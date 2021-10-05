package filesys

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"simple-webpage-document-storage-sys/logging"
)

// LoadUserIndexes loads the user-index JSON from disk; should be called only once
func LoadUserIndexes(path string) *IndexesOfUsers {
	jsonFile, err := os.Open(path)
	logging.ConditionalLogError(err)
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	logging.ConditionalLogError(err)

	r := &IndexesOfUsers{}
	err = json.Unmarshal(bytes, r)
	logging.ConditionalLogError(err)
	return r
}

// LoadUserDirs loads the JSON file that contains the info about the directories owned by a certain user
func LoadUserDirs(path string) *Collection {
	jsonFile, err := os.Open(path)
	logging.ConditionalLogError(err)
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	logging.ConditionalLogError(err)

	r := &Collection{}
	err = json.Unmarshal(bytes, r)
	logging.ConditionalLogError(err)
	return r
}

// OpenTxt opens the txt file and return all the bytes as string
func OpenTxt(path string) string {
	f, err := os.Open(path)
	logging.ConditionalLogError(err)
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		logging.Error(err)
		return "Error in reading file. Please contact our team for more details."
	}
	return string(bytes)
}
