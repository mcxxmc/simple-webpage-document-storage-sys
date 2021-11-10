package common

import (
	"math/rand"
	"time"
)

func GenerateRandomId(n int) string {
	id := make([]byte, n)
	length := len(LetterBytes)
	rand.Seed(time.Now().UnixNano())  // to avoid same random ids
	for i := range id {
		id[i] = LetterBytes[rand.Intn(length)]
	}
	return string(id)
}
