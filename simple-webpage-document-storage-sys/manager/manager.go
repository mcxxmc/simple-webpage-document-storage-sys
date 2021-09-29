package manager

import (
	"simple-webpage-document-storage-sys/filesys"
	"simple-webpage-document-storage-sys/logging"
)

type Manager struct {
	// username : their dirs.
	// Note that map, slice and channel is passed by pointer by golang, so there's no need for an additional '*'
	UsersDirs map[string]*filesys.UserDirs

	// UserTimestamp map[string]"time"
}

var defaultManager *Manager
var cached *filesys.IndexesOfUsers  //TODO: replace this with MySQL

// returns the number of users logged in
func (manager *Manager) numberOfUsers() int {
	return len(manager.UsersDirs)
}

// registers a user with the manager
func (manager *Manager) registerUser(username string) {
	if profile, ok := (*cached)[username]; ok {
		manager.UsersDirs[username] = filesys.LoadUserDirs(profile)
	} else {
		logging.Info("Fail to register user; username invalid.")
	}
}

// unregisters a user with the manager
func (manager *Manager) unregisterUser(username string) {
	if _, ok := manager.UsersDirs[username]; ok {
		delete(manager.UsersDirs, username)
	} else {
		logging.Info("Fail to unregister user; username invalid.")
	}
}

// returns the directories of a given user; if the user does not exist, returns nil
func (manager *Manager) userDirs(username string) *filesys.UserDirs {
	if dirs, ok := manager.UsersDirs[username]; ok {
		return dirs
	}
	return nil
}

// NumberOfUsers returns the number of users logged in
func NumberOfUsers() int {
	return defaultManager.numberOfUsers()
}

// RegisterUser registers a user with the manager
func RegisterUser(username string) {
	defaultManager.registerUser(username)
}

// UnregisterUser unregisters a user with the manager
func UnregisterUser(username string) {
	defaultManager.unregisterUser(username)
}

// UserDirs returns the directories of a given user; if the user does not exist, returns nil
func UserDirs(username string) *filesys.UserDirs {
	return defaultManager.userDirs(username)
}

// loads indexes of users into cache
func prepareCache(path string) {
	cached = filesys.LoadUserIndexes(path)
}

// StartManager configs an empty manager as the default manager;
// must be called before any other function calls
func StartManager(path string) {
	prepareCache(path)
	defaultManager = &Manager{UsersDirs: make(map[string]*filesys.UserDirs)}
}

//TODO: add a timer to manager to remove some logged users that have been idle for long
