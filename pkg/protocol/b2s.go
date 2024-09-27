//go:build go1.20

package protocol

import "unsafe"

func b2s(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
