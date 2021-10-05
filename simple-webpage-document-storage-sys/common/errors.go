package common

// ErrorCollectionNotFound returns the error message when the collection for a user id is not found
func ErrorCollectionNotFound(userId string) string {
	return "collection for the user id " + userId + " is not found"
}

// ErrorFileIdAlreadyExists returns the error message when the file id is duplicated
func ErrorFileIdAlreadyExists(fileId string) string {
	return "the file id " + fileId + " already exists"
}

// ErrorInvalidNewParentId returns the error msg if the new parent id is invalid
func ErrorInvalidNewParentId(newParentId string) string {
	return "the new parent id " + newParentId + " does not exist"
}
