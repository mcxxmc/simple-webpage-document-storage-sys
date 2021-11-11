package controller

import (
	"simple-webpage-document-storage-sys/filesys"
)

// HierarchyResponse for generating the hierarchy view
type HierarchyResponse struct {
	Ok bool `json:"ok"`
	Top string `json:"top"`  // the id of the root directory
	Dirs filesys.Collection `json:"dirs"`
}

// CommonResponse for making a msg response containing whether the previous request is successful
type CommonResponse struct {
	Ok bool `json:"ok"`
	Msg string `json:"msg"`
}

// RequestFile the request from user asking for an exact file; usually a counterpart of FileResponse
type RequestFile struct {
	Fid string `json:"fid"`  // the file id
}

// FileResponse the response to the user asking for an exact file; usually a counterpart of RequestFile
type FileResponse struct {
	Ok bool `json:"ok"`
	FileName string `json:"file_name"`  // the filename
	Content string `json:"content"`  // the file content
}

// RequestModifyFile the request from user to modify an exact file
type RequestModifyFile struct {
	Fid string `json:"fid"`  // the file id
	NewC string `json:"new_c"`  // the new content
}

// RequestRename the request from user to rename a file or a dir
type RequestRename struct {
	ObjId   string `json:"obj_id"`   // the object id
	Dir     bool   `json:"dir"`      // whether the file type is a directory
	NewName string `json:"new_name"` // the new name
}

// RequestCreate the request from user to create a file or a dir
type RequestCreate struct {
	Dir bool `json:"dir"`  // whether the file type is a directory
	Name string `json:"name"` // the name of the new file or dir
	NewC string `json:"new_c"`  // the new content (only available for file)
	ParentId string `json:"parent_id"`  // the parent dir id
}

// RequestDelete the request from user to delete a file or a dir
type RequestDelete struct {
	ObjId string `json:"obj_id"`  // the object id
	Dir bool `json:"dir"`  // whether the file type is a directory
}

// RequestMove the request from user to move a file or a dir
type RequestMove struct {
	ObjId string `json:"obj_id"`  // the object id
	Dir bool `json:"dir"`  // whether the file type is a directory
	NewParentId string `json:"new_parent_id"`  // the new parent dir id
}

// RequestLogin the request from user to login
type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse the response to users who want to log in
type LoginResponse struct {
	Ok bool `json:"ok"`
	Token string `json:"token"`
}
