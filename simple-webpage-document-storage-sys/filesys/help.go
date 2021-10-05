package filesys

import "simple-webpage-document-storage-sys/common"

func checkPath(path string, args []string) bool {
	userTxtPrefix := common.Path_txt_prefix
	for _, str := range args {
		userTxtPrefix = userTxtPrefix + str + common.Slash
	}
	if path[:len(userTxtPrefix)] != userTxtPrefix {
		return false
	}
	return true
}
