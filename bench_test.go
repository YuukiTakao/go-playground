package main

import (
	"math"
	"testing"
)

func BenchmarkAppend_AllocateEveryTime(b *testing.B) {
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		a := 1 << 62
		_ = a
	}
}

func BenchmarkAppend_AllocateOnce(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inf := int(math.Pow10(18))
		_ = inf
	}
}
