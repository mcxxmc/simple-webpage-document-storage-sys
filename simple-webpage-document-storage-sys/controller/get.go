package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/filesys"
	"simple-webpage-document-storage-sys/logging"
	"simple-webpage-document-storage-sys/manager"
)

func wrapUpUserDirs(dirs *filesys.Collection) *HierarchyResponse {
	return &HierarchyResponse{Ok: true, Top: common.RootId, Dirs: *dirs}
}

func ViewHierarchy(c *gin.Context) {
	uid, b := ExtractAndVerify(c)
	if !b {
		c.JSON(http.StatusOK, &HierarchyResponse{Ok: false, Top: common.RootId, Dirs: nil})
		return
	}
	c.JSON(http.StatusOK, wrapUpUserDirs(manager.UserCollection(uid)))
}

// Logout for logging out
func Logout(c *gin.Context) {
	uid, b := ExtractAndVerify(c)

	if !b {
		logging.Error(errors.New(errorInvalidToken))
		c.JSON(http.StatusOK, &CommonResponse{Ok: false, Msg: "fail to logout"})
		return
	}

	manager.UnregisterUser(uid)
	c.JSON(http.StatusOK, &CommonResponse{Ok: true, Msg: "you have logged out"})
}
