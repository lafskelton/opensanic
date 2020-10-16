package main

import (
	"github.com/lafskelton/sanicdb/pkg/server"
	"github.com/lafskelton/sanicdb/pkg/splash"
)

func main() {
	sanic := server.StartService()
	splash.Print()
	//via flags
	go sanic.Test()
	sanic.Network()
}
