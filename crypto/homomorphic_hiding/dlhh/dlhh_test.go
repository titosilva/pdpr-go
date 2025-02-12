package dlhh_test

import (
	"testing"

	"github.com/titosilva/pdpr-go/crypto/homomorphic_hiding/dlhh"
	"github.com/titosilva/pdpr-go/math/dl"
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
	data1 := random(32)
	data2 := random(32)

	// Act
	hidden1 := dlh.Hide(data1)
	hidden2 := dlh.Hide(data2)
	combined := dlh.CombineHidden(hidden1, hidden2)

	// Assert
	if combined == nil {
		t.Errorf("Expected a non-nil value")
	}

	sum := dlh.CombinePlain(data1, data2)
	if !dlh.Verify(sum, combined) {
		t.Errorf("Expected the combined data to be valid")
	}
}
