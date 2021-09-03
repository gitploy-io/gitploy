# GITPLOY_STORE_SOURCE

String value provides the source of store. Configures the database connection string. The default value is the path of the embedded sqlite3 database file.

```
GITPLOY_STORE_SOURCE=file:/data/sqlite3.db?cache=shared&_fk=1
```

Example mysql connection string (below). See the official driver [documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name) for more connection string details.

```
GITPLOY_STORE_SOURCE=root:password@tcp(1.2.3.4:3306)/gitploy?parseTime=true
```

Example postgres connection string (below). See the official driver [documentation](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING) for more connection string details.

```
GITPLOY_STORE_SOURCE=postgres://root:password@1.2.3.4:5432/gitploy?sslmode=disable
```