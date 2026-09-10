//go:build !arm64

package numop

// MaxI64 finds the maximum signed element in `src`, placing it in `dst[0]`.
func MaxI64(src, dst []int64) { maxI64Impl(src, dst) }

func maxI64Fallback(src, dst []int64) {}

// MinI64 finds the minimum signed element in `src`, placing it in `dst[0]`.
func MinI64(src, dst []int64) { minI64Impl(src, dst) }

func minI64Fallback(src, dst []int64) {}

// MaxI32 finds the maximum signed element in `src`, placing it in `dst[0]`.
func MaxI32(src, dst []int32) { maxI32Impl(src, dst) }

func maxI32Fallback(src, dst []int32) {}

// MinI32 finds the minimum signed element in `src`, placing it in `dst[0]`.
func MinI32(src, dst []int32) { minI32Impl(src, dst) }

func minI32Fallback(src, dst []int32) {}

// MaxF64 finds the maximum  element in `src` (ignoring NaN elements), placing it in `dst[0]`.
func MaxF64(src, dst []float64) { maxF64Impl(src, dst) }

func maxF64Fallback(src, dst []float64) {}

// MinF64 finds the minimum  element in `src` (ignoring NaN elements), placing it in `dst[0]`.
func MinF64(src, dst []float64) { minF64Impl(src, dst) }

func minF64Fallback(src, dst []float64) {}

// MaxF32 finds the maximum  element in `src` (ignoring NaN elements), placing it in `dst[0]`.
func MaxF32(src, dst []float32) { maxF32Impl(src, dst) }

func maxF32Fallback(src, dst []float32) {}

// MinF32 finds the minimum  element in `src` (ignoring NaN elements), placing it in `dst[0]`.
func MinF32(src, dst []float32) { minF32Impl(src, dst) }

func minF32Fallback(src, dst []float32) {}

// MinMaxI64 finds the minimum and maximum signed elements in `src`, placing the "min" in `dst[0]`
// and the "max" in `dst[1]`.
func MinMaxI64(src, dst []int64) { minMaxI64Impl(src, dst) }

func minMaxI64Fallback(src, dst []int64) {}

// MinMaxI32 finds the minimum and maximum signed elements in `src`, placing the "min" in `dst[0]`
// and the "max" in `dst[1]`.
func MinMaxI32(src, dst []int32) { minMaxI32Impl(src, dst) }

func minMaxI32Fallback(src, dst []int32) {}

// MinMaxF64 finds the minimum and maximum elements in `src` (ignoring NaN elements), placing
// the "min" in `dst[0]` and the "max" in `dst[1]`.
func MinMaxF64(src, dst []float64) { minMaxF64Impl(src, dst) }

func minMaxF64Fallback(src, dst []float64) {}

// MinMaxF32 finds the minimum and maximum elements in `src` (ignoring NaN elements), placing
// the "min" in `dst[0]` and the "max" in `dst[1]`.
func MinMaxF32(src, dst []float32) { minMaxF32Impl(src, dst) }

func minMaxF32Fallback(src, dst []float32) {}

// MaxI64WithValidity finds the maximum signed element in `src`, placing it in `dst[0]`. `validity`
// represents a validity bitmap, where only elements corresponding to set bits will be included in the
// calculation.
func MaxI64WithValidity(src, dst []int64, validity []byte) {
	maxI64WithValidityImpl(src, dst, validity)
}

func maxI64WithValidityFallback(src, dst []int64, validity []byte) {}

// MinI64WithValidity finds the minimum signed element in `src`, placing it in `dst[0]`. `validity`
// represents a validity bitmap, where only elements corresponding to set bits will be included in the
// calculation.
func MinI64WithValidity(src, dst []int64, validity []byte) {
	minI64WithValidityImpl(src, dst, validity)
}

func minI64WithValidityFallback(src, dst []int64, validity []byte) {}

// MaxI32WithValidity finds the maximum signed element in `src`, placing it in `dst[0]`. `validity`
// represents a validity bitmap, where only elements corresponding to set bits will be included in the
// calculation.
func MaxI32WithValidity(src, dst []int32, validity []byte) {
	maxI32WithValidityImpl(src, dst, validity)
}

