package schema

import (
	"math/rand"
	"time"
)

func generateHash() string {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 32)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func nowUTC() time.Time {
	return time.Now().UTC()
}

func zeroTime() time.Time {
	return time.Time{}
}
