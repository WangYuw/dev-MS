package client

//Client is an interface of client
type Client interface {
	//SendQRequest registry sends quality request to services(name, version, ip, port)
	SendQRequest(string, string, string, string)
}
