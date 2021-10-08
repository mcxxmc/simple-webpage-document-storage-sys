package main

import (
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/logging"
	"simple-webpage-document-storage-sys/manager"
	"simple-webpage-document-storage-sys/test"
)

//TODO: more test cases

func main() {
	defer logging.Sync()

	manager.StartManager(common.Path_index_of_users)

	// test manager
	test.ManagerDefaultView()
}
