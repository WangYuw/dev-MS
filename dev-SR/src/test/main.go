package main

import (
	"clientimpl"
	"config"
	"registry"
)

//test reg as client
func main() {
	reg := registry.NewRegistry()
	cln := clientimpl.NewClient()

	reg.SendQRequests(cln, config.DefaultPort)

	//cln.SendQRequest("Auth", "v1.0.0", "auth1", config.DefaultPort)
}
