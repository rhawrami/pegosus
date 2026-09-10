//go:build !amd64 && !arm64

package strop

var (
	toUpperASCIIImpl = toUpperASCIIFallback
	toLowerASCIIImpl = toLowerASCIIFallback
)
