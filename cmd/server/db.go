// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package main

import (
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/gitploy-io/gitploy/ent"
)

func OpenDB(driver string, dsn string) (*ent.Client, error) {
	if driver == dialect.SQLite || driver == dialect.MySQL {
		return ent.Open(driver, dsn)
	}

	if driver == dialect.Postgres {
		db, err := sql.Open("pgx", dsn)
		if err != nil {
			return nil, err
		}

		drv := entsql.OpenDB(dialect.Postgres, db)
		return ent.NewClient(ent.Driver(drv)), nil
	}

	return nil, fmt.Errorf("The driver have to be one of them: sqlite3, mysql, or postgres.")
}
