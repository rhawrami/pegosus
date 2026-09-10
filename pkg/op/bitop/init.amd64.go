//go:build amd64

package bitop

import "golang.org/x/sys/cpu"

var (
	bitWiseAndWithPopCountImpl  = bitWiseAndWithPopCount
	bitWiseOrWithPopCountImpl   = bitWiseOrWithPopCount
	bitWiseAndNWithPopCountImpl = bitWiseAndNWithPopCount
	bitWiseXorWithPopCountImpl  = bitWiseXorWithPopCount
	popCountImpl                = popCount
	broadcastU8Impl             = broadcastU8
	broadcastI64Impl            = broadcastI64
	broadcastI32Impl            = broadcastI32
	broadcastF64Impl            = broadcastF64
	broadcastF32Impl            = broadcastF32
)

func init() {
	if cpu.X86.HasAVX2 {
		return
	}

	bitWiseAndWithPopCountImpl = bitWiseAndWithPopCountFallback
	bitWiseOrWithPopCountImpl = bitWiseOrWithPopCountFallback
	bitWiseAndNWithPopCountImpl = bitWiseAndNWithPopCountFallback
	bitWiseXorWithPopCountImpl = bitWiseXorWithPopCountFallback
	popCountImpl = popCountFallback
	broadcastU8Impl = broadcastU8Fallback
	broadcastI64Impl = broadcastI64Fallback
	broadcastI32Impl = broadcastI32Fallback
	broadcastF64Impl = broadcastF64Fallback
	broadcastF32Impl = broadcastF32Fallback
}
