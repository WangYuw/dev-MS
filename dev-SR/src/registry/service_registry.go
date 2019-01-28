package registry

import (
	"client"
	"config"
	"db"
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

//GetAll all instances of services from db
func (r *Registry) GetAll(pdb *db.PostgresDB) ([]rentities.RegisterInfo, error) {
	return pdb.ListRegs()
	//return r.ServicesMap, nil
}

//Register adds RegisterInfo into the map of services in service registry
func (r *Registry) Register(pdb *db.PostgresDB, info *rentities.RegisterInfo) error {
	sList, ok := r.ServicesMap[info.TName]
	if ok {
		for _, ri := range sList {
			if ri.IID == info.IID || ri.IP == info.IP {
				return fmt.Errorf("Register error: service instance exists")
			}
		}
	}
	r.ServicesMap[info.TName] = append(r.ServicesMap[info.TName], info)
	//add into db
	pdb.InsertReg(*info)
	return nil
}

//Unregister deletes RegisterInfo from the map of services in service registry
func (r *Registry) Unregister(pdb *db.PostgresDB, info *rentities.RegisterInfo) error {
	sList, okName := r.ServicesMap[info.TName]
	if okName != true {
		return fmt.Errorf("Unregister error: service not exist")
	}
	for i, ri := range sList {
		if ri.IP == info.IP { //delete elem in slice
			sList[i] = sList[len(sList)-1]
			sList[len(sList)-1] = nil
			sList = sList[:len(sList)-1]
			//delete from db
			pdb.DeleteReg(info.IID)
			return nil
		}
	}
	return fmt.Errorf("Deregister error: service not exist")
}

//FindService finds required service
func (r *Registry) FindService(pdb *db.PostgresDB, req rentities.ServiceRequest) (*rentities.ServiceInfo, error) {
	/*sList, okName := r.ServicesMap[req.TName]
	if okName != true {
		return nil, fmt.Errorf("FindService error: service not exist")
	}

	//get min load instance
	minLoad := sList[0].Quality.Load
	minIndex := 0
	for i, ri := range sList {
		if ri.Version == req.Version {
			if minLoad >= ri.Quality.Load {
				minIndex = i
			}
		}
	}
	*/
	srv, err := pdb.FindMinLoadSrv(req.TName, req.Version)
	ri, err := rentities.NewServiceInfo(srv.TName, srv.IID, srv.IP, srv.Version, config.DefaultTTL)
	/*srv, err := rentities.NewServiceInfo(sList[minIndex].TName, sList[minIndex].IID,
	sList[minIndex].IP, sList[minIndex].Version, config.DefaultTTL)*/

	if err != nil {
		return nil, err
	}
	return ri, nil
}

//UpdateSQ updates service quality
func (r *Registry) UpdateSQ(pdb *db.PostgresDB, resp *rentities.SQualityInfo) error {
	//update db
	return pdb.UpdateLoad(resp.IID, resp.Quality.Load)

	/*
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
	*/
}

//SendQRequest sends quality request to all services
func (r *Registry) SendQRequest(c client.Client, sqr *rentities.SQualityReq, ip string, port string) error {
	/*pdb, err := db.NewPostgres(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	srvs, err := pdb.ListRegs()
	if err != nil {
		return err
	}
	for _, srv := range srvs {
		c.SendQRequest(srv.TName, srv.Version, srv.IP, port)
	}
	return nil*/
	return c.SendQRequest(sqr, ip, port)
}
