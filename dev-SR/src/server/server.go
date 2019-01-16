package server

//Server is an interface of server
type Server interface {
	//NewServer(string, string) *Server //protocol, port
	NewConnection(ip string, port string)
}
