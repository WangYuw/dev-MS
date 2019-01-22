package services

import (
	"bytes"
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"rentities"
	"time"
)

//Auth is an authentification service
type Auth struct {
	//metadata
	Name    string
	IID     int64
	IP      string
	Version string
	Quality *rentities.ServiceQuality
	//function ...
}

//NewAuth constructs a Auth service
func NewAuth(ip string, id int64, v string) (*Auth, error) {
	if ip == "" {
		return nil, fmt.Errorf("NewAuth: Auth's ip miss")
	}
	if v == "" {
		v = config.DefaultVersion
	}
	//random load
	rand.Seed(time.Now().UnixNano())
	load := rand.Float32()

	resp := &Auth{
		Name:    "Auth",
		IID:     id,
		IP:      ip,
		Version: v,
		Quality: rentities.NewServiceQuality(load),
	}
	return resp, nil
}

//GetConnection gets http connection (server)
func (a *Auth) GetConnection(ip string, port string) {
	log.Printf("service auth serving on %s port %s", ip, port)
	//connect db
	/*pdb, err := db.NewPostgres(config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		panic(err)
	}*/
	http.Handle("/services/quality", a.qualityHandler())
	log.Fatal(http.ListenAndServe(ip+":"+port, nil))

}

func (a *Auth) qualityHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		reponse, err := rentities.NewSQI(a.Name, a.IID, a.IP, a.Quality)
		if err != nil {
			log.Println(err)
			return
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(reponse)
	})
}

//ListAllServices sends request to list all services
func (a *Auth) ListAllServices(uri string) {
	resp, err := http.Get(uri + "/services/all")
	if err != nil {
		log.Fatalf("http.Get() failed with '%s'\n", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

//Register registers service
func (a *Auth) Register(info *rentities.RegisterInfo, uri string) {
	infoJ, err := json.Marshal(info)
	if err != nil {
		log.Fatalf("json.Marshal() failed with '%s'\n", err)
	}
	body := bytes.NewBuffer(infoJ)
	req, err := http.NewRequest(http.MethodPost, uri+"/services/registry", body)
	if err != nil {
		log.Fatalf("http.NewRequest() failed with '%s'\n", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client.Do() failed with '%s'\n", err)
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}

//FindService finds service
func (a *Auth) FindService(reqs rentities.ServiceRequest, uri string) {
	reqsJ, err := json.Marshal(reqs)
	if err != nil {
		log.Fatalf("json.Marshal() failed with '%s'\n", err)
	}
	body := bytes.NewBuffer(reqsJ)
	req, err := http.NewRequest(http.MethodPost, uri+"/services/info", body)
	if err != nil {
		log.Fatalf("http.NewRequest() failed with '%s'\n", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client.Do() failed with '%s'\n", err)
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}
