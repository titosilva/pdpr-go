package nmod

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"math/big"

	"filippo.io/bigmod"
)

type NatMod struct {
	value   *bigmod.Nat
	modulus *Mod
}

type Mod struct {
	value   *bigmod.Modulus
	hash    string
	byteLen int
}

var ErrDifferentModulus = errors.New("different modulus")

func ensureSameModulus(i, j *NatMod) error {
	return i.modulus.Equal(j.modulus)
}

func new(value *bigmod.Nat, modulus *Mod) *NatMod {
	return &NatMod{value, modulus}
}

func newNatFromUint(value uint64, modulus *Mod) *bigmod.Nat {
	r := bigmod.NewNat()
	valueBuf := make([]byte, 8)
	binary.BigEndian.PutUint64(valueBuf, value)

	r.SetBytes(valueBuf, modulus.value)
	return r
}

func NewFromUint(value uint64, modulus *Mod) *NatMod {
	r := newNatFromUint(value, modulus)
	return new(r, modulus)
}

func NewFromBigEndianBytes(value []byte, modulus *Mod) *NatMod {
	r := bigmod.NewNat()
	r.SetBytes(value, modulus.value)
	return new(r, modulus)
}

func NewModulusFromInt(value uint64) (*Mod, error) {
	bi := big.NewInt(int64(value))
	mod, err := bigmod.NewModulusFromBig(bi)
	if err != nil {
		return nil, err
	}

	biBytes := bi.Bytes()
	sha := sha256.New()
	sha.Write(biBytes)
	hash := string(sha.Sum(nil))

	return &Mod{mod, hash, len(biBytes)}, nil
}

func NewModulusFromBigEndianBytes(value []byte) (*Mod, error) {
	bi := big.NewInt(0)
	bi.SetBytes(value)
	mod, err := bigmod.NewModulusFromBig(bi)
	if err != nil {
		return nil, err
	}

	sha := sha256.New()
	sha.Write(bi.Bytes())
	hash := string(sha.Sum(nil))

	return &Mod{mod, hash, len(value)}, nil
}

func (i *NatMod) Add(j *NatMod) *NatMod {
	if err := ensureSameModulus(i, j); err != nil {
		return nil
	}

	result := NewFromUint(0, i.modulus)
	result.value.Add(i.value, result.modulus.value)
	result.value.Add(j.value, result.modulus.value)

	return result
}

func (i *NatMod) Exp(exp *NatMod) *NatMod {
	return i.ExpBytes(exp.value.Bytes(i.modulus.value))
}

func (i *NatMod) ExpBytes(exp []byte) *NatMod {
	r := bigmod.NewNat()
	result := r.Exp(i.value, exp, i.modulus.value)
	return new(result, i.modulus)
}

func (i *NatMod) Mul(j *NatMod) *NatMod {
	if err := ensureSameModulus(i, j); err != nil {
		return nil
	}

	result := NewFromUint(0, i.modulus)
	result.value.Add(i.value, result.modulus.value)
	result.value.Mul(j.value, result.modulus.value)

	return result
}

func (i *NatMod) Sub(j *NatMod) *NatMod {
	if err := ensureSameModulus(i, j); err != nil {
		return nil
	}

	result := NewFromUint(0, i.modulus)
	result.value.Add(i.value, result.modulus.value)
	result.value.Sub(j.value, result.modulus.value)

	return result
}

func (i *NatMod) Equal(j *NatMod) bool {
	if err := ensureSameModulus(i, j); err != nil {
		return false
	}

	return i.value.Equal(j.value) == 1
}

func (i *NatMod) Bytes() []byte {
	valueBs := i.value.Bytes(i.modulus.value)
	return valueBs
}

func (i *NatMod) ModulusIsSame(j *NatMod) bool {
	return i.modulus.Equal(j.modulus) == nil
}

func (i *NatMod) ModulusIs(m *Mod) bool {
	return i.modulus.Equal(m) == nil
}

func (m1 *Mod) Dec() (*NatMod, error) {
	r := m1.value.Nat()
	r.Sub(newNatFromUint(1, m1), m1.value)
	bs := r.Bytes(m1.value)
	return NewFromBigEndianBytes(bs, m1), nil
}

func (m1 *Mod) Equal(m2 *Mod) error {
	if m1.byteLen != m2.byteLen {
		return ErrDifferentModulus
	}

	if m1.hash != m2.hash {
		return ErrDifferentModulus
	}

	return nil
}
