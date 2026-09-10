//go:build amd64

package strop

import "golang.org/x/sys/cpu"

var (
	toUpperASCIIImpl = toUpperASCII
	toLowerASCIIImpl = toLowerASCII
)

func init() {
	if cpu.X86.HasAVX2 {
		return
	}

	toUpperASCIIImpl = toUpperASCIIFallback
	toLowerASCIIImpl = toLowerASCIIFallback
}
