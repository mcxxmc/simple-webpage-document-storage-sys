package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/manager"
)

func DefaultViewSkeleton(c *gin.Context) {
	manager.RegisterUser(common.DefaultUser)
	defer manager.UnregisterUser(common.DefaultUser)
	c.JSON(http.StatusOK, manager.UserDirs(common.DefaultUser))
}
