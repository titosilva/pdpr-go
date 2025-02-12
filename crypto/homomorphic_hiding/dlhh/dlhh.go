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

func (dlh *DLHider) CombineHidden(hidden1 []byte, hidden2 []byte) []byte {
	h1Nat := nmod.NewFromBigEndianBytes(hidden1, dlh.dlg.Mod)
	h2Nat := nmod.NewFromBigEndianBytes(hidden2, dlh.dlg.Mod)
	r := h1Nat.Mul(h2Nat)

	return r.Bytes()
}

func (dlh *DLHider) CombinePlain(data1 []byte, data2 []byte) []byte {
	d1Nat := nmod.NewFromBigEndianBytes(data1, dlh.dlg.MulMod)
	d2Nat := nmod.NewFromBigEndianBytes(data2, dlh.dlg.MulMod)
	r := d1Nat.Add(d2Nat)

	return r.Bytes()
}

func (dlh *DLHider) SubtractPlain(data1 []byte, data2 []byte) []byte {
	d1Nat := nmod.NewFromBigEndianBytes(data1, dlh.dlg.MulMod)
	d2Nat := nmod.NewFromBigEndianBytes(data2, dlh.dlg.MulMod)
	r := d1Nat.Sub(d2Nat)

	return r.Bytes()
}

func (dlh *DLHider) Verify(data []byte, hidden []byte) bool {
	hNat := nmod.NewFromBigEndianBytes(hidden, dlh.dlg.Mod)
	dNat := nmod.NewFromBigEndianBytes(data, dlh.dlg.Mod)
	r := dlh.dlg.Gen.Exp(dNat)

	return r.Equal(hNat)
}

func (dlh *DLHider) VerifyHidden(hidden1 []byte, hidden2 []byte) bool {
	h1Nat := nmod.NewFromBigEndianBytes(hidden1, dlh.dlg.Mod)
	h2Nat := nmod.NewFromBigEndianBytes(hidden2, dlh.dlg.Mod)

	return h1Nat.Equal(h2Nat)
}
