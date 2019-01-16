package serveimpl

import (
	"log"
	"net/http"
	"registry"
)

//HTTPServer is a http server implementing Server
type HTTPServer struct {
	Type string
}

//NewServer constructs a new http server on Port
func NewServer(t string) *HTTPServer {
	return &HTTPServer{
		Type: t,
	}
}

//NewConnection news a connection (impl. Server)
func (hs HTTPServer) NewConnection(ip string, port string) {
	r := registry.NewRegistry()
	mux := http.NewServeMux()
	mux.Handle("/services/all", ShowHandler(r))
	mux.Handle("/services/registry", SRegisterHandler(r))
	mux.Handle("/services/info", ServiceInfoHandler(r))
	//mux.Handle("/services/quality", SQualityHandler(r))

	log.Printf("http server serving on %s port %s", ip, port)
	log.Fatal(http.ListenAndServe(ip+":"+port, mux))
}
