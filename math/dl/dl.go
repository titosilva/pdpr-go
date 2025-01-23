package dl

import (
	"github.com/titosilva/pdpr-go/math/nmod"
)

type DiscreteLogGroup struct {
	Mod *nmod.Mod
	Gen *nmod.NatMod
}

func (dlg *DiscreteLogGroup) Index(n *nmod.NatMod) (*nmod.NatMod, error) {
	return dlg.Gen.Exp(n), nil
}

func NewOakley2Group() *DiscreteLogGroup {
	oakley2primeBytes := []byte{
		255, 255, 255, 255, 255, 255, 255, 255,
		201, 15, 218, 162, 33, 104, 194, 52,
		196, 198, 98, 139, 128, 220, 28, 209, 41,
		2, 78, 8, 138, 103, 204, 116, 2, 11, 190,
		166, 59, 19, 155, 34, 81, 74, 8, 121, 142,
		52, 4, 221, 239, 149, 25, 179, 205, 58, 67,
		27, 48, 43, 10, 109, 242, 95, 20, 55, 79,
		225, 53, 109, 109, 81, 194, 69, 228, 133,
		181, 118, 98, 94, 126, 198, 244, 76, 66, 233,
		166, 55, 237, 107, 11, 255, 92, 182, 244, 6,
		183, 237, 238, 56, 107, 251, 90, 137, 159, 165,
		174, 159, 36, 17, 124, 75, 31, 230, 73, 40, 102,
		81, 236, 230, 83, 129, 255, 255, 255, 255, 255,
		255, 255, 255,
	}

	m, _ := nmod.NewModulusFromBigEndianBytes(oakley2primeBytes)
	g := nmod.NewFromUint(2, m)

	return &DiscreteLogGroup{
		Mod: m,
		Gen: g,
	}
}
