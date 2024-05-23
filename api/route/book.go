package route

import (
	"send2kobo/api/controller"
	"send2kobo/bootstrap"
	"send2kobo/domain"
	"send2kobo/mongo"
	"send2kobo/repository"
	"send2kobo/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewBookRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	br := repository.NewBookRepository(db, domain.CollectionBook)
	bc := &controller.BookController{
		BookUsecase: usecase.NewBookUsecase(br, timeout),
		Env:         env,
	}
	group.GET("/book", bc.GetBooks)
	group.GET("/book/:id/", bc.GetBookByID)
	group.DELETE("/book/:id/", bc.DeleteBookByID)
	group.POST("/book/upload", bc.Upload)
	group.GET("/book/:id/download", bc.DownloadBookByID)
}
