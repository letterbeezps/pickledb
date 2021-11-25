# PickleDB-go

A lightweight and simple key-value store written in Go, inspired by [Python's PickleDB](https://github.com/patx/pickledb) and [PickleDB-rs](https://github.com/seladb/pickledb-rs).

## exmaple

```go
package main

import (
 "github.com/letterbeezps/pickledb"
)

func main() {
 db := pickledb.Load("test.db")

 db.Set("letter", "letter_value")
 err := db.Dump()
 if err != nil {
  panic(err)
 }
}
```

```shell
go run main.go
```

Create test.db in the current folder
