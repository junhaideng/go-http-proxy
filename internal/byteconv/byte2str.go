package byteconv

import "unsafe"

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
