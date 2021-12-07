# PickleDB-go

A lightweight and simple key-value store written in Go, inspired by [Python's PickleDB](https://github.com/patx/pickledb) and [PickleDB-rs](https://github.com/seladb/pickledb-rs).

When I used go to refactor the python project of my former colleague, I tried to find an alternative to [pickleDB](https://github.com/patx/pickledb) in [go.dev](https://pkg.go.dev/), but it failed. Fortunately, the source code of [pickleDB](https://github.com/patx/pickledb) is not complicated, so I try to develop a pickleDB-go by myself.

The core of the project is to user map[string]interface{} represent arbitray data.

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
