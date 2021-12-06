package byteconv

import "unsafe"

func Str2Byte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}