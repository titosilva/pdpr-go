package dlhh_test

import (
	"crypto/rand"
	"testing"

	"github.com/titosilva/pdpr-go/crypto/homomorphic_hiding/dlhh"
	"github.com/titosilva/pdpr-go/math/dl"
)

func runPdpr(size int) {
	m := random(size)
	k := random(size)

	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)

	enc := encrypt(dlh, m, k)
	proof := generateProof(dlh, enc)
	mH := dlh.Hide(m)
	verified := verifyProof(dlh, mH, proof, k)
	if !verified {
		panic("proof not verified")
	}

	dec := decrypt(dlh, enc, k)
	if string(dec[len(dec)-len(m):]) != string(m) {
		panic("decryption failed")
	}
}

func Benchmark__PDPr__HomomorphicHidings__256b(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runPdpr(32)
	}

	b.ReportMetric(float64(b.Elapsed().Milliseconds())/float64(b.N), "ms/pdpr")
}

func Benchmark__PDPr__HomomorphicHidings__128b(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runPdpr(16)
	}

	b.ReportMetric(float64(b.Elapsed().Milliseconds())/float64(b.N), "ms/pdpr")
}

func encrypt(dlh *dlhh.DLHider, m, k []byte) []byte {
	return dlh.CombinePlain(m, k)
}

func decrypt(dlh *dlhh.DLHider, enc, k []byte) []byte {
	return dlh.SubtractPlain(enc, k)
}

func generateProof(dlh *dlhh.DLHider, c []byte) []byte {
	return dlh.Hide(c)
}

func verifyProof(dlh *dlhh.DLHider, mH, proof, k []byte) bool {
	kH := dlh.Hide(k)
	return dlh.VerifyHidden(proof, dlh.CombineHidden(mH, kH))
}

func Benchmark__Encrypt__256b(b *testing.B) {
	m := random(32)
	k := random(32)
	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		encrypt(dlh, m, k)
	}

	b.ReportMetric(float64(b.Elapsed().Milliseconds())/float64(b.N), "ms")
}

func Benchmark__Decrypt__256b(b *testing.B) {
	m := random(32)
	k := random(32)
	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)
	enc := encrypt(dlh, m, k)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		decrypt(dlh, enc, k)
	}

	b.ReportMetric(float64(b.Elapsed().Milliseconds())/float64(b.N), "ms")
}

func Benchmark__GenerateProof__256b(b *testing.B) {
	m := random(32)
	k := random(32)
	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)
	enc := encrypt(dlh, m, k)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		generateProof(dlh, enc)
	}

	b.ReportMetric(float64(b.Elapsed().Milliseconds())/float64(b.N), "ms")
}

func Benchmark__VerifyProof__256b(b *testing.B) {
	m := random(32)
	k := random(32)
	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)
	enc := encrypt(dlh, m, k)
	proof := generateProof(dlh, enc)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		verifyProof(dlh, m, proof, k)
	}

	b.ReportMetric(float64(b.Elapsed().Milliseconds())/float64(b.N), "ms")
}

func Benchmark__Encrypt__128b(b *testing.B) {
	m := random(16)
	k := random(16)
	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		encrypt(dlh, m, k)
	}

	b.ReportMetric(float64(b.Elapsed().Milliseconds())/float64(b.N), "ms")
}

func Benchmark__Decrypt__128b(b *testing.B) {
	m := random(16)
	k := random(16)
	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)
	enc := encrypt(dlh, m, k)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		decrypt(dlh, enc, k)
	}

	b.ReportMetric(float64(b.Elapsed().Milliseconds())/float64(b.N), "ms")
}

func Benchmark__GenerateProof__128b(b *testing.B) {
	m := random(16)
	k := random(16)
	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)
	enc := encrypt(dlh, m, k)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		generateProof(dlh, enc)
	}

	b.ReportMetric(float64(b.Elapsed().Milliseconds())/float64(b.N), "ms")
}

func Benchmark__VerifyProof__128b(b *testing.B) {
	m := random(16)
	k := random(16)
	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)
	enc := encrypt(dlh, m, k)
	proof := generateProof(dlh, enc)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		verifyProof(dlh, m, proof, k)
	}

	b.ReportMetric(float64(b.Elapsed().Milliseconds())/float64(b.N), "ms")
}

func Benchmark__Expo__128b(b *testing.B) {
	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)
	m := random(16)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dlh.Hide(m)
	}

	b.ReportMetric(float64(b.Elapsed().Milliseconds())/float64(b.N), "ms")
}

func Benchmark__Expo__256b(b *testing.B) {
	dlg := dl.NewOakley2Group()
	dlh := dlhh.New(dlg)
	m := random(16)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dlh.Hide(m)
	}

	b.ReportMetric(float64(b.Elapsed().Milliseconds())/float64(b.N), "ms")
}

func random(size int) []byte {
	r := make([]byte, size)

	if _, err := rand.Read(r); err != nil {
		panic(err)
	}

	return r
}
