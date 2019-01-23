package main

import (
	"config"
	"registry"
	"serveimpl"
)

func main() {
	reg := registry.NewRegistry()

	//test reg as server
	srv := serveimpl.NewServer(config.DefaultProtocol)
	reg.GetConnection(srv, "reg1", config.DefaultPort)
}
