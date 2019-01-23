package main

import (
	"config"
	"services"
)

func main() {
	auth1, _ := services.NewAuth("auth1", 1, "v1.0.0")

	//test auth as server
	auth1.GetConnection(auth1.IP, config.ClientPort)
}
