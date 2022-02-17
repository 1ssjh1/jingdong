package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path"
	"strconv"
	"time"
)

func SaveFile(file *multipart.FileHeader, c *gin.Context) (string, error) {
	if file.Size > 5242880 {
		err := errors.New("文件过大，换个小点的吧")
		return "", err
	}
	ImagePath := path.Ext(file.Filename)
	fmt.Println(ImagePath)
	//if (ImagePath != ".jpg") && (ImagePath != ".png") {
	//	err := errors.New("文件不是我想要的再试试呗")
	//	return "", err
	//}
	t := time.Now().Unix()
	url := strconv.FormatInt(t, 10)
	err := c.SaveUploadedFile(file, "/www/static/"+url+ImagePath)
	fmt.Println(err)
	if err != nil {
		err = errors.New("文件保存失败")
		return url + ImagePath, err
	}
	return url + ImagePath, err
}
