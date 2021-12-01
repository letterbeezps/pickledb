package pickledb

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/letterbeezps/pickledb/db"
	"github.com/letterbeezps/pickledb/global"
)

func Load(location string, autoDump bool) *db.Pickledb {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	global.StoreLocation = fmt.Sprintf("%s/%s", dir, location)
	global.AutoDump = autoDump

	gob.Register([]interface{}{})
	gob.Register(map[string]interface{}{})

	newDB := db.LoadPickleDb()

	return newDB
}
