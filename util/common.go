package util

import "reflect"

func Copy(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
