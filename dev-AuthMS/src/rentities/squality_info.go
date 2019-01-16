package rentities

import (
	"fmt"
)

//SQualityInfo is the reponse of service quality request
type SQualityInfo struct {
	TName   string          `json:"type_name"`
	IID     int64           `json:"instance_id"`
	IP      string          `json:"ip"`
	Quality *ServiceQuality `json:"service_quality"`
}

//NewSQI constructs a response
func NewSQI(n string, id int64, ip string, q *ServiceQuality) (*SQualityInfo, error) {
	if n == "" {
		return nil, fmt.Errorf("NewSQI: service's name miss")
	}
	if id <= 0 {
		return nil, fmt.Errorf("NewSQI: istance id error")
	}
	if ip == "" {
		return nil, fmt.Errorf("NewSQI: service's ip miss")
	}
	resp := &SQualityInfo{
		TName:   n,
		IID:     id,
		IP:      ip,
		Quality: q,
	}
	return resp, nil
}
