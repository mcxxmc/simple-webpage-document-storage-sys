package main

import (
	"simple-webpage-document-storage-sys/logging"
	"simple-webpage-document-storage-sys/test"
)

func main() {
	defer logging.Sync()

	// test reading files
	test.ReadFiles()

	// test manager
	test.ManagerDefaultView()
}
