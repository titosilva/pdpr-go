package nmod_test

import (
	"testing"

	"github.com/titosilva/pdpr-go/math/nmod"
)

func Test__Add__ShouldAddNumbersCorrectly__WhenNumbersHaveSameModulus(t *testing.T) {
	m1, _ := nmod.NewModulusFromInt(7)
	n1 := nmod.NewFromUint(4, m1)

	m2, _ := nmod.NewModulusFromInt(7)
	n2 := nmod.NewFromUint(3, m2)

	result := n1.Add(n2)
	zero := nmod.NewFromUint(0, m1)

	if !result.Equal(zero) {
		t.Errorf("Expected the result to be equals zero")
	}
}

func Test__Add__ShouldAddNumbersCorrectly__WhenNumbersHaveSameModulus__2(t *testing.T) {
	m1, _ := nmod.NewModulusFromInt(7)
	n1 := nmod.NewFromUint(4, m1)

	m2, _ := nmod.NewModulusFromInt(7)
	n2 := nmod.NewFromUint(5, m2)

	result := n1.Add(n2)
	two := nmod.NewFromUint(2, m1)

	if !result.Equal(two) {
		t.Errorf("Expected the result to be equals two")
	}
}

func Test__Add__ShouldReturnError__WhenNumbersHaveDifferentModulus(t *testing.T) {
	m1, _ := nmod.NewModulusFromInt(7)
	n1 := nmod.NewFromUint(4, m1)

	m2, _ := nmod.NewModulusFromInt(5)
	n2 := nmod.NewFromUint(3, m2)

	result := n1.Add(n2)

	if result != nil {
		t.Errorf("Expected nil to be returned")
	}
}

func Test__Sub__ShouldSubtractNumbersCorrectly__WhenNumbersHaveSameModulus(t *testing.T) {
	m1, _ := nmod.NewModulusFromInt(7)
	n1 := nmod.NewFromUint(4, m1)

	m2, _ := nmod.NewModulusFromInt(7)
	n2 := nmod.NewFromUint(3, m2)

	result := n1.Sub(n2)
	one := nmod.NewFromUint(1, m1)

	if !result.Equal(one) {
		t.Errorf("Expected the result to be equals one")
	}
}

func Test__Sub__ShouldReturnError__WhenNumbersHaveDifferentModulus(t *testing.T) {
	m1, _ := nmod.NewModulusFromInt(7)
	n1 := nmod.NewFromUint(4, m1)

	m2, _ := nmod.NewModulusFromInt(5)
	n2 := nmod.NewFromUint(3, m2)

	result := n1.Sub(n2)

	if result != nil {
		t.Errorf("Expected nil to be returned")
	}
}

func Test__Exp_ShouldCorrectlyComputeExponentiation__WhenNumbersHaveCorrectModulus(t *testing.T) {
	m, err := nmod.NewModulusFromInt(15)
	if err != nil {
		t.Fatal(err)
	}

	g := nmod.NewFromUint(2, m)
	e := g.Exp(nmod.NewFromUint(4, m))

	if !e.Equal(nmod.NewFromUint(1, m)) {
		t.Errorf("Expected the result to be equals one")
	}

	if !g.Equal(nmod.NewFromUint(2, m)) {
		t.Errorf("Expected the generator to remain the same")
	}
}
