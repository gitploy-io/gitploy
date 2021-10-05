# Database

Gitploy requires the use of a database backend for persistence. Gitploy uses an embedded SQLite database by default. This article provides alternate databases: MySQL and Postgres.

## MySQL

Gitploy supports mysql `5.6` and higher as the database engine. The below example demonstrates mysql database configuration. See the official driver [documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name) for more connection string details.

```
GITPLOY_STORE_DRIVER=mysql
GITPLOY_STORE_SOURCE=root:password@tcp(1.2.3.4:3306)/gitploy?parseTime=true
```

### *⚠️ WARNING*

You have to setting the SQL mode by removing `NO_ZERO_DATE` from the default setting to enable the zero value for the `timestamp` type. 

```
sql-mode='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION'
```

To set the SQL mode, you can check the official [document](https://dev.mysql.com/doc/refman/8.0/en/sql-mode.html).

## Postgres

Gitploy supports postgres on the following 4 versions: `10`, `11`, `12` and `13`. The below example demonstrates postgres database configuration. See the official driver [documentation](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING) for more connection string details.

```
GITPLOY_STORE_DRIVER=postgres
GITPLOY_STORE_SOURCE=root:password@1.2.3.4:5432/gitploy?sslmode=disable
```
