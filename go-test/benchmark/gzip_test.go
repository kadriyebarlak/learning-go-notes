package gzip

import "testing"

func BenchmarkGzipWithoutPoll(b *testing.B) {
	gzip := NewGzipWithoutPool()

	b.ResetTimer()             //bi onceki testten kalanlari temizle
	b.ReportAllocs()           // report allocations
	for n := 0; n < b.N; n++ { //go birim zamanda calistirilabilecek sayi.
		err := gzip.Zip("Go Turkiye Egitim Kampi - 203 Test & Benchmarks")

		if err != nil {
			b.Fatal(err)
		}

	}
}

func BenchmarkGzipWithPoll(b *testing.B) {
	gzip := NewGzipWithPool()

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		err := gzip.Zip("Go Turkiye Egitim Kampi - 203 Test & Benchmarks")

		if err != nil {
			b.Fatal(err)
		}

	}
}
