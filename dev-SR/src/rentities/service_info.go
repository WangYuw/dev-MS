package rentities

import (
	"config"
	"fmt"
)

//ServiceInfo is the response of service request
type ServiceInfo struct {
	TName   string `json:"type_name"`
	IID     int64  `json:"instance_id"`
	IP      string `json:"ip"`
	Version string `json:"version"`
	TTL     int    `json:"ttl"`
}

//NewServiceInfo constructs a response
func NewServiceInfo(n string, id int64, ip string, v string, ttl int) (*ServiceInfo, error) {
	if n == "" {
		return nil, fmt.Errorf("NewServiceInfo: service's name miss")
	}
	if id <= 0 {
		return nil, fmt.Errorf("NewServiceInfo: istance id error")
	}
	if ip == "" {
		return nil, fmt.Errorf("NewServiceInfo: service's ip miss")
	}
	if v == "" {
		v = config.DefaultVersion
	}
	resp := &ServiceInfo{
		TName:   n,
		IID:     id,
		IP:      ip,
		Version: v,
		TTL:     ttl,
	}
	return resp, nil
}
