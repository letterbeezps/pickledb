package util

import (
	"encoding/json"
	"reflect"
)

func Copy(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

func Typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func Convert(v interface{}) interface{} {
	byteV, _ := json.Marshal(v)

	var ret interface{}

	if err := json.Unmarshal(byteV, &ret); err != nil {
		panic(err)
	}

	return ret
}
