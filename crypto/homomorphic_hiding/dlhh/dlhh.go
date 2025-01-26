package dlhh

import (
	"github.com/titosilva/pdpr-go/math/dl"
	"github.com/titosilva/pdpr-go/math/nmod"
)

type DLHider struct {
	dlg *dl.DiscreteLogGroup
}

func New(dlg *dl.DiscreteLogGroup) *DLHider {
	return &DLHider{dlg}
}

func (dlh *DLHider) Hide(data []byte) []byte {
	r := dlh.dlg.Gen.ExpBytes(data)
	return r.Bytes()
}

func (dlh *DLHider) Combine(hidden1 []byte, hidden2 []byte) []byte {
	h1Nat := nmod.NewFromBigEndianBytes(hidden1, dlh.dlg.Mod)
	h2Nat := nmod.NewFromBigEndianBytes(hidden2, dlh.dlg.Mod)
	r := h1Nat.Mul(h2Nat)

	return r.Bytes()
}

func (dlh *DLHider) Verify(data []byte, hidden []byte) bool {
	hNat := nmod.NewFromBigEndianBytes(hidden, dlh.dlg.Mod)
	dNat := nmod.NewFromBigEndianBytes(data, dlh.dlg.Mod)
	r := dlh.dlg.Gen.Exp(dNat)

	return r.Equal(hNat)
}
