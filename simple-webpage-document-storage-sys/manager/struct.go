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
		logging.Info("User: " + userId + " already logged in.")
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
		logging.Info("Fail to unregister user; userId invalid.")
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
		logging.Error(errors.New(common.ErrorCollectionNotFound(userId)))
		return "", ""
	}
	img := collection[fileId]
	return img.Name, filesys.OpenTxt(img.Children[0])
}

// modifies a txt file
func (manager *Manager) modifyTxt(userId string, fileId string, newContent string) bool {
	collection := *manager.userCollection(userId)
	if collection == nil {
		logging.Error(errors.New(common.ErrorCollectionNotFound(userId)))
		return false
	}
	img := collection[fileId]
	return filesys.RewriteTxt(img.Children[0], newContent)
}

// creates a txt file with content
func (manager *Manager) createTxt(userId string, newFileId string, level int, newFileName string, newContent string) bool {
	collection := *manager.userCollection(userId)
	if collection == nil {
		logging.Error(errors.New(common.ErrorCollectionNotFound(userId)))
		return false
	}
	// check if the newFileId is valid
	if _, exist := collection[newFileId]; exist == true {
		logging.Error(errors.New(common.ErrorFileIdAlreadyExists(newFileId)))
		return false
	}
	// create a new Image
	newPath := common.NewTxtPath(newFileName, userId)
	newImg := &filesys.Image{ID: newFileId, Dir: false, Name: newFileName, Level: level,
		Children: []string{newPath}}
	// append that Image to Collection
	collection[newFileId] = newImg
	return filesys.RewriteTxt(newPath, newContent)
}

// deletes a txt file
func (manager *Manager) deleteTxt(userId string, fileId string) bool {
	collection := *manager.userCollection(userId)
	if collection == nil {
		logging.Error(errors.New(common.ErrorCollectionNotFound(userId)))
		return false
	}
	ok := filesys.DeleteTxt(collection[fileId].Children[0])
	if ok == true {
		delete(collection, fileId)
		return true
	} else {
		return false
	}
}

// moves a txt file (within the range of a user);
//
// it does not move the real files on the disk due to the special design of this project
func (manager *Manager) moveTxt(userId string, fileId string, newParentId string) bool {
	collection := *manager.userCollection(userId)
	if collection == nil {
		logging.Error(errors.New(common.ErrorCollectionNotFound(userId)))
		return false
	}
	if _, exist := collection[newParentId]; exist == false {
		logging.Error(errors.New(common.ErrorInvalidNewParentId(newParentId)))
		return false
	}

	oldParentId := collection[fileId].Parent
	collection[oldParentId].Children = removeString(collection[oldParentId].Children, fileId)
	collection[fileId].Parent = newParentId
	collection[newParentId].Children = append(collection[newParentId].Children, fileId)
	return true
}