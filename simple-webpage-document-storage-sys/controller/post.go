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

// ExtractAndVerify checks if the uid is in the token
//
// returns the uid (if exists) and a bool.
func ExtractAndVerify(c *gin.Context) (string, bool) {
	uid, exist := c.Get(common.TokenUid)
	if !exist{
		return "", false
	}
	return uid.(string), true
}

// Login for logging in
func Login(c *gin.Context) {
	request := &RequestLogin{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	uid, b := manager.VerifyUserPassword(request.Username, request.Password)

	if !b {
		c.JSON(http.StatusOK, &LoginResponse{Ok: false, Token: ""})
		return
	}

	tk, err := token.GenerateToken(uid)
	if err != nil {
		logging.Error(err)
		c.JSON(http.StatusOK, &LoginResponse{Ok: false, Token: ""})
		return
	}

	manager.RegisterUser(uid)

	c.JSON(http.StatusOK, &LoginResponse{Ok: true, Token: tk})
}

// GetFile deals with requests that ask for a specific file
func GetFile(c *gin.Context) {
	request := &RequestFile{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	uid, b := ExtractAndVerify(c)

	if !b {
		c.JSON(http.StatusOK, &CommonResponse{Ok: false, Msg: "fail to get file"})
		return
	}

	fileName, content := manager.FetchTxt(uid, request.Fid)
	c.JSON(http.StatusOK, &FileResponse{Ok: true, FileName: fileName, Content: content})
}

// ModifyFile deals with requests that modify a specific file
func ModifyFile(c *gin.Context) {
	request := &RequestModifyFile{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	uid, b := ExtractAndVerify(c)

	if !b {
		c.JSON(http.StatusOK, &CommonResponse{Ok: false, Msg: "fail to modify file"})
		return
	}

	b = manager.ModifyTxt(uid, request.Fid, request.NewC)
	if b {
		c.JSON(http.StatusOK, &CommonResponse{Ok: true, Msg: "file modified"})
	} else {
		logging.Error(errors.New(errorModifyingTxt),
			logging.SS{S1: s1UserId, S2: uid}, logging.SS{S1: s1FileId, S2: request.Fid},
			logging.SS{S1: s1NewC, S2: request.NewC})
		c.JSON(http.StatusOK, &CommonResponse{Ok: false, Msg: "fail to modify file"})
	}
}

// Rename deals with requests that rename a txt or a directory
func Rename(c *gin.Context) {
	request := &RequestRename{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	uid, b := ExtractAndVerify(c)

	if !b {
		c.JSON(http.StatusOK, &CommonResponse{Ok: false, Msg: "fail to rename"})
		return
	}

	b = false
	if request.Dir {
		b = manager.RenameDir(uid, request.ObjId, request.NewName)
	} else {
		b = manager.RenameTxt(uid, request.ObjId, request.NewName)
	}
	if b {
		c.JSON(http.StatusOK, &CommonResponse{Ok: true, Msg: "renamed"})
	} else {
		logging.Error(errors.New(errorRenaming),
			logging.SS{S1: s1UserId, S2: uid}, logging.SS{S1: s1ObjId, S2: request.ObjId},
			logging.SS{S1: s1IsDir, S2: strconv.FormatBool(request.Dir)}, logging.SS{S1: s1newName, S2: request.NewName})
		c.JSON(http.StatusOK, &CommonResponse{Ok: false, Msg: "fail to rename"})
	}
}

// Create deals with requests that create a new file or a new dir
func Create(c *gin.Context) {
	request := &RequestCreate{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	uid, b := ExtractAndVerify(c)

	if !b {
		c.JSON(http.StatusOK, &CommonResponse{Ok: false, Msg: "fail to create"})
		return
	}

	b = false
	objId := common.GenerateRandomId(64)
	if request.Dir {
		b = manager.CreateDir(uid, objId, request.Name, request.ParentId)
	} else {
		b = manager.CreateTxt(uid, objId, request.Name, request.NewC, request.ParentId)
	}
	if b {
		c.JSON(http.StatusOK, &CommonResponse{Ok: true, Msg: "created"})  // TODO may need to send a new dirs
	} else {
		logging.Error(errors.New(errorCreating),
			logging.SS{S1: s1UserId, S2: uid}, logging.SS{S1: s1ObjId, S2: objId},
			logging.SS{S1: s1IsDir, S2: strconv.FormatBool(request.Dir)}, logging.SS{S1: s1newName, S2: request.Name},
			logging.SS{S1: s1NewC, S2: request.NewC}, logging.SS{S1: s1ParentId, S2: request.ParentId})
		c.JSON(http.StatusOK, &CommonResponse{Ok: false, Msg: "fail to create"})
	}
}

// Delete deals with requests that delete a file or a dir
func Delete(c * gin.Context) {
	request := &RequestDelete{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	uid, b := ExtractAndVerify(c)

	if !b {
		c.JSON(http.StatusOK, &CommonResponse{Ok: false, Msg: "fail to delete"})
		return
	}

	b = false
	if request.Dir {
		b = manager.DeleteDir(uid, request.ObjId)
	} else {
		b = manager.DeleteTxt(uid, request.ObjId)
	}
	if b {
		c.JSON(http.StatusOK, &CommonResponse{Ok: true, Msg: "deleted"})
	} else {
		logging.Error(errors.New(errorDeleting),
			logging.SS{S1: s1UserId, S2: uid}, logging.SS{S1: s1ObjId, S2: request.ObjId},
			logging.SS{S1: s1IsDir, S2: strconv.FormatBool(request.Dir)})
		c.JSON(http.StatusOK, &CommonResponse{Ok: false, Msg: "fail to delete"})
	}
}

// Move deals with requests that move a file or a dir to a different location
func Move(c *gin.Context) {
	request := &RequestMove{}
	err := c.BindJSON(request)
	logging.ConditionalLogError(err)

	uid, b := ExtractAndVerify(c)

	if !b {
		c.JSON(http.StatusOK, &CommonResponse{Ok: false, Msg: "fail to move"})
		return
	}

	b = false
	if request.Dir {
		b = manager.MoveDir(uid, request.ObjId, request.NewParentId)
	} else {
		b = manager.MoveTxt(uid, request.ObjId, request.NewParentId)
	}
	if b {
		c.JSON(http.StatusOK, &CommonResponse{Ok: true, Msg: "moved"})  // TODO may need to send a new dirs
	} else {
		logging.Error(errors.New(errorMoving),
			logging.SS{S1: s1UserId, S2: uid}, logging.SS{S1: s1ObjId, S2: request.ObjId},
			logging.SS{S1: s1IsDir, S2: strconv.FormatBool(request.Dir)},
			logging.SS{S1: s1ParentId, S2: request.NewParentId})
		c.JSON(http.StatusOK, &CommonResponse{Ok: false, Msg: "fail to move"})
	}
}
