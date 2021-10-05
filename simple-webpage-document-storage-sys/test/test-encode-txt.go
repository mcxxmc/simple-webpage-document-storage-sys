package test

import (
	"errors"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/filesys"
	"simple-webpage-document-storage-sys/logging"
)

var testTxtPath = common.NewTxtPath("test_createTxt.txt", common.DefaultUserId)
var testMsg = "CreateTxt(): this is a test."

func CreateTxt() {
	filesys.RewriteTxt(testTxtPath, testMsg)
	read := filesys.OpenTxt(testTxtPath)
	logging.Info(read)
	if testMsg != read {
		logging.Fatal(errors.New("the read does not match the expected message"))
	}
}

func DeleteTxt() {
	filesys.DeleteTxt(testTxtPath)
}
