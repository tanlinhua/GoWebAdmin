package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/pkg/utils"
)

// curl -X POST http://127.0.0.1:2014/api/test/upload -F "upload=@D:\1.png" -H "Content-Type: multipart/form-data"
func TestUpload(c *gin.Context) {
	file, err := c.FormFile("upload")
	if err != nil {
		c.String(500, "上传文件出错:"+err.Error())
		return
	}
	fileName := utils.UUID() + "_" + file.Filename
	dst := "runtime/upload/" + fileName
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	}
	c.String(http.StatusOK, fileName)
}
