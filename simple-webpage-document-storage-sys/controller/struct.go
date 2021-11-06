package controller

import (
	"simple-webpage-document-storage-sys/filesys"
)

// DefaultResponse the response for default-view GET request
type DefaultResponse struct {
	Top string `json:"top"`  // the id of the root directory
	Dirs filesys.Collection `json:"dirs"`
}

// RequestFile the request from user asking for an exact file; usually a counterpart of FileResponse
type RequestFile struct {
	User string `json:"user"`  // the user id
	Fid string `json:"fid"`  // the file id
}

// FileResponse the response to the user asking for an exact file; usually a counterpart of RequestFile
type FileResponse struct {
	FileName string `json:"file_name"`  // the filename
	Content string `json:"content"`  // the file content
}

// RequestModifyFile the response from user to modify an exact file
type RequestModifyFile struct {
	User string `json:"user"`  // the user id
	Fid string `json:"fid"`  // the file id
	NewC string `json:"new_c"`  // the new content
}

// RequestRename the response from user to rename a file or a dir
type RequestRename struct {
	User    string `json:"user"`     // the user id
	ObjId   string `json:"obj_id"`   // the object id
	Dir     bool   `json:"dir"`      // whether the file type is a directory
	NewName string `json:"new_name"` // the new name
}

// RequestCreate the response from user to create a file or a dir
type RequestCreate struct {
	User string `json:"user"`  // the user id
	ObjId string `json:"obj_id"`  // the object id
	Dir bool `json:"dir"`  // whether the file type is a directory
	Name string `json:"name"` // the name of the new file or dir
	NewC string `json:"new_c"`  // the new content (only available for file)
	ParentId string `json:"parent_id"`  // the parent dir id
}

// RequestDelete the response from user to delete a file or a dir
type RequestDelete struct {
	User string `json:"user"`  // the user id
	ObjId string `json:"obj_id"`  // the object id
	Dir bool `json:"dir"`  // whether the file type is a directory
}

// RequestMove the response from user to move a file or a dir
type RequestMove struct {
	User string `json:"user"`  // the user id
	ObjId string `json:"obj_id"`  // the object id
	Dir bool `json:"dir"`  // whether the file type is a directory
	NewParentId string `json:"new_parent_id"`  // the new parent dir id
}
