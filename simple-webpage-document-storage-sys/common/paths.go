package common

const Path_users_prefix = "D:/ProgrammingTestsGo/simple-webpage-document-storage-sys/file-sys/users/"
const Path_index_of_users = "D:/ProgrammingTestsGo/simple-webpage-document-storage-sys/file-sys/index/index_of_users.json"
const Path_txt_prefix = "D:/ProgrammingTestsGo/simple-webpage-document-storage-sys/file-sys/txt/"
const Slash = "/"

// NewTxtPath returns the full path of the new txt file;
// the 2nd parameter should be the uid instead of name of the user
func NewTxtPath(txtName string, userUid string) string {
	return Path_txt_prefix + userUid + Slash + txtName
}
