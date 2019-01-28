package main

import (
	"clientimpl"
	"config"
	"registry"
	"rentities"
)

//test reg as client
func main() {
	reg := registry.NewRegistry()
	cln := clientimpl.NewClient()
	sqr, _ := rentities.NewSQR("Auth", "v1.0.0")
	reg.SendQRequest(cln, sqr, "auth1", config.DefaultPort)

	//cln.SendQRequest("Auth", "v1.0.0", "auth1", config.DefaultPort)
}
