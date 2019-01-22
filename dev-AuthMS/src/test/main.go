package main

import (
	"config"
	"rentities"
	"services"
)

//test auth as client
func main() {
	auth1, _ := services.NewAuth("auth1", 1, "v1.0.0")
	auth2, _ := services.NewAuth("auth2", 2, "v1.0.0")
	ri1, _ := rentities.NewRegisterInfo(auth1.Name, auth1.IID, auth1.IP, auth1.Version, auth1.Quality)
	ri2, _ := rentities.NewRegisterInfo(auth2.Name, auth2.IID, auth2.IP, auth2.Version, auth2.Quality)
	req, _ := rentities.NewServiceRequest("Auth", "v1.0.0")

	//client
	const uri string = "http://" + "reg1" + ":" + config.DefaultPort

	auth1.ListAllServices(uri)
	auth1.Register(ri1, uri)
	auth2.Register(ri2, uri)
	auth1.ListAllServices(uri)
	auth1.FindService(*req, uri)
}
