package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-webpage-document-storage-sys/logging"
	"simple-webpage-document-storage-sys/manager"
)

// GetFile deals with requests that ask for a specific file
func GetFile(c *gin.Context) {
	request := &FileRequest{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)
	fileName, content := manager.FetchTxt(request.User, request.Fid)
	c.JSON(http.StatusOK, &FileResponse{FileName: fileName, Content: content})
}
