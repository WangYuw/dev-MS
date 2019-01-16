package rentities

//SQualityReq is service quality request, service registery -> services
type SQualityReq struct {
	CommonReq *ServiceRequest
}

//NewSQR constructs a service quality request (service name, service version)
func NewSQR(n string, v string) (*SQualityReq, error) {
	req, err := NewServiceRequest(n, v)
	return &SQualityReq{
		CommonReq: req,
	}, err
}
