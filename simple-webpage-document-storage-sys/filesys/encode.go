package filesys

import (
	"errors"
	"io/ioutil"
	"os"
	"simple-webpage-document-storage-sys/logging"
)

// RewriteTxt modifies a txt if it exists; creates a new one if not.
func RewriteTxt(path string, newContent string) bool {
	err := ioutil.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		logging.Error(err)
		return false
	}
	return true
}

// DeleteTxt deletes a txt if it exists;
//
// args: the userId. Should not be empty if it is not in the path.
func DeleteTxt(path string, args...string) bool {
	if checkPath(path, args) == false {
		logging.Error(errors.New("invalid path"))
		return false
	}
	err := os.Remove(path)
	if err != nil {
		logging.Error(err)
		return false
	}
	return true
}
