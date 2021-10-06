package manager

// removes a string from an array of strings and returns the new string (not the original one)
func removeString(strings []string, toRemove string) []string {
	index := -1
	for i, str := range strings {
		if str == toRemove {
			index = i
			break
		}
	}
	if index == -1 {
		return strings
	} else {
		return append(append([]string{}, strings[:index]...), strings[index + 1:]...)
	}
}
