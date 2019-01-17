package main

import (
	"config"
	"services"
)

func main() {
	auth1, _ := services.NewAuth("1.0.0.0", 1, "v1.0.0")
	/*auth2, _ := services.NewAuth("1.0.0.1", 2, "v1.0.0")
	ri1, _ := rentities.NewRegisterInfo(auth1.Name, auth1.IID, auth1.IP, auth1.Version)
	ri2, _ := rentities.NewRegisterInfo(auth2.Name, auth2.IID, auth2.IP, auth2.Version)
	req, _ := rentities.NewServiceRequest("Auth", "v1.0.0")*/

	//test auth as server
	auth1.GetConnection("172.25.0.2", config.ClientPort)

	//test auth as client
	/*auth1.ListAllServices()
	auth1.Register(ri1)
	auth2.Register(ri2)
	auth1.ListAllServices()
	auth1.FindService(*req)*/
}
