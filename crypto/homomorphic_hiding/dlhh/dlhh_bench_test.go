package dlhh_test

import (
	"crypto/rand"
	"testing"

	"github.com/titosilva/pdpr-go/crypto/homomorphic_hiding/dlhh"
	"github.com/titosilva/pdpr-go/math/dl"
)

func runPdpr() {
	m := random(32)
	k := random(32)

	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)

	mH := dlh.Hide(m)

	enc := dlh.CombinePlain(m, k)

	kH := dlh.Hide(k)

	encH2 := dlh.CombineHidden(mH, kH)

	if !dlh.Verify(enc, encH2) {
		panic("Failed verification")
	}
}

func Benchmark__PDPr__HomomorphicHidings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runPdpr()
	}
}

func random(size int) []byte {
	r := make([]byte, size)

	if _, err := rand.Read(r); err != nil {
		panic(err)
	}

	return r
}
