package pickledb

import (
	"fmt"
	"testing"

	"github.com/letterbeezps/pickledb/util"
	"github.com/stretchr/testify/assert"
)

type testJson struct {
	Name    string
	Age     int
	Friends []string
}

// go test -v -run TestSetGet
func TestSetGet(t *testing.T) {
	testDB := Load("test.db", false)

	testDB.Set("zp", "zp_value")

	value, ok := testDB.Get("zp")

	assert.Equal(t, ok, true)

	assert.Equal(t, value, "zp_value")

	fmt.Println(value)

	fakeJson := testJson{
		Name:    "zp",
		Age:     24,
		Friends: []string{"z1", "z2", "z3"},
	}

	fakeValue := util.Convert(fakeJson)

	testDB.Set("fakeZp", fakeValue)

	fmt.Println(fakeValue)
}

// go test -v -run TestRem
func TestRem(t *testing.T) {
	testDB := Load("test.db", false)

	testDB.Set("zp", "zp_value")

	value, ok := testDB.Get("zp")

	assert.Equal(t, ok, true)

	assert.Equal(t, value, "zp_value")

	testDB.Rem("zp")

	_, ok = testDB.Get("zp")

	assert.Equal(t, ok, false)

	fmt.Println(value)
}

// go test -v -run TestDump
func TestDump(t *testing.T) {
	testDB := Load("test.db", false)

	testDB.Set("zp", "zp_value")

	testDB.Set("zz", "zz_value")

	fakeJson := testJson{
		Name:    "zp",
		Age:     24,
		Friends: []string{"z1", "z2", "z3"},
	}

	fakeValue := util.Convert(fakeJson)

	testDB.Set("fakeZp", fakeValue)

	testDB.Dump()
}

// go test -v -run TestLoad
func TestLoad(t *testing.T) {
	testDB := Load("test.db", false)

	// testDB.Set("zp", "zp_value")

	value, ok := testDB.Get("zp")

	assert.Equal(t, ok, true)

	assert.Equal(t, value, "zp_value")

	fakeValue, ok := testDB.Get("fakeZp")

	assert.Equal(t, ok, true)

	fmt.Println(fakeValue)
}

// go test -v -run TestGetAll
func TestGetAll(t *testing.T) {
	testDB := Load("test.db", false)

	testDB.Set("zp", "zp_value")

	testDB.Set("zz", "zz_value")

	keys := testDB.GetAll()

	fmt.Println(keys)
}

//

// go test -v -run TestListValue
func TestListValue(t *testing.T) {
	testDB := Load("testList.db", false)

	// test step 1
	testDB.ListCreate("zpList")

	l1, ok := testDB.Get("zpList")

	assert.Equal(t, ok, true)

	fmt.Println(l1)

	// test step 2
	testDB.ListAdd("zpList", "zp")
	testDB.ListAdd("zpList", "zp1")
	testDB.ListAdd("zpList", "zp2")

	l2, ok := testDB.Get("zpList")

	assert.Equal(t, ok, true)

	fmt.Println(l2)

	fmt.Println(testDB)

	// test step 3

	arrs := []string{"zl1", "zl2", "zl3"}

	testDB.ListExtend("zpList", arrs)

	l3, ok := testDB.Get("zpList")

	assert.Equal(t, ok, true)

	fmt.Println(l3)

	fmt.Println(testDB)

	testDB.Dump()
}

// go test -v -run TestListLoad
func TestListLoad(t *testing.T) {
	testDB := Load("testList.db", false)

	l2, ok := testDB.Get("zpList")

	assert.Equal(t, ok, true)

	fmt.Println(l2)

	fmt.Println(testDB)
}

// go test -v -run TestDictValue
func TestDictValue(t *testing.T) {
	testDB := Load("testDict.db", false)

	testDB.DictCreate("zpDict")

	testDB.DictAdd("zpDict", "dict1", 1)

	arr1 := []int{1, 2, 3}
	testDB.DictAdd("zpDict", "dictarr1", arr1)

	map1 := map[string]string{
		"dict": "value",
	}

	testDB.DictAdd("zpDict", "dictmap1", map1)

	value, _ := testDB.Get("zpDict")

	testDB.Dump()

	fmt.Printf("%v", value)
}

// go test -v -run TestDictLoad
func TestDictLoad(t *testing.T) {
	testDB := Load("testDict.db", false)

	value, ok := testDB.Get("zpDict")
	assert.Equal(t, ok, true)

	fmt.Println(value)

	fmt.Println(testDB)
}
