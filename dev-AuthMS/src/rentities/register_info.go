package rentities

import (
	"config"
	"fmt"
)

//RegisterInfo gives the information of registering service
type RegisterInfo struct {
	TName   string          `json:"type_name"`
	IID     int64           `json:"instance_id"`
	IP      string          `json:"ip"`
	Version string          `json:"version"`
	Quality *ServiceQuality `json:"quality"`
}

//NewRegisterInfo constructs a register info
func NewRegisterInfo(n string, id int64, ip string, v string, q *ServiceQuality) (*RegisterInfo, error) {
	if n == "" {
		return nil, fmt.Errorf("NewRegisterInfo: service's name miss")
	}
	if id <= 0 {
		return nil, fmt.Errorf("NewRegisterInfo: istance id error")
	}
	if ip == "" {
		return nil, fmt.Errorf("NewRegisterInfo: service's ip miss")
	}
	if v == "" {
		v = config.DefaultVersion
	}

	ri := &RegisterInfo{
		TName:   n,
		IID:     id,
		IP:      ip,
		Version: v,
		Quality: q,
	}
	return ri, nil
}
