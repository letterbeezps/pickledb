package db

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	pickleUitl "github.com/letterbeezps/pickledb/util"
)

type Pickledb struct {
	data map[string][]byte

	lock *sync.RWMutex
}

func NewPickleDb() *Pickledb {
	return &Pickledb{
		data: make(map[string][]byte, 256),
		lock: &sync.RWMutex{},
	}
}

func (db *Pickledb) Get(key string) ([]byte, bool) {
	return db.get(key)
}

func (db *Pickledb) Set(key string, value interface{}) {
	if val, ok := value.(string); ok {
		db.set(key, []byte(val))
	} else {
		val, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}

		db.set(key, val)
	}

}

func (db *Pickledb) Dump(location string) error {

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	newLocation := fmt.Sprintf("%s/%s", dir, location)

	return db.dump(newLocation)
}

/****************************************************/

func (db *Pickledb) get(key string) ([]byte, bool) {
	db.lock.RLock()
	defer db.lock.RUnlock()
	value, ok := db.data[key]
	return value, ok
}

func (db *Pickledb) set(key string, value []byte) {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.data[key] = pickleUitl.Copy(value)
}

func (db *Pickledb) dump(location string) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	return newDump(db).dump(location)
}
