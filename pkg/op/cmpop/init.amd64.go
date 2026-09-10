//go:build amd64

package cmpop

import "golang.org/x/sys/cpu"

var (
	gtI32Impl   = gtI32
	ltI32Impl   = ltI32
	geI32Impl   = geI32
	leI32Impl   = leI32
	eqI32Impl   = eqI32
	gtF32Impl   = gtF32
	ltF32Impl   = ltF32
	geF32Impl   = geF32
	leF32Impl   = leF32
	eqF32Impl   = eqF32
	neqI32Impl  = neqI32
	neqF32Impl  = neqF32
	betI32Impl  = betI32
	nBetI32Impl = nBetI32
	betF32Impl  = betF32
	nBetF32Impl = nBetF32
	gtI64Impl   = gtI64
	ltI64Impl   = ltI64
	geI64Impl   = geI64
	leI64Impl   = leI64
	eqI64Impl   = eqI64
	gtF64Impl   = gtF64
	ltF64Impl   = ltF64
	geF64Impl   = geF64
	leF64Impl   = leF64
	eqF64Impl   = eqF64
	neqI64Impl  = neqI64
	neqF64Impl  = neqF64
	betI64Impl  = betI64
	nBetI64Impl = nBetI64
	betF64Impl  = betF64
	nBetF64Impl = nBetF64
)

func init() {
	if cpu.X86.HasAVX2 {
		return
	}

	gtI32Impl = gtI32Fallback
	ltI32Impl = ltI32Fallback
	geI32Impl = geI32Fallback
	leI32Impl = leI32Fallback
	eqI32Impl = eqI32Fallback
	gtF32Impl = gtF32Fallback
	ltF32Impl = ltF32Fallback
	geF32Impl = geF32Fallback
	leF32Impl = leF32Fallback
	eqF32Impl = eqF32Fallback
	neqI32Impl = neqI32Fallback
	neqF32Impl = neqF32Fallback
	betI32Impl = betI32Fallback
	nBetI32Impl = nBetI32Fallback
	betF32Impl = betF32Fallback
	nBetF32Impl = nBetF32Fallback
	gtI64Impl = gtI64Fallback
	ltI64Impl = ltI64Fallback
	geI64Impl = geI64Fallback
	leI64Impl = leI64Fallback
	eqI64Impl = eqI64Fallback
	gtF64Impl = gtF64Fallback
	ltF64Impl = ltF64Fallback
	geF64Impl = geF64Fallback
	leF64Impl = leF64Fallback
	eqF64Impl = eqF64Fallback
	neqI64Impl = neqI64Fallback
	neqF64Impl = neqF64Fallback
	betI64Impl = betI64Fallback
	nBetI64Impl = nBetI64Fallback
	betF64Impl = betF64Fallback
	nBetF64Impl = nBetF64Fallback
}
