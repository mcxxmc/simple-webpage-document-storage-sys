package test

import (
	"errors"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/logging"
	"simple-webpage-document-storage-sys/manager"
)

func ManagerDefaultView() {
	n := manager.NumberOfUsers()
	logging.InfoInt("number of users logged in: ", "number:", n)
	if n != 0 {
		logging.Fatal(errors.New("numbers of users do not match"))
	}

	manager.RegisterUser(common.DefaultUserId)
	n = manager.NumberOfUsers()
	logging.InfoInt("number of users logged in: ", "number:", n)
	if n != 1 {
		logging.Fatal(errors.New("numbers of users do not match"))
	}

	manager.UnregisterUser(common.DefaultUserId)
	n = manager.NumberOfUsers()
	logging.InfoInt("number of users logged in: ", "number:", n)
	if n != 0 {
		logging.Fatal(errors.New("numbers of users do not match"))
	}
}


