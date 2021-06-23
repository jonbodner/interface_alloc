package interface_alloc

import "testing"

var total int

func BenchmarkALoop(b *testing.B) {
	for i :=0 ;i < b.N;i++ {
		total += ALoop()
	}
}

func BenchmarkBLoop(b *testing.B) {
	for i :=0 ;i < b.N;i++ {
		total += BLoop()
	}
}
