package test

import (
	"errors"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/logging"
	"simple-webpage-document-storage-sys/manager"
)

const userId = common.DefaultUserId
const dirId1 = "1001"
const dirId2 = "1002"

const newDirId1 = "dir1"
const newDirId2 = "dir2"
const newFileId1 = "file1"
const newFileId2 = "file2"

// ManagerDefaultView tests using the default user
func ManagerDefaultView() {

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
	// PART1: without saving
	// create new dir
	testDirName1 := "TestDir1"
	yes := manager.CreateDir(userId, newDirId1, testDirName1, common.RootId)
	if yes == false {
		logging.Fatal(errors.New("fails to create a new directory"))
	}

	// create new txt file
	testFileName1 := "TestFile1"
	testFileContent1 := "ManagerDefaultView(): test file 1."
	yes = manager.CreateTxt(userId, newFileId1, testFileName1,
		testFileContent1, newDirId1)

	if yes == false {
		logging.Fatal(errors.New("fails to create txt"))
	}

	// fetch txt
	name, ctx := manager.FetchTxt(userId, newFileId1)
	if name != testFileName1 || ctx != testFileContent1 {
		logging.Fatal(errors.New("file name or content does not match"))
	}

	// rename the dir
	testDirName2 := "TestDir2"
	yes = manager.RenameDir(userId, newDirId1, testDirName2)
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
	testFileName2 := "TestFile2"
	testFileContent2 := "ManagerDefaultView(): test file 2."

	yes = manager.ModifyTxt(common.DefaultUserId, newFileId1, testFileContent2)
	if yes == false {
		logging.Fatal(errors.New("fails to modify txt"))
	}

	yes = manager.RenameTxt(common.DefaultUserId, newFileId1, testFileName2)
	if yes == false {
		logging.Fatal(errors.New("fails to rename txt"))
	}

	name, ctx = manager.FetchTxt(common.DefaultUserId, newFileId1)
	if name != testFileName2 || ctx != testFileContent2 {
		logging.Fatal(errors.New("file name or content does not match"))
	}

	// delete txt
	yes = manager.DeleteTxt(common.DefaultUserId, newFileId1)

	if yes == false {
		logging.Fatal(errors.New("fails to delete txt"))
	}

	//PART2: with saving
	testDirName3 := "TestDir3"
	yes = manager.CreateDir(userId, newDirId2, testDirName3, dirId2)
	if yes == false {
		logging.Fatal(errors.New("fails to create a new directory"))
	}
	testFileName3 := "TestFile3"
	testFileContent3 := "ManagerDefaultView(): test file 3."
	yes = manager.CreateTxt(userId, newFileId2, testFileName3,
		testFileContent3, newDirId2)
	if yes == false {
		logging.Fatal(errors.New("fails to create txt"))
	}
	manager.SaveUserCollection(userId)


	// Test: log out
	manager.UnregisterUser(common.DefaultUserId)
	n = manager.NumberOfUsers()
	if n != 0 {
		logging.Fatal(errors.New("numbers of users do not match"))
	}
}


