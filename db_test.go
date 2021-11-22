package pickledb

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -run TestSetGet
func TestSetGet(t *testing.T) {
	testDB := Load()

	testDB.Set("zp", "zp_value")

	value, ok := testDB.Get("zp")

	assert.Equal(t, ok, true)

	assert.Equal(t, string(value), "zp_value")

	fmt.Println(string(value))
}

// go test -v -run TestSetGetJson

type testJson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestSetGetJson(t *testing.T) {
	testDB := Load()

	fakeJson := testJson{
		Name: "zp",
		Age:  24,
	}

	testDB.Set("zp", fakeJson)

	value, ok := testDB.Get("zp")

	assert.Equal(t, ok, true)

	outJson := testJson{}

	json.Unmarshal(value, &outJson)

	fmt.Println(outJson)
}

// go test -v -run TestDump
func TestDump(t *testing.T) {
	testDB := Load()

	testDB.Set("zp", "zp_value")

	testDB.Set("zz", "zz_value")

	testDB.Dump("test.db")
}
