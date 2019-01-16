package main

import (
	"clientimpl"
	"config"
	"registry"
)

//TestClient tests unit
func main() {
	reg := registry.NewRegistry()
	cln := clientimpl.NewClient()
	cln.SendQRequest("Auth", "v1.0.0", "auth1", config.DefaultPort)
	reg.SendQRequests(cln, config.DefaultPort)
}
