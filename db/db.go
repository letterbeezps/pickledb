package db

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/letterbeezps/pickledb/global"
	"github.com/letterbeezps/pickledb/util"
)

type Pickledb struct {
	data map[string]*value

	lock *sync.RWMutex
}

func NewPickleDb() *Pickledb {
	return &Pickledb{
		data: make(map[string]*value, 256),
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
		// return NewPickleDb()
		panic(err)
	}

	return newdb
}

func (db *Pickledb) Get(key string) (interface{}, bool) {
	value, ok := db.get(key)
	if !ok {
		return nil, ok
	}
	return value.Data, ok
}

func (db *Pickledb) GetAll() []string {
	return db.getall()
}

func (db *Pickledb) Set(key string, v interface{}) {
	convertValue := util.Convert(v)
	val := newValue(convertValue, "N")
	db.set(key, val)

}

func (db *Pickledb) Rem(key string) {
	if _, ok := db.get(key); ok {
		db.del(key)
	}
}

// list operation

// create a list
func (db *Pickledb) ListCreate(key string) bool {
	val := newValue([]interface{}{}, "L")
	db.set(key, val)

	return true
}

// add a value to list
func (db *Pickledb) ListAdd(key string, v interface{}) {
	value, ok := db.get(key)

	if !ok {
		fmt.Printf("%s not exist", key)
		return
	}
	oldList := value.Data
	list := oldList.([]interface{})
	convertValue := util.Convert(v)
	list = append(list, convertValue)
	value.Data = list
	db.set(key, value)

}

// extend a list with a array
func (db *Pickledb) ListExtend(key string, v interface{}) {
	value, ok := db.get(key)

	if !ok {
		fmt.Printf("%s not exist", key)
		return
	}
	oldList := value.Data
	list := oldList.([]interface{})

	convertV := util.Convert(v)
	vList := convertV.([]interface{})

	list = append(list, vList...)
	value.Data = list

	db.set(key, value)
}

// dump operation
func (db *Pickledb) Dump() {

	if err := db.dump(global.StoreLocation); err != nil {
		panic(err)
	}
}

// dict operation

// create a dict
func (db *Pickledb) DictCreate(key string) bool {
	val := newValue(map[string]interface{}{}, "D")
	db.set(key, val)

	return true
}

// add a key-value to a dict
func (db *Pickledb) DictAdd(key string, dictKey string, dictValue interface{}) {
	value, ok := db.get(key)
	if !ok {
		fmt.Printf("%s not exist", key)
		return
	}
	oldDict := value.Data
	dict, ok := oldDict.(map[string]interface{})
	if !ok {
		fmt.Printf("value of %s not a dict", key)
		return
	}
	convertValue := util.Convert(dictValue)
	dict[dictKey] = convertValue
	value.Data = dict
	db.set(key, value)

}

func (db *Pickledb) DictDelete(key, dictKey string) {
	value, ok := db.get(key)
	if !ok {
		fmt.Printf("%s not exist", key)
		return
	}
	oldDict := value.Data
	dict, ok := oldDict.(map[string]interface{})
	if !ok {
		fmt.Printf("value of %s not a dict", key)
		return
	}
	delete(dict, dictKey)
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

func (db *Pickledb) get(key string) (*value, bool) {
	db.lock.RLock()
	defer db.lock.RUnlock()
	value, ok := db.data[key]
	return value, ok
}

func (db *Pickledb) getall() []string {
	keys := make([]string, 0, len(db.data))
	for k := range db.data {
		keys = append(keys, k)
	}
	return keys
}

func (db *Pickledb) set(key string, value *value) {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.data[key] = value
}

func (db *Pickledb) del(key string) {
	db.lock.Lock()
	defer db.lock.Unlock()

	delete(db.data, key)
}

func (db *Pickledb) dump(location string) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	return newDump(db).dump(location)
}
