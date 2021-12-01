# PickleDB-go

A lightweight and simple key-value store written in Go, inspired by [Python's PickleDB](https://github.com/patx/pickledb) and [PickleDB-rs](https://github.com/seladb/pickledb-rs).

## exmaple

```go
package main

import (
    "fmt"

    "github.com/letterbeezps/pickledb"
)

func main() {
    Dump()

    Load()
}

func Dump() {
    db := pickledb.Load("test.db", false)

    db.Set("letter", "letter_value")
    db.Dump()
}

func Load() {
    db := pickledb.Load("test.db", false)

    v, _ := db.Get("letter")
    fmt.Println(v)
}

```

```shell
go run main.go
```

Create test.db in the current folder
