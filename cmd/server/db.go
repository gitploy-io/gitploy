package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"

	"contrib.go.opencensus.io/integrations/ocsql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-sql-driver/mysql"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	MySQLConnector struct {
		dsn string
	}
)

func (c MySQLConnector) Connect(context.Context) (driver.Conn, error) {
	return c.Driver().Open(c.dsn)
}

func (c MySQLConnector) Driver() driver.Driver {
	return ocsql.Wrap(
		mysql.MySQLDriver{},
		ocsql.WithAllTraceOptions(),
		ocsql.WithRowsClose(false),
		ocsql.WithRowsNext(false),
		ocsql.WithDisableErrSkip(true),
	)
}

func OpenDB(driver string, dsn string) (*ent.Client, error) {
	if driver == dialect.SQLite {
		return ent.Open(driver, dsn)
	}

	if driver == dialect.MySQL {
		db := sql.OpenDB(MySQLConnector{dsn})
		drv := entsql.OpenDB(dialect.MySQL, db)
		return ent.NewClient(ent.Driver(drv)), nil
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
