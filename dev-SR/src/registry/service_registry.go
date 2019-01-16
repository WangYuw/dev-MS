package registry

import (
	"client"
	"config"
	"fmt"
	"rentities"
	"server"
)

//ServiceList is a list of services
type ServiceList []*rentities.RegisterInfo //map[string]*RegisterInfo //Slice???

//Registry is a service registry
type Registry struct {
	//map of slice {"Auth":{{n, ip, v}, {n, ip, v}}, "User": {...}, ...}
	ServicesMap map[string]ServiceList
}

//NewRegistry returns a new Registry
func NewRegistry() *Registry {
	return &Registry{
		ServicesMap: make(map[string]ServiceList),
	}
}

//GetConnection get net connection
func (r *Registry) GetConnection(s server.Server, ip string, port string) {
	s.NewConnection(ip, port)
}

//GetAll all instances of services
func (r *Registry) GetAll() (map[string]ServiceList, error) {
	return r.ServicesMap, nil
}

//Register adds RegisterInfo into the map of services in service registry
func (r *Registry) Register(info *rentities.RegisterInfo) error {
	sList, ok := r.ServicesMap[info.TName]
	if ok {
		for _, ri := range sList {
			if ri.IID == info.IID || ri.IP == info.IP {
				return fmt.Errorf("Register error: service instance exists")
			}
		}
	}
	r.ServicesMap[info.TName] = append(r.ServicesMap[info.TName], info)
	return nil
}

//Unregister deletes RegisterInfo from the map of services in service registry
func (r *Registry) Unregister(info *rentities.RegisterInfo) error {
	sList, okName := r.ServicesMap[info.TName]
	if okName != true {
		return fmt.Errorf("Unregister error: service not exist")
	}
	for i, ri := range sList {
		if ri.IP == info.IP { //delete elem in slice
			sList[i] = sList[len(sList)-1]
			sList[len(sList)-1] = nil
			sList = sList[:len(sList)-1]
			return nil
		}
	}
	return fmt.Errorf("Deregister error: service not exist")
}

//FindService finds required service
func (r *Registry) FindService(req rentities.ServiceRequest) (*rentities.ServiceInfo, error) {
	sList, okName := r.ServicesMap[req.TName]
	if okName != true {
		return nil, fmt.Errorf("FindService error: service not exist")
	}
	minLoad := sList[0].Quality.Load
	minIndex := 0
	for i, ri := range sList {
		if ri.Version == req.Version {
			if minLoad >= ri.Quality.Load {
				minIndex = i
			}
		}
	}
	serv, err := rentities.NewServiceInfo(sList[minIndex].TName, sList[minIndex].IID,
		sList[minIndex].IP, sList[minIndex].Version, config.DefaultTTL)
	if err != nil {
		return nil, err
	}
	return serv, nil
}

//UpdateSQ updates serv=>ice quality
func (r *Registry) UpdateSQ(resp *rentities.SQualityInfo) error {
	_, okName := r.ServicesMap[resp.TName]
	if okName != true {
		return fmt.Errorf("UpdateSQ erroSr: service not exist")
	}
	for _, ri := range r.ServicesMap[resp.TName] {
		if ri.IID == resp.IID && ri.IP == resp.IP {
			ri.Quality.Load = resp.Quality.Load
			return nil
		}
	}
	return fmt.Errorf("UpdateSQ error: service instance not exist")
}

//SendQRequests sends quality request to all services
func (r *Registry) SendQRequests(c client.Client, port string) {
	for _, sList := range r.ServicesMap {
		for _, s := range sList {
			c.SendQRequest(s.TName, s.Version, s.IP, port)
		}
	}
}
