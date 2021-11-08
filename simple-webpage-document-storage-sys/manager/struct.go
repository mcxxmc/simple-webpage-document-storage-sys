package manager

import (
	"errors"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/filesys"
	"simple-webpage-document-storage-sys/logging"
)

// Manager manages the collections of all users;
type Manager struct {
	// user uid (not name): their images.
	Collections map[string]*filesys.Collection

	// Note that map, slice and channel is passed by pointer by golang, so there's no need for an additional '*'
	// UserTimestamp map[string]"time"
	//TODO: add a go routine timer to manager to remove some logged users that have been idle for long
}

// returns the number of users logged in
func (manager *Manager) numberOfUsers() int {
	return len(manager.Collections)
}

// registers a user with the manager
func (manager *Manager) registerUser(userId string) {
	if _, ok := manager.Collections[userId]; ok {
		logging.Error(errors.New(errorUserAlreadyLoggedIn), logging.S(s1userId, userId))
		return
	}
	if userInfo, ok := (*cached)[userId]; ok {
		manager.Collections[userId] = filesys.LoadUserDirs(userInfo.Profile)
	} else {

	}
}

// unregisters a user with the manager
func (manager *Manager) unregisterUser(userId string) {
	if _, ok := manager.Collections[userId]; ok {
		delete(manager.Collections, userId)
	} else {
		logging.Error(errors.New(errorUserNotLoggedIn), logging.S(s1userId, userId))
	}
}

// returns the directories of a given user; if the user does not exist, returns nil
func (manager *Manager) userCollection(userId string) *filesys.Collection {
	if dirs, ok := manager.Collections[userId]; ok {
		return dirs
	}
	return nil
}

// returns the txt file (name and content) by its id of and the user that owns it
func (manager *Manager) fetchTxt(userId string, fileId string) (string, string) {
	collection := *manager.userCollection(userId)

	if collection == nil {
		logging.Error(errors.New(errorCollectionNotFound), logging.S(s1userId, userId))
		return "", ""
	}

	img, exist := collection[fileId]
	if exist == false {
		logging.Error(errors.New(errorFileIdNotExist), logging.S(s1userId, userId), logging.S(s1fileId, fileId))
		return "", ""
	}

	if img.Dir == true {
		logging.Error(errors.New(errorNotATxt), logging.S(s1userId, userId), logging.S(s1fileId, fileId))
		return "", ""
	}

	return img.Name, filesys.OpenTxt(img.Children[0])
}

// modifies a txt file
func (manager *Manager) modifyTxt(userId string, fileId string, newContent string) bool {
	collection := *manager.userCollection(userId)

	if collection == nil {
		logging.Error(errors.New(errorCollectionNotFound), logging.S(s1userId, userId))
		return false
	}

	img, exist := collection[fileId]
	if exist == false {
		logging.Error(errors.New(errorFileIdNotExist), logging.S(s1userId, userId), logging.S(s1fileId, fileId))
		return false
	}

	if img.Dir == true {
		logging.Error(errors.New(errorNotATxt), logging.S(s1userId, userId), logging.S(s1fileId, fileId))
		return false
	}

	return filesys.RewriteTxt(img.Children[0], newContent)
}

// creates a txt file with content
func (manager *Manager) createTxt(userId string, newFileId string, newFileName string,
	newContent string, parentId string) bool {

	if parentId == common.RootParentId {
		logging.Error(errors.New(errorUnauthorizedRoot), logging.S(s1userId, userId), logging.S(s1fileId, newFileId))
		return false
	}

	collection := *manager.userCollection(userId)

	// check if the user id is valid
	if collection == nil {
		logging.Error(errors.New(errorCollectionNotFound), logging.S(s1userId, userId))
		return false
	}

	// check if the newFileId is valid
	if _, exist := collection[newFileId]; exist == true {
		logging.Error(errors.New(errorFileIdAlreadyExists), logging.S(s1userId, userId), logging.S(s1fileId, newFileId))
		return false
	}

	// check if the parentId is valid
	parent, exist := collection[parentId]
	if exist != true {
		logging.Error(errors.New(errorInvalidNewParentId), logging.S(s1parentId, parentId))
		return false
	}

	// create a new Image
	newPath := common.NewTxtPath(newFileName, userId)
	level := parent.Level + 1
	newImg := &filesys.Image{ID: newFileId, Dir: false, Name: newFileName, Level: level,
		Children: []string{newPath}, Parent: parentId}

	// append that Image to Collection
	collection[newFileId] = newImg

	// append the new txt as a child to its parent
	collection[parentId].Children = append(collection[parentId].Children, newFileId)

	return filesys.RewriteTxt(newPath, newContent)
}

