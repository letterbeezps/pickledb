package db

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/letterbeezps/pickledb/global"
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

func LoadPickleDb() *Pickledb {

	if _, err := os.Stat(global.StoreLocation); err != nil {
		if os.IsNotExist(err) {
			return NewPickleDb()
		} else {
			panic(err)
		}
	}

	newdb, err := newEmptyDump().load(global.StoreLocation)
	if err != nil {
		return NewPickleDb()
	}

	return newdb
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

func (db *Pickledb) Dump() error {

	return db.dump(global.StoreLocation)
}

// Think before you act when use this function
func (db *Pickledb) AutoDumpByTime() {
	go func() {
		ticker := time.NewTicker(time.Duration(global.DumpDuration) * time.Minute)
		for {
			select {
			case <-ticker.C:
				db.Dump()
			}
		}
	}()
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