func maxI32WithValidityFallback(src, dst []int32, validity []byte) {}

// MinI32WithValidity finds the minimum signed element in `src`, placing it in `dst[0]`. `validity`
// represents a validity bitmap, where only elements corresponding to set bits will be included in the
// calculation.
func MinI32WithValidity(src, dst []int32, validity []byte) {
	minI32WithValidityImpl(src, dst, validity)
}

func minI32WithValidityFallback(src, dst []int32, validity []byte) {}

// MaxF64WithValidity finds the maximum  element in `src` (ignoring NaN elements), placing it in `dst[0]`. `validity`
// represents a validity bitmap, where only elements corresponding to set bits will be included in the
// calculation.
func MaxF64WithValidity(src, dst []float64, validity []byte) {
	maxF64WithValidityImpl(src, dst, validity)
}

func maxF64WithValidityFallback(src, dst []float64, validity []byte) {}

// MinF64WithValidity finds the minimum  element in `src` (ignoring NaN elements), placing it in `dst[0]`. `validity`
// represents a validity bitmap, where only elements corresponding to set bits will be included in the
// calculation.
func MinF64WithValidity(src, dst []float64, validity []byte) {
	minF64WithValidityImpl(src, dst, validity)
}

func minF64WithValidityFallback(src, dst []float64, validity []byte) {}

// MaxF32WithValidity finds the maximum  element in `src` (ignoring NaN elements), placing it in `dst[0]`. `validity`
// represents a validity bitmap, where only elements corresponding to set bits will be included in the
// calculation.
func MaxF32WithValidity(src, dst []float32, validity []byte) {
	maxF32WithValidityImpl(src, dst, validity)
}

func maxF32WithValidityFallback(src, dst []float32, validity []byte) {}

// MinF32WithValidity finds the minimum  element in `src` (ignoring NaN elements), placing it in `dst[0]`. `validity`
// represents a validity bitmap, where only elements corresponding to set bits will be included in the
// calculation.
func MinF32WithValidity(src, dst []float32, validity []byte) {
	minF32WithValidityImpl(src, dst, validity)
}

func minF32WithValidityFallback(src, dst []float32, validity []byte) {}

// MinMaxI64WithValidity finds the minimum and maximum signed elements in `src`, placing the "min" in `dst[0]`
// and the "max" in `dst[1]`. `validity` represents a validity bitmap, where only elements
// corresponding to set bits will be included in the calculation.
func MinMaxI64WithValidity(src, dst []int64, validity []byte) {
	minMaxI64WithValidityImpl(src, dst, validity)
}

func minMaxI64WithValidityFallback(src, dst []int64, validity []byte) {}

// MinMaxI32WithValidity finds the minimum and maximum signed elements in `src`, placing the "min" in `dst[0]`
// and the "max" in `dst[1]`. `validity` represents a validity bitmap, where only elements
// corresponding to set bits will be included in the calculation.
func MinMaxI32WithValidity(src, dst []int32, validity []byte) {
	minMaxI32WithValidityImpl(src, dst, validity)
}

func minMaxI32WithValidityFallback(src, dst []int32, validity []byte) {}

// MinMaxF64WithValidity finds the minimum and maximum elements in `src` (ignoring NaN elements), placing
// the "min" in `dst[0]` and the "max" in `dst[1]`. `validity` represents a validity bitmap,
// where only elements corresponding to set bits will be included in the calculation.
func MinMaxF64WithValidity(src, dst []float64, validity []byte) {
	minMaxF64WithValidityImpl(src, dst, validity)
}

func minMaxF64WithValidityFallback(src, dst []float64, validity []byte) {}

// MinMaxF32WithValidity finds the minimum and maximum elements in `src` (ignoring NaN elements), placing
// the "min" in `dst[0]` and the "max" in `dst[1]`. `validity` represents a validity bitmap,
// where only elements corresponding to set bits will be included in the calculation.
func MinMaxF32WithValidity(src, dst []float32, validity []byte) {
	minMaxF32WithValidityImpl(src, dst, validity)
}

func minMaxF32WithValidityFallback(src, dst []float32, validity []byte) {}
