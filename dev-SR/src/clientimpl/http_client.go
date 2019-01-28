package clientimpl

import (
	"bytes"
	"config"
	"db"
	"encoding/json"
	"fmt"
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
func (hc HTTPClient) SendQRequest(sqr *rentities.SQualityReq, ip string, port string) error {
	//New SQR and to Json
	sqrJSON, err := json.Marshal(sqr)
	if err != nil {
		return fmt.Errorf("json.Marshal() failed with '%s'", err)
	}
	//New http client
	hc.HClient.Timeout = time.Second * 15
	//New http request
	uri := "http://" + ip + ":" + port + "/services/quality"
	body := bytes.NewBuffer(sqrJSON)
	req, err := http.NewRequest(http.MethodPost, uri, body)
	if err != nil {
		return fmt.Errorf("http.NewRequest() failed with '%s'", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//Get http response
	resp, err := hc.HClient.Do(req)
	if err != nil {
		return fmt.Errorf("client.Do() failed with '%s'", err)
	}
	defer resp.Body.Close()

	//handle response
	fmt.Printf("Response status code: %d, text:\n", resp.StatusCode)

	var sqi rentities.SQualityInfo
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&sqi)
	if err != nil {
		return fmt.Errorf("JSON Decoder error")
	}

	r := registry.NewRegistry()
	pdb, err := db.NewPostgres(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	err = r.UpdateSQ(pdb, &sqi)
	fmt.Printf("resp status code: %d\n", resp.StatusCode)
	if err != nil {
		return err
	}
	fmt.Printf("After updating load: T_Name: %s, I_ID: %d, Load: %f\n", sqi.TName, sqi.IID, sqi.Quality.Load)
	return nil
}
