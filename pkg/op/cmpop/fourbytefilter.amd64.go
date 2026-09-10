//go:build amd64

package cmpop

//go:noescape
func gtI32(src []int32, dst []byte, lit int32)

//go:noescape
func ltI32(src []int32, dst []byte, lit int32)

//go:noescape
func geI32(src []int32, dst []byte, lit int32)

//go:noescape
func leI32(src []int32, dst []byte, lit int32)

//go:noescape
func eqI32(src []int32, dst []byte, lit int32)

//go:noescape
func gtF32(src []float32, dst []byte, lit float32)

//go:noescape
func ltF32(src []float32, dst []byte, lit float32)

//go:noescape
func geF32(src []float32, dst []byte, lit float32)

//go:noescape
func leF32(src []float32, dst []byte, lit float32)

//go:noescape
func eqF32(src []float32, dst []byte, lit float32)

//go:noescape
func neqI32(src []int32, dst []byte, lit int32)

//go:noescape
func neqF32(src []float32, dst []byte, lit float32)

//go:noescape
func betI32(src []int32, dst []byte, min int32, max int32)

//go:noescape
func nBetI32(src []int32, dst []byte, min int32, max int32)

//go:noescape
func betF32(src []float32, dst []byte, min float32, max float32)

//go:noescape
func nBetF32(src []float32, dst []byte, min float32, max float32)
