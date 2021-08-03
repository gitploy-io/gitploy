package slack

import (
	"strconv"
)

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
