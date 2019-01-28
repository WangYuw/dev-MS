package client

import "rentities"

//Client is an interface of client
type Client interface {
	//SendQRequest registry sends quality request to services(qualityrequest, ip, port)
	SendQRequest(*rentities.SQualityReq, string, string) error
}
