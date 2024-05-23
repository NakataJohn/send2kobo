package fileutil

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

func GetFilePath(uploadPath, fileName string) string {
	// 创建多层级目录
	os.MkdirAll(uploadPath, os.ModePerm)

	filePath := filepath.Join(uploadPath, fileName)
	return filePath
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func HashFileMd5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

func Filenamify(input string) string {
	pattern := regexp.MustCompile(`[\s\\/:\*\?"<>\|]`)
	output := pattern.ReplaceAllString(input, "_")
	return output
}
