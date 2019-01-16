package main

import (
	"config"
	"rentities"
	"services"
)

//TestServer tests unit
func main() {
	auth1, _ := services.NewAuth("1.0.0.0", 1, "v1.0.0")
	auth2, _ := services.NewAuth("1.0.0.1", 2, "v1.0.0")
	ri1, _ := rentities.NewRegisterInfo(auth1.Name, auth1.IID, auth1.IP, auth1.Version)
	ri2, _ := rentities.NewRegisterInfo(auth2.Name, auth2.IID, auth2.IP, auth2.Version)
	req, _ := rentities.NewServiceRequest("Auth", "v1.0.0")

	//client
	const uri string = "http://" + "reg1" + ":" + config.DefaultPort

	auth1.ListAllServices(uri)
	auth1.Register(ri1, uri)
	auth2.Register(ri2, uri)
	auth1.ListAllServices(uri)
	auth1.FindService(*req, uri)
}
