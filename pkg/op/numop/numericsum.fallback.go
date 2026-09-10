//go:build !arm64

package numop

func SumI64(src, dst []int64) { sumI64Impl(src, dst) }

func sumI64Fallback(src, dst []int64) {}

func SumI32(src []int32, dst []int64) { sumI32Impl(src, dst) }

func sumI32Fallback(src []int32, dst []int64) {}

func SumF64(src, dst []float64) { sumF64Impl(src, dst) }

func sumF64Fallback(src, dst []float64) {}

func SumF32(src []float32, dst []float64) { sumF32Impl(src, dst) }

func sumF32Fallback(src []float32, dst []float64) {}

func SumI64WithValidity(src, dst []int64, validity []byte) {
	sumI64WithValidityImpl(src, dst, validity)
}

func sumI64WithValidityFallback(src, dst []int64, validity []byte) {}

func SumI32WithValidity(src []int32, dst []int64, validity []byte) {
	sumI32WithValidityImpl(src, dst, validity)
}

func sumI32WithValidityFallback(src []int32, dst []int64, validity []byte) {}

func SumF64WithValidity(src, dst []float64, validity []byte) {
	sumF64WithValidityImpl(src, dst, validity)
}

func sumF64WithValidityFallback(src, dst []float64, validity []byte) {}

func SumF32WithValidity(src []float32, dst []float64, validity []byte) {
	sumF32WithValidityImpl(src, dst, validity)
}

func sumF32WithValidityFallback(src []float32, dst []float64, validity []byte) {}
