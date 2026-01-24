package main

import (
	"math"
	"testing"
)

// go test -bench=. conversion_test.go
// formula of shift n << m == 2^m
// 1 << 10 == 2^10

var result int

func BenchmarkConversion1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = 10 << 20
		_ = result
	}
}

func BenchmarkConversion2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = 10 * 1024 * 1024
		_ = result
	}
}

func BenchmarkConversion3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = 10 * (int)(math.Pow(2, 20))
		_ = result
	}
}
