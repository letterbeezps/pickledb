package pickledb

import (
	"fmt"
	"os"

	"github.com/letterbeezps/pickledb/db"
	"github.com/letterbeezps/pickledb/global"
)

func Load(location string) *db.Pickledb {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	global.StoreLocation = fmt.Sprintf("%s/%s", dir, location)

	newDB := db.LoadPickleDb()

	return newDB
}
