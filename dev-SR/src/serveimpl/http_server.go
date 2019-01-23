package serveimpl

import (
	"config"
	"db"
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
	pdb, err := db.NewPostgres(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	mux.Handle("/services/all", ShowHandler(r, pdb))
	mux.Handle("/services/registry", SRegisterHandler(r, pdb))
	mux.Handle("/services/info", ServiceInfoHandler(r, pdb))

	log.Printf("http server serving on %s port %s", ip, port)
	log.Fatal(http.ListenAndServe(ip+":"+port, mux))
}
