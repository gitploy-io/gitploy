package repos

import "strconv"

const (
	defaultQueryPage    = "1"
	defaultQueryPerPage = "30"
)

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
