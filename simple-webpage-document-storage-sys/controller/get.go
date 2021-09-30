package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/manager"
)

func DefaultViewSkeleton(c *gin.Context) {
	manager.RegisterUser(common.DefaultUser)
	c.JSON(http.StatusOK, wrapUpUserDirs(manager.UserDirs(common.DefaultUser)))
}
