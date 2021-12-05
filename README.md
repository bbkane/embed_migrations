# embed_migrations

`embed_migrations` is how I finally figured out how to embed sql migration
files into a go app and use them from the binary.

## Update: 2021-12-06

I've updated to use go's new `embed` library/compiler directive and modernc.org/sqlite . No need for CGO or go generate any more!!

Usage:

```
go run .
```

The rest of this README is no longer needed!

## Dependencies

[go-sqlite3](https://github.com/mattn/go-sqlite3) is a CGO app. There's
currently an [issue](https://github.com/golang/go/issues/36025) on GitHub about
the linker errors that building CGO apps creates.

```
# ensure gcc is installed
CGO_ENABLED=1 go install github.com/mattn/go-sqlite3

# clean in prep for rebuild if necessary
go clean -i  github.com/mattn/go-sqlite3...
```

- [sql-migrate](https://github.com/rubenv/sql-migrate/) can read from an `http.FileSystem`
- [vfsgen](https://github.com/shurcooL/vfsgen) can provide an `http.FileSystem`


## Read Migrations From Directory

```
go run -tags=dev main.go
```

## Embed Migrations Directory Into Code

```
go generate
go run main.go
```

## Things that didn't work

- [goose](https://github.com/pressly/goose), in its library form, can only read migrations from files
- [migrate](https://github.com/golang-migrate/migrate) is not tested against SQLite3
- [packr](https://github.com/gobuffalo/packr/tree/master/v2) never did find my files.
- [statik](https://github.com/rakyll/statik) doesn't support directories (and doesn't mention this in the README.md, only in the code)
- `sql-migrate` doesn't mention `http.FileSystem` in the README.md. I have a [pull request](https://github.com/rubenv/sql-migrate/pull/162) open for this
- `vfsgen`'s `go generate` [example](https://github.com/shurcooL/vfsgen#go-generate-usage) didn't work for me. This project is the tweaks I made to get it working

## Problems

`vfsgen` doesn't seem to show up in `go.mod`.
