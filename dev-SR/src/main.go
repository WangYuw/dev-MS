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
	go reg.GetConnection(srv, "localhost", config.DefaultPort)

	/*//test reg as client
	cln := clientimpl.NewClient()
	cln.SendQRequest("Auth", "v1.0.0", config.DefaultIP, config.DefaultPort)
	reg.SendQRequests(cln, config.DefaultPort)*/
}
