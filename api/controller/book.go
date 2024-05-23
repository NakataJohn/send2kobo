package controller

import (
	"fmt"
	"net/http"
	"os"
	"path"
	fp "path/filepath"
	"send2kobo/bootstrap"
	"send2kobo/domain"
	"send2kobo/internal/cmdutil"
	"send2kobo/internal/fileutil"
	"send2kobo/logger"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookController struct {
	BookUsecase domain.BookUsecase
	Env         *bootstrap.Env
}

// @Summary Fetch All Books
// @Tags Book
// @Description  获取图书列表
// @Security token
// @Accept  json
// @Produce  json
// @Success 200 {object} []domain.Book
// @Failure 200 {object} domain.ErrorResponse
// @Router /api/v1/book [get]
func (bc *BookController) GetBooks(c *gin.Context) {
	books, err := bc.BookUsecase.Fetch(c)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// @Summary Fetch Book by id
// @Tags Book
// @Description  获取图书
// @Security token
// @Produce  json
// @Param id path domain.RequestID true "用id查询Book"
// @Success 200 {object} domain.Book
// @Failure 200 {object} domain.ErrorResponse
// @Router /api/v1/book/{id}/ [get]
func (bc *BookController) GetBookByID(c *gin.Context) {
	ID := com.StrTo(c.Param("id")).String()

	book, err := bc.BookUsecase.GetByID(c, ID)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Book not found with the given id"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// @Summary Download Book by id
// @Tags Book
// @Description  下载图书
// @Security token
// @Produce  json
// @Param id path domain.RequestID true "通过id下载Book"
// @Success 200 {object} domain.SuccessResponse
// @Failure 200 {object} domain.ErrorResponse
// @Router /api/v1/book/{id}/download [get]
func (bc *BookController) DownloadBookByID(c *gin.Context) {
	ID := com.StrTo(c.Param("id")).String()

	book, err := bc.BookUsecase.GetByID(c, ID)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Book not found with the given id"})
		return
	}

	filePath := book.Kepubpath
	fileTmp, errByOpenFile := os.Open(filePath)
	defer fileTmp.Close()

	//获取文件的名称
	fileName := path.Base(filePath)

	if !fileutil.Exists(filePath) || errByOpenFile != nil {
		logger.Error("获取文件失败")
		c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: "文件获取失败"})
		return
	}

	stat, err := fileTmp.Stat()
	if err != nil {
		logger.Error("获取文件状态失败")
		c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: "文件获取失败"})
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	//强制浏览器下载
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	//浏览器下载或预览
	c.Header("Content-Disposition", "inline;filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.Writer.Header().Set("Content,Length", strconv.FormatInt(stat.Size(), 10))
	c.Writer.Flush()

	c.File(filePath)

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: book.Title + "：可提供下载。"})
}

// @Summary Delete Book by id
// @Tags Book
// @Description  删除图书
// @Security token
// @Produce  json
// @Param id path domain.RequestID true "用id删除Book"
// @Success 200 {object} domain.SuccessResponse
// @Failure 200 {object} domain.ErrorResponse
// @Router /api/v1/book/{id}/ [delete]
func (bc *BookController) DeleteBookByID(c *gin.Context) {
	ID := com.StrTo(c.Param("id")).String()

	err := bc.BookUsecase.DeleteByID(c, ID)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Book not found with the given id"})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "删除成功"})
}

// @Summary upload epub book file
// @Tags Book
// @Description 传文件：epub文件
// @Accept multipart/form-data
// @Security token
// @Param title formData string false "书籍名"
// @Param fileName formData file true "file"
// @Produce  json
// @Success 200 {object} domain.SuccessResponse
// @Failure 200 {object} domain.ErrorResponse
// @Router /api/v1/book/upload [post]
func (bc *BookController) Upload(c *gin.Context) {
	title := c.PostForm("title")
	file, err := c.FormFile("fileName")
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "文件上传失败。"})
		return
	}
	filename := fileutil.Filenamify(file.Filename)
	if !strings.HasSuffix(filename, "epub") {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "请上传epub文件。"})
		return
	}

	filepath := fileutil.GetFilePath(bc.Env.UploadPath, filename)
	if fileutil.Exists(filepath) {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "请勿上传同一文件。"})
		return
	}
	err = c.SaveUploadedFile(file, filepath)
	if err != nil || !fileutil.Exists(filepath) {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "文件保存失败。"})
		return
	}

	hash, _ := fileutil.HashFileMd5(filepath)
	kfile := strings.Split(filename, "epub")[0] + "kepub.epub"
	kpath := fp.Join(bc.Env.KepubPath, kfile)
	// 转kepub
	err = exkepub(filepath, kpath)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "文件转换失败。"})
		return
	}

	book := &domain.Book{
		ID:        primitive.NewObjectID(),
		Hash:      hash,
		Title:     title,
		Path:      filepath,
		Kepubpath: kpath,
		KHash:     "",
	}
	bc.BookUsecase.Create(c, book)

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "上传成功。"})
}

func exkepub(epath, kpath string) error {
	cmdline := fmt.Sprintf("kepubify %s -o %s", epath, kpath)
	logger.Info(cmdline)
	_, err := cmdutil.RunCmd(cmdline)
	return err
}
