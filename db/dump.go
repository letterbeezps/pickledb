package db

import (
	"encoding/gob"
	"os"
	"sync"
)

type dumpData struct {
	Data map[string]*value
}

func newEmptyDump() *dumpData {
	return &dumpData{}
}

func newDump(db *Pickledb) *dumpData {
	return &dumpData{
		Data: db.data,
	}
}

func (d *dumpData) dump(location string) error {
	file, err := os.OpenFile(location, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gob.NewEncoder(file).Encode(d)

	if err != nil {
		file.Close()
		os.Remove(location)
		return err
	}

	file.Close()
	return nil
}

func (d *dumpData) load(location string) (*Pickledb, error) {
	file, err := os.Open(location)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	if err = gob.NewDecoder(file).Decode(d); err != nil {
		return nil, err
	}

	return &Pickledb{
		data: d.Data,
		lock: &sync.RWMutex{},
	}, nil
}
