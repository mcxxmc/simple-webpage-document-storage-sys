package common

const Path_user_profiles_prefix = "D:/ProgrammingTestsGo/simple-webpage-document-storage-sys/file-sys/users/"
const Path_index_of_users = "D:/ProgrammingTestsGo/simple-webpage-document-storage-sys/file-sys/index/index_of_users.json"
const Path_txt_prefix = "D:/ProgrammingTestsGo/simple-webpage-document-storage-sys/file-sys/txt/"
const Slash = "/"
const ProfileAppendix = ".json"

// NewTxtPath returns the full path of the new txt file;
// the 2nd parameter should be the uid instead of name of the user
func NewTxtPath(txtName string, userUid string) string {
	return Path_txt_prefix + userUid + Slash + txtName
}

// NewUserProfilePath returns the full path of a new user profile
func NewUserProfilePath(uid string) string {
	return Path_user_profiles_prefix + uid + ProfileAppendix
}
