package controller

import (
	"simple-webpage-document-storage-sys/filesys"
)

// DefaultResponse the response for default-view GET request
type DefaultResponse struct {
	Top string `json:"top"`  // the id of the root directory
	Dirs filesys.Collection `json:"dirs"`
}

// FileRequest the request from user asking for an exact file; usually a counterpart of FileResponse
type FileRequest struct {
	User string `json:"user"`  // the username
	Fid string `json:"fid"`  // the file id
}

// FileResponse the response to the user asking for an exact file; usually a counterpart of FileRequest
type FileResponse struct {
	FileName string `json:"file_name"`  // the filename
	Content string `json:"content"`  // the file content
}
