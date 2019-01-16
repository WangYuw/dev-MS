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
	reg.GetConnection(srv, "172.25.0.3", config.DefaultPort)

	/*//test reg as client
	cln := clientimpl.NewClient()
	cln.SendQRequest("Auth", "v1.0.0", config.DefaultIP, config.DefaultPort)
	reg.SendQRequests(cln, config.DefaultPort)*/
}