// deletes a txt file
func (manager *Manager) deleteTxt(userId string, fileId string) bool {
	collection := *manager.userCollection(userId)

	if collection == nil {
		logging.Error(errors.New(errorCollectionNotFound), logging.S(s1userId, userId))
		return false
	}

	img, exist := collection[fileId]
	if exist == false {
		logging.Error(errors.New(errorFileIdNotExist), logging.S(s1userId, userId), logging.S(s1fileId, fileId))
		return false
	}

	if img.Dir == true {
		logging.Error(errors.New(errorNotATxt), logging.S(s1userId, userId), logging.S(s1fileId, fileId))
		return false
	}

	oldParent, exist := collection[img.Parent]
	if exist != true {
		logging.Error(errors.New(errorInvalidOldParentId), logging.S(s1userId, userId), logging.S(s1parentId, img.Parent))
		return false
	}

	ok := filesys.DeleteTxt(collection[fileId].Children[0])
	if ok == true {
		oldParent.Children = removeString(oldParent.Children, fileId)
		delete(collection, fileId)
		return true
	} else {
		return false
	}
}

// moves a txt file or a directory (within the range of a user);
//
// it does not move the real files on the disk due to the special design of this project
func (manager *Manager) move(userId string, objectId string, newParentId string) bool {
	collection := *manager.userCollection(userId)

	if collection == nil {
		logging.Error(errors.New(errorCollectionNotFound), logging.S(s1userId, userId))
		return false
	}

	newParent, exist := collection[newParentId]

	if exist == false {
		logging.Error(errors.New(errorInvalidNewParentId), logging.S(s1parentId, newParentId))
		return false
	}

	if newParent.Dir == false {
		logging.Error(errors.New(errorNotADir), logging.S(s1parentId, newParentId))
	}

	img, exist := collection[objectId]
	if exist == false {
		logging.Error(errors.New(errorFileIdNotExist), logging.S(s1userId, userId), logging.S(s1objectId, objectId))
		return false
	}

	oldParentId := img.Parent
	oldParent, exist := collection[oldParentId]

	if exist == false {
		logging.Error(errors.New(errorInvalidOldParentId), logging.S(s1userId, userId), logging.S(s1parentId, oldParentId))
		return false
	}

	oldParent.Children = removeString(oldParent.Children, objectId)

	img.Parent = newParentId
	img.Level = newParent.Level + 1
	newParent.Children = append(newParent.Children, objectId)

	return true
}

// renames a file or a directory
func (manager *Manager) rename(userId string, objectId string, newName string) bool {
	collection := *manager.userCollection(userId)

	if collection == nil {
		logging.Error(errors.New(errorCollectionNotFound), logging.S(s1userId, userId))
		return false
	}

	img, exist := collection[objectId]
	if exist == false {
		logging.Error(errors.New(errorFileIdNotExist), logging.S(s1userId, userId), logging.S(s1objectId, objectId))
		return false
	}

	img.Name = newName
	// todo: if that is a file, considering updating its real filename and the path in img
	return true
}

// creates a new directory
func (manager *Manager) createDir(userId string, newDirId string, newDirName string, parentId string) bool {
	if parentId == common.RootParentId {
		logging.Error(errors.New(errorUnauthorizedRoot), logging.S(s1userId, userId), logging.S(s1dirId, newDirId))
		return false
	}

	collection := *manager.userCollection(userId)

	// check if the user id is valid
	if collection == nil {
		logging.Error(errors.New(errorCollectionNotFound), logging.S(s1userId, userId))
		return false
	}

	// check if the newDirId is valid
	if _, exist := collection[newDirId]; exist == true {
		logging.Error(errors.New(errorDirIdAlreadyExists), logging.S(s1userId, userId), logging.S(s1dirId, newDirId))
		return false
	}

	// check if the parentId is valid
	parent, exist := collection[parentId]
	if exist != true {
		logging.Error(errors.New(errorInvalidNewParentId), logging.S(s1parentId, parentId))
		return false
	}

	// create a new Image
	level := parent.Level + 1
	newImg := &filesys.Image{ID: newDirId, Dir: true, Name: newDirName, Level: level,
		Children: []string{}, Parent: parentId}

	// append that Image to Collection
	collection[newDirId] = newImg

	// append the new txt as a child to its parent
	collection[parentId].Children = append(collection[parentId].Children, newDirId)

	return true
}

// deletes a directory (only works when the directory is empty)
func (manager *Manager) deleteDir(userId string, dirId string) bool {
	collection := *manager.userCollection(userId)

	// check if the user id is valid
	if collection == nil {
		logging.Error(errors.New(errorCollectionNotFound), logging.S(s1userId, userId))
		return false
	}

	img, exist := collection[dirId]
	if exist == false {
		logging.Error(errors.New(errorDirIdNotExist), logging.S(s1userId, userId), logging.S(s1dirId, dirId))
		return false
	}

	if img.Dir == false {
		logging.Error(errors.New(errorNotADir), logging.S(s1userId, userId), logging.S(s1dirId, dirId))
		return false
	}

	if len(img.Children) > 0 {
		logging.Error(errors.New(errorDirNotEmpty), logging.S(s1userId, userId), logging.S(s1dirId, dirId))
		return false
	}

	oldParent, exist := collection[img.Parent]
	if exist != true {
		logging.Error(errors.New(errorInvalidOldParentId), logging.S(s1userId, userId), logging.S(s1parentId, img.Parent))
		return false
	}

	oldParent.Children = removeString(oldParent.Children, dirId)
	delete(collection, dirId)
	return true
}
