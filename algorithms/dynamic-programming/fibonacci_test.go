package main

import "testing"

//BenchmarkFibonacci-8   	  405768	      3171 ns/op	       0 B/op	       0 allocs/op

func BenchmarkFibonacci(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_ = Fib(15)
	}
}

//BenchmarkFibonacci2-8   	671626482	         1.615 ns/op	       0 B/op	       0 allocs/op

func BenchmarkFibonacci2(b *testing.B) {
	n := 15
	memo := make([]int, n)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Fib2(n, memo)
	}
}

//BenchmarkFibonacci3-8   	122153936	         9.493 ns/op	       0 B/op	       0 allocs/op

func BenchmarkFibonacci3(b *testing.B) {
	n := 15
	k := make([]int, n+1)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Fib3(n, k)
	}
}
