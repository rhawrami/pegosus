//go:build !amd64 && !arm64

package numop

var (
	addF64LitImpl = addF64LitFallback
	subF64LitImpl = subF64LitFallback
	mulF64LitImpl = mulF64LitFallback
	divF64LitImpl = divF64LitFallback
	addF32LitImpl = addF32LitFallback
	subF32LitImpl = subF32LitFallback
	mulF32LitImpl = mulF32LitFallback
	divF32LitImpl = divF32LitFallback
	addF64VecImpl = addF64VecFallback
	subF64VecImpl = subF64VecFallback
	mulF64VecImpl = mulF64VecFallback
	divF64VecImpl = divF64VecFallback
	addF32VecImpl = addF32VecFallback
	subF32VecImpl = subF32VecFallback
	mulF32VecImpl = mulF32VecFallback
	divF32VecImpl = divF32VecFallback

	sqrtF64Impl  = sqrtF64Fallback
	sqF64Impl    = sqF64Fallback
	absF64Impl   = absF64Fallback
	negF64Impl   = negF64Fallback
	recipF64Impl = recipF64Fallback
	sqrtF32Impl  = sqrtF32Fallback
	sqF32Impl    = sqF32Fallback
	absF32Impl   = absF32Fallback
	negF32Impl   = negF32Fallback
	recipF32Impl = recipF32Fallback

	addI64LitImpl = addI64LitFallback
	subI64LitImpl = subI64LitFallback
	mulI64LitImpl = mulI64LitFallback
	divI64LitImpl = divI64LitFallback
	addI32LitImpl = addI32LitFallback
	subI32LitImpl = subI32LitFallback
	mulI32LitImpl = mulI32LitFallback
	divI32LitImpl = divI32LitFallback
	addI64VecImpl = addI64VecFallback
	subI64VecImpl = subI64VecFallback
	mulI64VecImpl = mulI64VecFallback
	divI64VecImpl = divI64VecFallback
	addI32VecImpl = addI32VecFallback
	subI32VecImpl = subI32VecFallback
	mulI32VecImpl = mulI32VecFallback
	divI32VecImpl = divI32VecFallback

	sqrtI64Impl  = sqrtI64Fallback
	sqI64Impl    = sqI64Fallback
	absI64Impl   = absI64Fallback
	negI64Impl   = negI64Fallback
	recipI64Impl = recipI64Fallback
	sqrtI32Impl  = sqrtI32Fallback
	sqI32Impl    = sqI32Fallback
	absI32Impl   = absI32Fallback
	negI32Impl   = negI32Fallback
	recipI32Impl = recipI32Fallback

	maxI64Impl    = maxI64Fallback
	minI64Impl    = minI64Fallback
	maxI32Impl    = maxI32Fallback
	minI32Impl    = minI32Fallback
	maxF64Impl    = maxF64Fallback
	minF64Impl    = minF64Fallback
	maxF32Impl    = maxF32Fallback
	minF32Impl    = minF32Fallback
	minMaxI64Impl = minMaxI64Fallback
	minMaxI32Impl = minMaxI32Fallback
	minMaxF64Impl = minMaxF64Fallback
	minMaxF32Impl = minMaxF32Fallback

	maxI64WithValidityImpl    = maxI64WithValidityFallback
	minI64WithValidityImpl    = minI64WithValidityFallback
	maxI32WithValidityImpl    = maxI32WithValidityFallback
	minI32WithValidityImpl    = minI32WithValidityFallback
	maxF64WithValidityImpl    = maxF64WithValidityFallback
	minF64WithValidityImpl    = minF64WithValidityFallback
	maxF32WithValidityImpl    = maxF32WithValidityFallback
	minF32WithValidityImpl    = minF32WithValidityFallback
	minMaxI64WithValidityImpl = minMaxI64WithValidityFallback
	minMaxI32WithValidityImpl = minMaxI32WithValidityFallback
	minMaxF64WithValidityImpl = minMaxF64WithValidityFallback
	minMaxF32WithValidityImpl = minMaxF32WithValidityFallback

	castI64ToF64Impl = castI64ToF64Fallback
	castI32ToF32Impl = castI32ToF32Fallback
	castF64ToI64Impl = castF64ToI64Fallback
	castF32ToI32Impl = castF32ToI32Fallback
	castF32ToF64Impl = castF32ToF64Fallback
	castI32ToI64Impl = castI32ToI64Fallback
	castI32ToF64Impl = castI32ToF64Fallback
	castF32ToI64Impl = castF32ToI64Fallback
	castI64ToF32Impl = castI64ToF32Fallback
	castF64ToI32Impl = castF64ToI32Fallback

	clipF64WithF64BoundsImpl = clipF64WithF64BoundsFallback
	clipF32WithF32BoundsImpl = clipF32WithF32BoundsFallback
	clipI32WithI32BoundsImpl = clipI32WithI32BoundsFallback
	clipI64WithI64BoundsImpl = clipI64WithI64BoundsFallback
	clipI64WithF64BoundsImpl = clipI64WithF64BoundsFallback
	clipI32WithF32BoundsImpl = clipI32WithF32BoundsFallback

	sumI64Impl             = sumI64Fallback
	sumI32Impl             = sumI32Fallback
	sumF64Impl             = sumF64Fallback
	sumF32Impl             = sumF32Fallback
	sumI64WithValidityImpl = sumI64WithValidityFallback
	sumI32WithValidityImpl = sumI32WithValidityFallback
	sumF64WithValidityImpl = sumF64WithValidityFallback
	sumF32WithValidityImpl = sumF32WithValidityFallback
)
