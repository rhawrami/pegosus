//go:build !arm64

package cmpop

// GtI32 checks if `src` elements are greater than `lit`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func GtI32(src []int32, dst []byte, lit int32) {
	gtI32Impl(src, dst, lit)
}

func gtI32Fallback(src []int32, dst []byte, lit int32) {}

// LtI32 checks if `src` elements are less than `lit`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func LtI32(src []int32, dst []byte, lit int32) {
	ltI32Impl(src, dst, lit)
}

func ltI32Fallback(src []int32, dst []byte, lit int32) {}

// GeI32 checks if `src` elements are greater than or equal to `lit`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func GeI32(src []int32, dst []byte, lit int32) {
	geI32Impl(src, dst, lit)
}

func geI32Fallback(src []int32, dst []byte, lit int32) {}

// LeI32 checks if `src` elements are less than or equal to `lit`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func LeI32(src []int32, dst []byte, lit int32) {
	leI32Impl(src, dst, lit)
}

func leI32Fallback(src []int32, dst []byte, lit int32) {}

// EqI32 checks if `src` elements are equal to `lit`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func EqI32(src []int32, dst []byte, lit int32) {
	eqI32Impl(src, dst, lit)
}

func eqI32Fallback(src []int32, dst []byte, lit int32) {}

// GtF32 checks if `src` elements are greater than `lit`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func GtF32(src []float32, dst []byte, lit float32) {
	gtF32Impl(src, dst, lit)
}

func gtF32Fallback(src []float32, dst []byte, lit float32) {}

// LtF32 checks if `src` elements are less than `lit`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func LtF32(src []float32, dst []byte, lit float32) {
	ltF32Impl(src, dst, lit)
}

func ltF32Fallback(src []float32, dst []byte, lit float32) {}

// GeF32 checks if `src` elements are greater than or equal to `lit`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func GeF32(src []float32, dst []byte, lit float32) {
	geF32Impl(src, dst, lit)
}

func geF32Fallback(src []float32, dst []byte, lit float32) {}

// LeF32 checks if `src` elements are less than or equal to `lit`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func LeF32(src []float32, dst []byte, lit float32) {
	leF32Impl(src, dst, lit)
}

func leF32Fallback(src []float32, dst []byte, lit float32) {}

// EqF32 checks if `src` elements are equal to `lit`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func EqF32(src []float32, dst []byte, lit float32) {
	eqF32Impl(src, dst, lit)
}

func eqF32Fallback(src []float32, dst []byte, lit float32) {}

// NeqI32 checks if `src` elements are not equal to `lit`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func NeqI32(src []int32, dst []byte, lit int32) {
	neqI32Impl(src, dst, lit)
}

func neqI32Fallback(src []int32, dst []byte, lit int32) {}

// NeqF32 checks if `src` elements are equal to `lit`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func NeqF32(src []float32, dst []byte, lit float32) {
	neqF32Impl(src, dst, lit)
}

func neqF32Fallback(src []float32, dst []byte, lit float32) {}

// BetI32 checks if `src` elements are between `min` and `max`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func BetI32(src []int32, dst []byte, min int32, max int32) {
	betI32Impl(src, dst, min, max)
}

func betI32Fallback(src []int32, dst []byte, min int32, max int32) {}

// NBetI32 checks if `src` elements are not between `min` and `max`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func NBetI32(src []int32, dst []byte, min int32, max int32) {
	nBetI32Impl(src, dst, min, max)
}

func nBetI32Fallback(src []int32, dst []byte, min int32, max int32) {}

// BetF32 checks if `src` elements are between `min` and `max`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func BetF32(src []float32, dst []byte, min float32, max float32) {
	betF32Impl(src, dst, min, max)
}

func betF32Fallback(src []float32, dst []byte, min float32, max float32) {}

// NBetF32 checks if `src` elements are not between `min` and `max`, and places
// the bitpacked results in `dst` (bit is set for elements passing condition).
func NBetF32(src []float32, dst []byte, min float32, max float32) {
	nBetF32Impl(src, dst, min, max)
}

func nBetF32Fallback(src []float32, dst []byte, min float32, max float32) {}
