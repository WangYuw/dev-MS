package rentities

//ServiceQuality is service quality request, request of service registery asks services
type ServiceQuality struct {
	Load float32 `json:"load"` //random [0.0,1.0)
}

//NewServiceQuality constructs a service quality request
func NewServiceQuality(l float32) *ServiceQuality {
	sq := &ServiceQuality{
		Load: l,
	}
	return sq
}
