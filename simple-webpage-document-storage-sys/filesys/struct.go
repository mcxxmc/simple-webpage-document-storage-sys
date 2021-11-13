package filesys

// UserInfo the basic user info
type UserInfo struct {
	Name string `json:"name"`
	Uid string `json:"uid"`
	Password string `json:"password"`
	Profile string `json:"profile"`  // the path where the user profile is stored
}

// IndexesOfUsers maintains the basic info of all users;
//
// map structure = uid : *UserInfo
type IndexesOfUsers map[string]*UserInfo

// Image keeps the info of a directory or a file owned by the user;
//
// for any file, Children only contains 1 element which is the exact path of that file.
type Image struct {
	ID       string   `json:"id"`  // the id, should be consistent with the key of the map
	Dir      bool     `json:"dir"`  // if it is a directory or is linked with a real txt file
	Name     string   `json:"name"`
	Level    int      `json:"level"` // the level of that directory
	Children []string `json:"children"` //TODO: for privacy and safety, add encoding and decoding for link if it points to a real file
	Parent   string   `json:"parent"`
}

// Collection keeps the info of all directories and files owned by one user;
//
// map structure = Image id : *Image
type Collection map[string]*Image // passed by pointer

// Profile a new profile for a new user containing only the root directory
type Profile struct {
	ROOT struct {
		Id       string        `json:"id"`
		Dir      bool          `json:"dir"`
		Name     string        `json:"name"`
		Level    int           `json:"level"`
		Children []string      `json:"children"`
		Parent   string        `json:"parent"`
	} `json:"ROOT"`
}
