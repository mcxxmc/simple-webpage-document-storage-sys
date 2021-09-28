package file_sys

type IndexesOfUsers map[string]string

type UserDir struct {
	ID string `json:"id"`  // the id, should be consistent with the key of the map
	Dir bool `json:"dir"`  // if it is a directory or is linked with a real txt file
	Name string `json:"name"`
	Level int `json:"level"`  // the level of that directory
	Link []string `json:"link"`
}

type UserDirs map[string]UserDir
