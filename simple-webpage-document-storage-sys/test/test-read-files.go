package test

import (
	"fmt"
	"simple-webpage-document-storage-sys/common"
	file_sys "simple-webpage-document-storage-sys/file-sys"
)

func TestReadFiles() {
	userindexes := file_sys.LoadUserIndexes(common.Path_index_of_users)
	nxt := userindexes[common.DefaultUser]
	userdirs := file_sys.LoadUserDirs(nxt)
	for _, v := range userdirs {
		if v.Dir == false {
			fmt.Println(file_sys.OpenTxt(v.Link[0]))
		}
	}
}
