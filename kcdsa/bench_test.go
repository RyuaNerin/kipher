package kcdsa

import (
	"bufio"
	"crypto/rand"
	"testing"
)

func benchmarkAllSizes(b *testing.B, f func(*testing.B, ParameterSizes)) {
	tests := []struct {
		name  string
		sizes ParameterSizes
	}{
		{"L2048 N224 SHA224", L2048N224SHA224},
		{"L2048 N224 SHA256", L2048N224SHA256},
		{"L2048 N256 SHA256", L2048N256SHA256},
		{"L3072 N256 SHA256", L3072N256SHA256},
	}
	for _, test := range tests {
		test := test
		b.Run(test.name, func(b *testing.B) {
			f(b, test.sizes)
		})
	}
}

var rnd = bufio.NewReaderSize(rand.Reader, 1<<15)

func Benchmark_GenerateParameters_GO(b *testing.B) {
	benchmarkAllSizes(b, func(b *testing.B, ps ParameterSizes) {

		var params Parameters
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := GenerateParameters(&params, rnd, ps); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_GenerateParameters_TTAK(b *testing.B) {
	benchmarkAllSizes(b, func(b *testing.B, ps ParameterSizes) {
		var params Parameters
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if _, _, err := GenerateParametersTTAK(&params, rnd, ps); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_GenerateKey_Go(b *testing.B) {
	benchmarkAllSizes(b, func(b *testing.B, ps ParameterSizes) {
		var priv PrivateKey
		if err := GenerateParameters(&priv.Parameters, rnd, ps); err != nil {
			b.Error(err)
		}

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := GenerateKey(&priv, rnd); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_GenerateKey_TTAK(b *testing.B) {
	benchmarkAllSizes(b, func(b *testing.B, ps ParameterSizes) {
		var priv PrivateKey
		if _, _, err := GenerateParametersTTAK(&priv.Parameters, rnd, ps); err != nil {
			b.Error(err)
		}

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := GenerateKeyTTAK(&priv, rnd, UserProvidedRandomInput); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_Sign(b *testing.B) {
	benchmarkAllSizes(b, func(b *testing.B, ps ParameterSizes) {
		data := []byte(`text`)

		var priv PrivateKey
		if err := GenerateParameters(&priv.Parameters, rnd, ps); err != nil {
			b.Error(err)
		}
		if err := GenerateKey(&priv, rnd); err != nil {
			b.Error(err)
		}

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r, _, err := Sign(rnd, &priv, data)
			if err != nil {
				b.Error(err)
			}
			data = r.Bytes()
		}
	})

}

func Benchmark_Verify(b *testing.B) {
	benchmarkAllSizes(b, func(b *testing.B, ps ParameterSizes) {
		data := []byte(`text`)

		var priv PrivateKey
		if err := GenerateParameters(&priv.Parameters, rnd, ps); err != nil {
			b.Error(err)
		}
		if err := GenerateKey(&priv, rnd); err != nil {
			b.Error(err)
		}

		r, s, err := Sign(rnd, &priv, data)
		if err != nil {
			b.Error(err)
		}

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ok := Verify(&priv.PublicKey, data, r, s)
			if !ok {
				b.Errorf("%d: Verify failed", i)
			}
		}
	})
}
