package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gitploy-io/gitploy/extent"
)

const (
	year = 365 * 24 * time.Hour
)

var (
	limit = flag.Int("limit", 0, "Member count")
)

func main() {
	flag.Parse()

	if *limit == 0 {
		log.Fatal("Set the limit.")
	}

	d := &extent.SigningData{
		MemberLimit: *limit,
		ExpiredAt:   time.Now().Add(year),
	}

	j, err := json.Marshal(d)
	if err != nil {
		log.Fatalf("It has failed to marshal.")
	}

	fmt.Print(string(j))
}
