// +build oss

package main

import (
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/gitploy-io/gitploy/model/ent"
)

func OpenDB(driver string, dsn string) (*ent.Client, error) {
	if driver != dialect.SQLite {
		return nil, fmt.Errorf("The community edition support sqlite only.")
	}

	return ent.Open(driver, dsn)
}
