package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/filesys"
	"simple-webpage-document-storage-sys/manager"
)

func wrapUpUserDirs(dirs *filesys.Collection) *DefaultResponse {
	return &DefaultResponse{Top: common.RootId, Dirs: *dirs}
}

func DefaultViewSkeleton(c *gin.Context) {
	manager.RegisterUser(common.DefaultUserId)
	c.JSON(http.StatusOK, wrapUpUserDirs(manager.UserCollection(common.DefaultUserId)))
}
