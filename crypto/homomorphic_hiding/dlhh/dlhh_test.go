package dlhh_test

import (
	"crypto/rand"
	"encoding/binary"
	"testing"

	"github.com/titosilva/pdpr-go/crypto/homomorphic_hiding/dlhh"
	"github.com/titosilva/pdpr-go/math/dl"
	"github.com/titosilva/pdpr-go/math/nmod"
)

func Test__Hide__ShouldHideDataCorrectly__WhenDataIsPassed(t *testing.T) {
	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)
	data := []byte{1, 2, 3, 4}
	hidden := dlh.Hide(data)

	if hidden == nil {
		t.Errorf("Expected a non-nil value")
	}

	if !dlh.Verify(data, hidden) {
		t.Errorf("Expected the hidden data to be valid")
	}
}

func Test__Combine__ShouldCombineHiddenDataCorrectly__WhenTwoHiddenDataArePassed(t *testing.T) {
	// Arrange
	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)
	n1 := randUint()
	data1 := nmod.NewFromUint(n1, dlg.Mod).Bytes()
	n2 := randUint()
	data2 := nmod.NewFromUint(n2, dlg.Mod).Bytes()

	// Act
	hidden1 := dlh.Hide(data1)
	hidden2 := dlh.Hide(data2)
	combined := dlh.CombineHidden(hidden1, hidden2)

	// Assert
	if combined == nil {
		t.Errorf("Expected a non-nil value")
	}

	sum := n1 + n2
	if !dlh.Verify(nmod.NewFromUint(sum, dlg.Mod).Bytes(), combined) {
		t.Errorf("Expected the combined data to be valid")
	}
}

func randUint() uint64 {
	buf := make([]byte, 8)
	rand.Read(buf)
	return binary.BigEndian.Uint64(buf)
}
