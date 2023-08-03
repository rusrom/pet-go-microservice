## chi
Lightweight, idiomatic and composable router for building Go HTTP services  
https://go-chi.io/  
https://github.com/go-chi/chi

```shell
go get -u github.com/go-chi/chi/v5
```

`chi` comes equipped with an optional `middleware` package, providing a suite of standard net/http middlewares
https://github.com/go-chi/chi#middlewares

_**Extra middlewares & packages_**  
https://github.com/go-chi/chi#extra-middlewares--packages

https://github.com/go-chi/cors

```shell
go get github.com/go-chi/cors
```


## sql
https://pkg.go.dev/database/sql

Package `sql` provides a generic interface around SQL (or SQL-like) databases.  
The sql package must be used in conjunction with a database driver. See for a [list of drivers](https://golang.org/s/sqldrivers).
Drivers that do not support context cancellation will not return until after the query is completed.

For usage examples, see the [wiki page](https://golang.org/s/sqlwiki)


_**Tutorial: Accessing a relational database**_  
https://go.dev/doc/tutorial/database-access

_**SQL database drivers**_  
The `database/sql` and `database/sql/driver` packages are designed for using databases from Go and implementing database drivers, respectively.  
https://github.com/golang/go/wiki/SQLDrivers

Install pgx - PostgreSQL Driver and Toolkit  
https://github.com/jackc/pgx

```shell
go get github.com/jackc/pgconn
go get github.com/jackc/pgx/v4
go get github.com/jackc/pgx/v4/stdlib
```
