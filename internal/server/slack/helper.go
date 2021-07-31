package slack

import (
	"math/rand"
	"strconv"
)

func randstr() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
