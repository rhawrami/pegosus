//go:build !arm64

package numop

// SqrtF64 takes the square root of elements in `src`, and places the result in `dst`.
func SqrtF64(src, dst []float64) { sqrtF64Impl(src, dst) }

func sqrtF64Fallback(src, dst []float64) {}

// SqF64 takes the square of elements in `src`, and places the result in `dst`.
func SqF64(src, dst []float64) { sqF64Impl(src, dst) }

func sqF64Fallback(src, dst []float64) {}

// AbsF64 takes the absolute value of elements in `src`, and places the result in `dst`.
func AbsF64(src, dst []float64) { absF64Impl(src, dst) }

func absF64Fallback(src, dst []float64) {}

// NegF64 negates elements in `src`, and places the result in `dst`.
func NegF64(src, dst []float64) { negF64Impl(src, dst) }

func negF64Fallback(src, dst []float64) {}

// RecipF64 takes the reciprocal elements in `src`, and places the result in `dst`.
func RecipF64(src, dst []float64) { recipF64Impl(src, dst) }

func recipF64Fallback(src, dst []float64) {}

// SqrtF32 takes the square root of elements in `src`, and places the result in `dst`.
func SqrtF32(src, dst []float32) { sqrtF32Impl(src, dst) }

func sqrtF32Fallback(src, dst []float32) {}

// SqF32 takes the square of elements in `src`, and places the result in `dst`.
func SqF32(src, dst []float32) { sqF32Impl(src, dst) }

func sqF32Fallback(src, dst []float32) {}

// AbsF32 takes the absolute value of elements in `src`, and places the result in `dst`.
func AbsF32(src, dst []float32) { absF32Impl(src, dst) }

func absF32Fallback(src, dst []float32) {}

// NegF32 negates elements in `src`, and places the result in `dst`.
func NegF32(src, dst []float32) { negF32Impl(src, dst) }

func negF32Fallback(src, dst []float32) {}

// RecipF32 takes the reciprocal elements in `src`, and places the result in `dst`.
func RecipF32(src, dst []float32) { recipF32Impl(src, dst) }

func recipF32Fallback(src, dst []float32) {}
