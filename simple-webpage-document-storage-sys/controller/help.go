package controller

import (
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/filesys"
)

// DefaultResponse the response for default-view GET request
type DefaultResponse struct {
	Tops []string `json:"tops"`  // the ids of level-1 directories
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

func wrapUpUserDirs(dirs *filesys.Collection) *DefaultResponse {
	tops := make([]string, 10)
	for k, dir := range *dirs {
		if dir.Level == common.RootLevel {
			tops = append(tops, k)
		}
	}
	return &DefaultResponse{Tops: tops, Dirs: *dirs}
}
