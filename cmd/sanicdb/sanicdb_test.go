package main

import (
	"runtime"
	"testing"

	"github.com/lafskelton/sanicdb/pkg/server"
)

//
func BenchmarkSet(b *testing.B) {
	runtime.GOMAXPROCS(16)
	sanic := server.StartService()
	//Responce callback
	//Responce callback
	resp := func(data interface{}) {
		_ = data
		//This contains completion data about the op.
		//in open sanic, only numerical keys are enabled.
		//the returned data type will be store.COMPLETED
		//see /pkg/tasks.go
		return
	}
	for n := 1; n < b.N; n++ {
		sanic.Set("test", uint32(n), "hello", resp)
	}

}
