package pickledb

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -run TestSetGet
func TestSetGet(t *testing.T) {
	testDB := Load("test.db")

	testDB.Set("zp", "zp_value")

	value, ok := testDB.Get("zp")

	assert.Equal(t, ok, true)

	assert.Equal(t, string(value), "zp_value")

	fmt.Println(string(value))
}

// go test -v -run TestJson

type testJson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJson(t *testing.T) {
	testDB := Load("test.db")

	fakeJson := testJson{
		Name: "zp",
		Age:  24,
	}

	testDB.Set("zp", fakeJson)
	testDB.Set("zz", "zz_value")

	value, ok := testDB.Get("zp")

	assert.Equal(t, ok, true)

	outJson := testJson{}

	fmt.Println(string(value))

	json.Unmarshal(value, &outJson)

	fmt.Println(outJson)

	testDB.Dump()

}

// go test -v -run TestDump
func TestDump(t *testing.T) {
	testDB := Load("test.db")

	testDB.Set("zp", "zp_value")

	testDB.Set("zz", "zz_value")

	testDB.Dump()
}

// go test -v -run TestLoad
func TestLoad(t *testing.T) {
	testDB := Load("test.db")

	// testDB.Set("zp", "zp_value")

	value, ok := testDB.Get("zp")

	assert.Equal(t, ok, true)

	outJson := testJson{}

	fmt.Println(string(value))

	json.Unmarshal(value, &outJson)

	fmt.Println(outJson)
}

// go test -v -run TestGetAll
func TestGetAll(t *testing.T) {
	testDB := Load("test1.db")

	testDB.Set("zp", "zp_value")

	testDB.Set("zz", "zz_value")

	keys := testDB.GetAll()

	fmt.Println(keys)
}
