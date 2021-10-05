package test

import (
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/filesys"
	"simple-webpage-document-storage-sys/logging"
)

func ReadFiles() {
	userindexes := *filesys.LoadUserIndexes(common.Path_index_of_users)
	nxt := userindexes[common.DefaultUserId]
	userdirs := *filesys.LoadUserDirs(nxt.Profile)
	for _, v := range userdirs {
		if v.Dir == false {
			logging.Info(filesys.OpenTxt(v.Children[0]))
		}
	}
}
