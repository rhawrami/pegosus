//go:build amd64

package cmpop

//go:noescape
func gtI64(src []int64, dst []byte, lit int64)

//go:noescape
func ltI64(src []int64, dst []byte, lit int64)

//go:noescape
func geI64(src []int64, dst []byte, lit int64)

//go:noescape
func leI64(src []int64, dst []byte, lit int64)

//go:noescape
func eqI64(src []int64, dst []byte, lit int64)

//go:noescape
func gtF64(src []float64, dst []byte, lit float64)

//go:noescape
func ltF64(src []float64, dst []byte, lit float64)

//go:noescape
func geF64(src []float64, dst []byte, lit float64)

//go:noescape
func leF64(src []float64, dst []byte, lit float64)

//go:noescape
func eqF64(src []float64, dst []byte, lit float64)

//go:noescape
func neqI64(src []int64, dst []byte, lit int64)

//go:noescape
func neqF64(src []float64, dst []byte, lit float64)

//go:noescape
func betI64(src []int64, dst []byte, min int64, max int64)

//go:noescape
func nBetI64(src []int64, dst []byte, min int64, max int64)

//go:noescape
func betF64(src []float64, dst []byte, min float64, max float64)

//go:noescape
func nBetF64(src []float64, dst []byte, min float64, max float64)
