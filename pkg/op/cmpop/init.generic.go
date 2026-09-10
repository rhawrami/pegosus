//go:build !amd64 && !arm64

package cmpop

var (
	gtI32Impl   = gtI32Fallback
	ltI32Impl   = ltI32Fallback
	geI32Impl   = geI32Fallback
	leI32Impl   = leI32Fallback
	eqI32Impl   = eqI32Fallback
	gtF32Impl   = gtF32Fallback
	ltF32Impl   = ltF32Fallback
	geF32Impl   = geF32Fallback
	leF32Impl   = leF32Fallback
	eqF32Impl   = eqF32Fallback
	neqI32Impl  = neqI32Fallback
	neqF32Impl  = neqF32Fallback
	betI32Impl  = betI32Fallback
	nBetI32Impl = nBetI32Fallback
	betF32Impl  = betF32Fallback
	nBetF32Impl = nBetF32Fallback
	gtI64Impl   = gtI64Fallback
	ltI64Impl   = ltI64Fallback
	geI64Impl   = geI64Fallback
	leI64Impl   = leI64Fallback
	eqI64Impl   = eqI64Fallback
	gtF64Impl   = gtF64Fallback
	ltF64Impl   = ltF64Fallback
	geF64Impl   = geF64Fallback
	leF64Impl   = leF64Fallback
	eqF64Impl   = eqF64Fallback
	neqI64Impl  = neqI64Fallback
	neqF64Impl  = neqF64Fallback
	betI64Impl  = betI64Fallback
	nBetI64Impl = nBetI64Fallback
	betF64Impl  = betF64Fallback
	nBetF64Impl = nBetF64Fallback
)
