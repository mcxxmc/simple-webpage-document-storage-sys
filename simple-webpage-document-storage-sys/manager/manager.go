package manager

import (
	"simple-webpage-document-storage-sys/filesys"
)

var defaultManager *Manager
var cached *filesys.IndexesOfUsers  //TODO: replace this with MySQL

// NumberOfUsers returns the number of users logged in
func NumberOfUsers() int {
	return defaultManager.numberOfUsers()
}

// RegisterUser registers a user with the manager
func RegisterUser(userId string) {
	defaultManager.registerUser(userId)
}

// UnregisterUser unregisters a user with the manager
func UnregisterUser(userId string) {
	defaultManager.unregisterUser(userId)
}

// UserCollection returns the directories of a given user; if the user does not exist, returns nil
func UserCollection(userId string) *filesys.Collection {
	return defaultManager.userCollection(userId)
}

// FetchTxt returns the txt file (name and content) by its id of and the user that owns it
func FetchTxt(userId string, fileId string) (string, string) {
	return defaultManager.fetchTxt(userId, fileId)
}

// ModifyTxt modifies a txt file
func ModifyTxt(userId string, fileId string, newContent string) bool {
	return defaultManager.modifyTxt(userId, fileId, newContent)
}

// CreateTxt creates a txt file with content
func CreateTxt(userId string, newFileId string, level int, newFileName string, newContent string) bool {
	return defaultManager.createTxt(userId, newFileId, level, newFileName, newContent)
}

// DeleteTxt deletes a txt file
func DeleteTxt(userId string, fileId string) bool {
	return defaultManager.deleteTxt(userId, fileId)
}

// MoveTxt moves a txt file (within the range of a user);
//
// it does not move the real files on the disk due to the special design of this project
func MoveTxt(userId string, fileId string, newParentId string) bool {
	return defaultManager.moveTxt(userId, fileId, newParentId)
}



// loads indexes of users into cache
func prepareCache(path string) {
	cached = filesys.LoadUserIndexes(path)
}

// StartManager configs an empty manager as the default manager;
// must be called before any other function calls
func StartManager(path string) {
	prepareCache(path)
	defaultManager = &Manager{Collections: make(map[string]*filesys.Collection)}
}


