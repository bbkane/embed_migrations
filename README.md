https://github.com/rubenv/sql-migrat://github.com/rubenv/sql-migrate


```
# ensure gcc is installed
go get github.com/mattn/go-sqlite3
CGO_ENABLED=1 go install github.com/mattn/go-sqlite3
```

Get the packr binary

go get -u github.com/gobuffalo/packr/v2/packr2

My God this gets a lot of dependencies!

It also increases compilation :( - .8s for no migrations!
