package filesys

type UserInfo struct {
	Uid string `json:"uid"`
	Profile string `json:"profile"`
}

type IndexesOfUsers map[string]*UserInfo

type UserDir struct {
	ID string `json:"id"`  // the id, should be consistent with the key of the map
	Dir bool `json:"dir"`  // if it is a directory or is linked with a real txt file
	Name string `json:"name"`
	Level int `json:"level"`  // the level of that directory
	Link []string `json:"link"`  //TODO: for privacy and safety, add encoding and decoding for link if it points to a real file
}

type UserDirs map[string]*UserDir  // passed by pointer
