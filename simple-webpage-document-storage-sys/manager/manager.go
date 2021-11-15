package manager

import (
	"errors"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/filesys"
	"simple-webpage-document-storage-sys/logging"
)

var defaultManager *Manager
var cached *filesys.IndexesOfUsers
// TODO: replace "cached" with MySQL; also, it can be created as 2 related tables (uid and uname as keys for each table)

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
//
// note that the real new filename on disk equals to newFileId.txt instead of newFileName (to avoid collision)
func CreateTxt(userId string, newFileId string, newFileName string, newContent string, parentId string) bool {
	return defaultManager.createTxt(userId, newFileId, newFileName, newContent, parentId)
}

// DeleteTxt deletes a txt file
func DeleteTxt(userId string, fileId string) bool {
	return defaultManager.deleteTxt(userId, fileId)
}

// MoveTxt moves a txt file (within the range of a user);
//
// it does not move the real files on the disk due to the special design of this project
func MoveTxt(userId string, fileId string, newParentId string) bool {
	return defaultManager.move(userId, fileId, newParentId)
}

// RenameTxt renames a txt file
func RenameTxt(userId string, fileId string, newName string) bool {
	return defaultManager.rename(userId, fileId, newName)
}

// RenameDir renames a directory
func RenameDir(userId string, dirId string, newName string) bool {
	return defaultManager.rename(userId, dirId, newName)
}

// MoveDir moves a directory (within the range of a user);
//
// it does not move the real files on the disk due to the special design of this project
func MoveDir(userId string, dirId string, newParentId string) bool {
	return defaultManager.move(userId, dirId, newParentId)
}

// CreateDir creates a new directory
func CreateDir(userId string, newDirId string, newDirName string, parentId string) bool {
	return defaultManager.createDir(userId, newDirId, newDirName, parentId)
}

// DeleteDir deletes a new directory (the directory needs to be empty)
func DeleteDir(userId string, dirId string) bool {
	return defaultManager.deleteDir(userId, dirId)
}



// loads indexes of users into cache
func prepareCache(path string) {
	cached = filesys.LoadUserIndexes(path)
}

// StartManager configs an empty manager as the default manager;
// must be called before any other function calls
func StartManager(path string) {
	prepareCache(path)
	defaultManager = &Manager{Collections: make(map[string]*filesys.Collection), Modified: make(map[string]bool)}
}

// saveUserCollection saves the user collection (as a JSON file) to disk
func saveUserCollection(userId string) {
	err := filesys.SaveUserCollection((*cached)[userId].Profile, defaultManager.userCollection(userId))
	logging.ConditionallyLoggingError(err, logging.S(s1userId, userId))
}

// SaveModifiedUserCollections saves all the modified user collections to disk
//
// should be called periodically or when the program shuts down.
func SaveModifiedUserCollections() {
	for uid, b := range defaultManager.Modified {
		if b {
			saveUserCollection(uid)
			defaultManager.Modified[uid] = false
		}
	}
}

// SaveIndexesOfUsers saves the indexes of users to the disk
//
// should be called periodically or when the program shuts down.
func SaveIndexesOfUsers() {
	err := filesys.SaveIndexesOfUsers(common.Path_index_of_users, cached)
	logging.ConditionallyLoggingError(err)
}

// SaveWhenShuttingDown saves necessary info to disk when shutting down
func SaveWhenShuttingDown() {
	SaveModifiedUserCollections()
	SaveIndexesOfUsers()
}

// VerifyUserPassword verifies the username - password pair and returns the userId and a bool
//
// the bool will be false if the userId does not exist or the pair is invalid
//
// TODO: change to SQL; detailed return (if the username does not exist)
func VerifyUserPassword(username, password string) (string, bool) {
	for _, user := range *cached {
		if user.Name == username {
			if user.Password == password {
				return user.Uid, true
			}
			return "", false
		}
	}
	return "", false
}

// IsUsernameAvailable checks if the username is available. True if the username has not been used.
//
// TODO: change to SQL
func IsUsernameAvailable(username string) bool {
	for _, user := range *cached {
		if user.Name == username {
			return false
		}
	}
	return true
}

// IsUserIdAvailable checks if the user id is available. True if the user id has not been used.
//
// TODO: change to SQL
func IsUserIdAvailable(id string) bool {
	if _, exist := (*cached)[id]; exist {
		return false
	}
	return true
}

// CreateNewUser creates a new user and add it to cache
//
// TODO: change to SQL
func CreateNewUser(name, password string) (string, bool) {
	if !IsUsernameAvailable(name) {
		logging.Error(errors.New(errorUsernameUsed), logging.S(s1username, name))
		return "username unavailable", false
	}

	uid := common.GenerateRandomId(64)
	for {
		if IsUserIdAvailable(uid) {
			break
		}
		uid = common.GenerateRandomId(64)
	}

	profile := common.NewUserProfilePath(uid)
	err := filesys.CreateNewUserProfile(profile)  // create and save the new profile to disk
	if err != nil {
		logging.Error(err)
		return "internal error: fail to create user profile", false
	}

	newUser := &filesys.UserInfo{Name: name, Uid: uid, Password: password, Profile: profile}
	(*cached)[uid] = newUser  // add it to the cache

	// no need to add it to the manager here; that will be done by the controller

	// IMPORTANT: create a physical dir for it on disk
	b := filesys.CreatePhysicalDirForUser(uid)
	if !b {
		return "internal error: fail to create physical dir", false
	}

	return uid, true
}


