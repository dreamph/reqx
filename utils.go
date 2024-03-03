package reqx

import "unsafe"

// toBytes returns a byte pointer without allocation.
func toBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// toString returns a string pointer without allocation
func toString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
