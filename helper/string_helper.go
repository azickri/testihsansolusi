package helper

import (
	"math/rand"
)

func RandomNumber(length int) string {
	letterBytes := "1234567890"

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}
