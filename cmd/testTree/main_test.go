package main

import (
	"runtime"
	"testing"

	"github.com/lafskelton/sanicdb/pkg/store"
)

//this benchmarks the internal engine directly.
func BenchmarkSet(b *testing.B) {
	runtime.GOMAXPROCS(8)
	var s store.SanicBST
	s.Init()
	for n := 1; n < b.N; n++ {
		go s.TestTree(n)
	}
}
