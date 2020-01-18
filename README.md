# SQLite

```
# ensure gcc is installed
CGO_ENABLED=1 go get github.com/mattn/go-sqlite3
CGO_ENABLED=1 go install github.com/mattn/go-sqlite3
```

## Clean SQLite

```
go clean -i  github.com/mattn/go-sqlite3...
```


## GitHub

https://github.com/golang/go/issues/36025

# sql-migrate

https://github.com/rubenv/sql-migrat://github.com/rubenv/sql-migrate

# vfsgen

https://github.com/shurcooL/vfsgen


## Build in dev:

```
go build -tags=dev
```

```
go run -tags=dev main.go
```

That builds a binary that reads from the filesystem

## Package for prod

```
# Need -tags=dev so it can find the data package
go run -tags=dev generate_assets.go // or go generate
go run main.go
```

