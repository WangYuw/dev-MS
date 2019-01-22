package clientimpl

import (
	"bytes"
	"config"
	"db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"registry"
	"rentities"
	"time"
)

//HTTPClient is a http client
type HTTPClient struct {
	HClient *http.Client
}

//NewClient creates a new HTTPClient
func NewClient() *HTTPClient {
	return &HTTPClient{
		HClient: &http.Client{},
	}
}

//SendQRequest (impl. Client) (service name, service version)
func (hc HTTPClient) SendQRequest(sn string, sv string, ip string, port string) {
	//New SQR and to Json
	sqr, _ := rentities.NewSQR(sn, sv)
	sqrJSON, err := json.Marshal(sqr)
	if err != nil {
		log.Fatalf("json.Marshal() failed with '%s'\n", err)
	}
	//New http client
	hc.HClient.Timeout = time.Second * 15
	//New http request
	uri := "http://" + ip + ":" + port + "/services/quality"
	body := bytes.NewBuffer(sqrJSON)
	req, err := http.NewRequest(http.MethodPost, uri, body)
	if err != nil {
		log.Fatalf("http.NewRequest() failed with '%s'\n", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//Get http response
	resp, err := hc.HClient.Do(req)
	if err != nil {
		log.Fatalf("client.Do() failed with '%s'\n", err)
	}
	defer resp.Body.Close()

	//handle response
	fmt.Printf("Response status code: %d, text:\n", resp.StatusCode)

	var sqi rentities.SQualityInfo
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&sqi)
	if err != nil {
		log.Println("JSON Decoder error")
		return
	}
	r := registry.NewRegistry()
	pdb, err := db.NewPostgres(config.DBUser, config.DBPassword, config.DBName)

	fmt.Printf("%s %d %f\n", sqi.TName, sqi.IID, sqi.Quality.Load)

	err = r.UpdateSQ(pdb, &sqi)
	fmt.Printf("resp status code: %d\n", resp.StatusCode)
	if err != nil {
		log.Println(err)
		return
	}
}
