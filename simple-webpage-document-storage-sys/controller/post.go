package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/logging"
	"simple-webpage-document-storage-sys/manager"
	"simple-webpage-document-storage-sys/token"
	"strconv"
)

// verify checks if the uid in the token matches the uid in the request
func verify(c *gin.Context, userId string) bool {
	uid, exist := c.Get(common.TokenUid)
	if (!exist && userId != common.DefaultUserId) || uid != userId {
		return false
	}
	return true
}

// Login for logging in
func Login(c *gin.Context) {
	request := &RequestLogin{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	uid, b := manager.VerifyUserPassword(request.Username, request.Password)

	if !b {
		c.JSON(http.StatusOK, &LoginResponse{Success: false, User: "", Token: ""})
		return
	}

	tk, err := token.GenerateToken(uid)
	if err != nil {
		logging.Error(err)
		return
	}

	manager.RegisterUser(uid)

	c.JSON(http.StatusOK, &LoginResponse{Success: true, User: uid, Token: tk})
}

// Logout for logging out
func Logout(c *gin.Context) {
	request := &RequestLogout{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	uid, b := token.DecodeToken(request.Token)

	if !b || uid != request.User {
		logging.Error(errors.New(errorInvalidToken),
			logging.S("uid", request.User), logging.S("uid from token", uid), logging.S("token", request.Token))
		c.Status(http.StatusBadRequest)
		return
	}

	manager.UnregisterUser(uid)
	c.Status(http.StatusOK)
}

// GetFile deals with requests that ask for a specific file
func GetFile(c *gin.Context) {
	request := &RequestFile{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	if !verify(c, request.User) {
		c.Status(http.StatusBadRequest)
		return
	}

	fileName, content := manager.FetchTxt(request.User, request.Fid)
	c.JSON(http.StatusOK, &FileResponse{FileName: fileName, Content: content})
}

// ModifyFile deals with requests that modify a specific file
func ModifyFile(c *gin.Context) {
	request := &RequestModifyFile{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	if ! verify(c, request.User) {
		c.Status(http.StatusBadRequest)
		return
	}

	b := manager.ModifyTxt(request.User, request.Fid, request.NewC)
	if b {
		c.Status(http.StatusOK)
	} else {
		logging.Error(errors.New(errorModifyingTxt),
			logging.SS{S1: s1UserId, S2: request.User}, logging.SS{S1: s1FileId, S2: request.Fid},
			logging.SS{S1: s1NewC, S2: request.NewC})
		c.Status(http.StatusBadRequest)
	}
}

// Rename deals with requests that rename a txt or a directory
func Rename(c *gin.Context) {
	request := &RequestRename{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	if !verify(c, request.User) {
		c.Status(http.StatusBadRequest)
		return
	}

	b := false
	if request.Dir {
		b = manager.RenameDir(request.User, request.ObjId, request.NewName)
	} else {
		b = manager.RenameTxt(request.User, request.ObjId, request.NewName)
	}
	if b {
		c.Status(http.StatusOK)
	} else {
		logging.Error(errors.New(errorRenaming),
			logging.SS{S1: s1UserId, S2: request.User}, logging.SS{S1: s1ObjId, S2: request.ObjId},
			logging.SS{S1: s1IsDir, S2: strconv.FormatBool(request.Dir)}, logging.SS{S1: s1newName, S2: request.NewName})
		c.Status(http.StatusBadRequest)
	}
}

// Create deals with requests that create a new file or a new dir
func Create(c *gin.Context) {
	request := &RequestCreate{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	if !verify(c, request.User) {
		c.Status(http.StatusBadRequest)
		return
	}

	b := false
	objId := common.GenerateRandomId(64)
	if request.Dir {
		b = manager.CreateDir(request.User, objId, request.Name, request.ParentId)
	} else {
		b = manager.CreateTxt(request.User, objId, request.Name, request.NewC, request.ParentId)
	}
	if b {
		c.Status(http.StatusOK)  // TODO may need to send a new dirs
	} else {
		logging.Error(errors.New(errorCreating),
			logging.SS{S1: s1UserId, S2: request.User}, logging.SS{S1: s1ObjId, S2: objId},
			logging.SS{S1: s1IsDir, S2: strconv.FormatBool(request.Dir)}, logging.SS{S1: s1newName, S2: request.Name},
			logging.SS{S1: s1NewC, S2: request.NewC}, logging.SS{S1: s1ParentId, S2: request.ParentId})
		c.Status(http.StatusBadRequest)
	}
}

// Delete deals with requests that delete a file or a dir
func Delete(c * gin.Context) {
	request := &RequestDelete{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	if !verify(c, request.User) {
		c.Status(http.StatusBadRequest)
		return
	}

	b := false
	if request.Dir {
		b = manager.DeleteDir(request.User, request.ObjId)
	} else {
		b = manager.DeleteTxt(request.User, request.ObjId)
	}
	if b {
		c.Status(http.StatusOK)
	} else {
		logging.Error(errors.New(errorDeleting),
			logging.SS{S1: s1UserId, S2: request.User}, logging.SS{S1: s1ObjId, S2: request.ObjId},
			logging.SS{S1: s1IsDir, S2: strconv.FormatBool(request.Dir)})
		c.Status(http.StatusBadRequest)
	}
}

// Move deals with requests that move a file or a dir to a different location
func Move(c *gin.Context) {
	request := &RequestMove{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	if !verify(c, request.User) {
		c.Status(http.StatusBadRequest)
		return
	}

	b := false
	if request.Dir {
		b = manager.MoveDir(request.User, request.ObjId, request.NewParentId)
	} else {
		b = manager.MoveTxt(request.User, request.ObjId, request.NewParentId)
	}
	if b {
		c.Status(http.StatusOK)  // TODO may need to send a new dirs
	} else {
		logging.Error(errors.New(errorMoving),
			logging.SS{S1: s1UserId, S2: request.User}, logging.SS{S1: s1ObjId, S2: request.ObjId},
			logging.SS{S1: s1IsDir, S2: strconv.FormatBool(request.Dir)},
			logging.SS{S1: s1ParentId, S2: request.NewParentId})
		c.Status(http.StatusBadRequest)
	}
}
