package test

import (
	"errors"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/logging"
	"simple-webpage-document-storage-sys/manager"
)

// ManagerDefaultView tests using the default user
func ManagerDefaultView() {
	userId := common.DefaultUserId
	dirId1 := "1001"
	dirId2 := "1002"

	newDirId1 := "dir1"
	newFileId1 := "file1"

	n := manager.NumberOfUsers()
	if n != 0 {
		logging.Fatal(errors.New("numbers of users do not match"))
	}

	// Test: new register
	manager.RegisterUser(userId)
	n = manager.NumberOfUsers()
	if n != 1 {
		logging.Fatal(errors.New("numbers of users do not match"))
	}

	// Test dir & txt:
	// create new dir
	testDirName := "TestDir1"
	yes := manager.CreateDir(userId, newDirId1, testDirName, common.RootId)
	if yes == false {
		logging.Fatal(errors.New("fails to create a new directory"))
	}

	// create new txt file
	testFileName := "TestFile1"
	testFileContent := "ManagerDefaultView(): test file 1."
	yes = manager.CreateTxt(userId, newFileId1, testFileName,
		testFileContent, newDirId1)

	if yes == false {
		logging.Fatal(errors.New("fails to create txt"))
	}

	// fetch txt
	name, ctx := manager.FetchTxt(userId, newFileId1)
	if name != testFileName || ctx != testFileContent {
		logging.Fatal(errors.New("file name or content does not match"))
	}

	// rename the dir
	testDirNameNew := "NewTestDir1"
	yes = manager.RenameDir(userId, newDirId1, testDirNameNew)
	if yes == false {
		logging.Fatal(errors.New("fails to rename directory"))
	}

	// move the dir
	yes = manager.MoveDir(userId, newDirId1, dirId1)
	if yes == false {
		logging.Fatal(errors.New("fails to move directory"))
	}

	// try to delete that directory; it should fail here because it is not empty
	yes = manager.DeleteDir(userId, newDirId1)
	if yes == true {
		logging.Fatal(errors.New("a mistake occurs as a non-empty directory is deleted"))
	}

	// move txt
	yes = manager.MoveTxt(common.DefaultUserId, newFileId1, dirId2)

	if yes == false {
		logging.Fatal(errors.New("fails to move txt"))
	}

	// again, try to delete that directory; it should work now
	yes = manager.DeleteDir(userId, newDirId1)
	if yes == false {
		logging.Fatal(errors.New("fails to delete directory"))
	}

	// modify txt & changes txt name
	testFileNameNew := "NewTestFile1"
	testFileContentNew := "ManagerDefaultView(): new test file 1."

	yes = manager.ModifyTxt(common.DefaultUserId, newFileId1, testFileContentNew)
	if yes == false {
		logging.Fatal(errors.New("fails to modify txt"))
	}

	yes = manager.RenameTxt(common.DefaultUserId, newFileId1, testFileNameNew)
	if yes == false {
		logging.Fatal(errors.New("fails to rename txt"))
	}

	name, ctx = manager.FetchTxt(common.DefaultUserId, newFileId1)
	if name != testFileNameNew || ctx != testFileContentNew {
		logging.Fatal(errors.New("file name or content does not match"))
	}

	// delete txt
	yes = manager.DeleteTxt(common.DefaultUserId, newFileId1)

	if yes == false {
		logging.Fatal(errors.New("fails to delete txt"))
	}

	// Test: log out
	manager.UnregisterUser(common.DefaultUserId)
	n = manager.NumberOfUsers()
	if n != 0 {
		logging.Fatal(errors.New("numbers of users do not match"))
	}
}


