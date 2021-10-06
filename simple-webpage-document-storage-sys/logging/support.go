package logging

import "go.uber.org/zap"

// S returns a struct used by the logger in the form of string : string
func S(s1 string, s2 string) SS {
	return SS{S1: s1, S2: s2}
}

// converts []SS to []zap.Field
func convert(sss []SS) []zap.Field {
	length := len(sss)
	if length == 0 {
		return nil
	}
	fields := make([]zap.Field, length)
	for i, ss := range sss {
		fields[i] = ss.convert()
	}
	return fields
}
