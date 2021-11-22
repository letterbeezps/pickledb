package pickledb

import "github.com/letterbeezps/pickledb/db"

func Load() *db.Pickledb {
	newDB := db.NewPickleDb()

	return newDB
}
