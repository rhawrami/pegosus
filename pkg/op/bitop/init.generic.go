//go:build !amd64 && !arm64

package bitop

var (
	bitWiseAndWithPopCountImpl  = bitWiseAndWithPopCountFallback
	bitWiseOrWithPopCountImpl   = bitWiseOrWithPopCountFallback
	bitWiseAndNWithPopCountImpl = bitWiseAndNWithPopCountFallback
	bitWiseXorWithPopCountImpl  = bitWiseXorWithPopCountFallback
	popCountImpl                = popCountFallback
	broadcastU8Impl             = broadcastU8Fallback
	broadcastI64Impl            = broadcastI64Fallback
	broadcastI32Impl            = broadcastI32Fallback
	broadcastF64Impl            = broadcastF64Fallback
	broadcastF32Impl            = broadcastF32Fallback
)
