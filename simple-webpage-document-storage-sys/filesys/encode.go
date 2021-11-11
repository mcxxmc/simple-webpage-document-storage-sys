package filesys

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/logging"
)

// checks if the path is valid
func checkPath(path string, args []string) bool {
	userTxtPrefix := common.Path_txt_prefix
	for _, str := range args {
		userTxtPrefix = userTxtPrefix + str + common.Slash
	}
	if path[:len(userTxtPrefix)] != userTxtPrefix {
		return false
	}
	return true
}

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

// SaveUserCollection saves the user collection (as a JSON file) to disk
func SaveUserCollection(path string, modified *Collection) error {
	byteVal, err := json.Marshal(modified)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, byteVal, 0644)
	return err
}

// SaveIndexesOfUsers saves the indexes of users to the disk
func SaveIndexesOfUsers(path string, indexes *IndexesOfUsers) error {
	byteVal, err := json.Marshal(indexes)
	if err != nil {
		return nil
	}
	err = ioutil.WriteFile(path, byteVal, 0644)
	return err
}

// CreateNewUserProfile creates a new profile and save it to disk
func CreateNewUserProfile(path string) error {
	newProfile := &Profile{
		ROOT: struct {
			Id       string        `json:"id"`
			Dir      bool          `json:"dir"`
			Name     string        `json:"name"`
			Level    int           `json:"level"`
			Children []string      `json:"children"`
			Parent   string        `json:"parent"`
		}(struct {
			Id       string
			Dir      bool
			Name     string
			Level    int
			Children []string
			Parent   string
		}{Id: "ROOT", Dir: true, Name: "ROOT", Level: 0, Children: make([]string, 0), Parent: "NONE"}),
	}
	byteVal, err := json.Marshal(newProfile)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, byteVal, 0644)
	return err
}
