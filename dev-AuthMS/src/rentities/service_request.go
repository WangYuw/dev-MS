package rentities

import (
	"config"
	"fmt"
)

//ServiceRequest is a request for searching services
type ServiceRequest struct {
	TName   string `json:"type_name"`
	Version string `json:"version"`
}

//NewServiceRequest constructs a service request
func NewServiceRequest(n string, v string) (*ServiceRequest, error) {
	if n == "" {
		return nil, fmt.Errorf("NewServiceRequest: service's name miss")
	}
	if v == "" {
		v = config.DefaultVersion
	}
	req := &ServiceRequest{
		TName:   n,
		Version: v,
	}
	return req, nil
}
