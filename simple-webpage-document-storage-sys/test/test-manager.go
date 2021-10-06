package test

import (
	"errors"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/logging"
	"simple-webpage-document-storage-sys/manager"
	"strconv"
)

func ManagerDefaultView() {
	n := manager.NumberOfUsers()
	logging.Info("number of users logged in: ", logging.S("number:", strconv.Itoa(n)))
	if n != 0 {
		logging.Fatal(errors.New("numbers of users do not match"))
	}

	// Test: new register
	manager.RegisterUser(common.DefaultUserId)
	n = manager.NumberOfUsers()
	logging.Info("number of users logged in: ", logging.S("number:", strconv.Itoa(n)))
	if n != 1 {
		logging.Fatal(errors.New("numbers of users do not match"))
	}

	// Test:

	// Test: log out
	manager.UnregisterUser(common.DefaultUserId)
	n = manager.NumberOfUsers()
	logging.Info("number of users logged in: ", logging.S("number:", strconv.Itoa(n)))
	if n != 0 {
		logging.Fatal(errors.New("numbers of users do not match"))
	}
}


