package manager

import file_sys "simple-webpage-document-storage-sys/file-sys"

type Manager struct {
	Users file_sys.IndexesOfUsers // the index containing all user ids
	UsersDirs map[string]file_sys.UserDirs  // user id : their dirs
}

func (manager *Manager) NumberOfUsers() int {
	return len(manager.Users)
}

func (manager *Manager) AddNewUser() {}

func (manager *Manager) RemoveUser() {}
