package serveimpl

import (
	"db"
	"encoding/json"
	"log"
	"net/http"
	"registry"
	"rentities"
)

//ShowHandler is a handler func to show all services
func ShowHandler(reg *registry.Registry, pdb *db.PostgresDB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		reponse, err := reg.GetAll(pdb)
		if err != nil {
			log.Println(err)
			return
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(reponse)
	})
}

//SRegisterHandler is a handler func to register services
func SRegisterHandler(db *registry.Registry, pdb *db.PostgresDB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		var request rentities.RegisterInfo
		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			log.Println("JSON Decoder error")
			return
		}
		err = db.Register(pdb, &request)
		if err != nil {
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

//ServiceInfoHandler is a handler func to find services
func ServiceInfoHandler(db *registry.Registry, pdb *db.PostgresDB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		var request rentities.ServiceRequest
		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			log.Println("JSON Decoder error")
			return
		}
		reponse, err := db.FindService(pdb, request)
		if err != nil {
			log.Println(err)
			return
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(reponse)
	})
}

//SQualityHandler is a handler func to handle service quality
func SQualityHandler(db *registry.Registry, pdb *db.PostgresDB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		var request rentities.SQualityInfo
		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			log.Println("JSON Decoder error")
			return
		}
		err = db.UpdateSQ(pdb, &request)
		if err != nil {
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
