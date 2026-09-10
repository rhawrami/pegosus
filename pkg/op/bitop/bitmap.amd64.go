//go:build amd64

package bitop

//go:noescape
func bitWiseAndWithPopCount(src1, src2, dst []byte) uint64

//go:noescape
func bitWiseOrWithPopCount(src1, src2, dst []byte) uint64

//go:noescape
func bitWiseXorWithPopCount(src1, src2, dst []byte) uint64

//go:noescape
func bitWiseAndNWithPopCount(src1, src2, dst []byte) uint64

//go:noescape
func popCount(src []byte) uint64
