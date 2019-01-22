package db

import (
	"database/sql"
	"fmt"
	"log"
	"rentities"

	//drive postfresql
	_ "github.com/lib/pq"
)

//PostgresDB is repository of postgresql db
type PostgresDB struct {
	db *sql.DB
}

//NewPostgres inits db
func NewPostgres(user string, password string, name string) (*PostgresDB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, name)
	//dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, name)
	db, err := sql.Open("postgres", dbinfo)
	fmt.Println("# Postgres opened ...")
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	//defer db.Close()
	return &PostgresDB{
		db,
	}, nil
}

//Close closes db
func (p *PostgresDB) Close() {
	p.db.Close()
}

//InsertReg inserts data
func (p *PostgresDB) InsertReg(reg rentities.RegisterInfo) error {
	fmt.Println("# Inserting RegisterInfo ...")
	_, err := p.db.Exec("INSERT INTO registry(type_name, instance_id, ip, version, load) VALUES($1, $2, $3, $4, $5)", reg.TName, reg.IID, reg.IP, reg.Version, reg.Quality.Load)
	return err
}

//ListRegs finds all values in registry db
func (p *PostgresDB) ListRegs() ([]rentities.RegisterInfo, error) {
	fmt.Println("# Quarying All ...")
	rows, err := p.db.Query("SELECT * FROM registry ORDER BY type_name, instance_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//Parse all rows into a slice of RegisterInfo
	regs := []rentities.RegisterInfo{}
	for rows.Next() {
		q := &rentities.ServiceQuality{}
		reg := rentities.RegisterInfo{Quality: q}
		err = rows.Scan(&reg.TName, &reg.IID, &reg.IP, &reg.Version, &reg.Quality.Load)
		if err == nil {
			regs = append(regs, reg)
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return regs, nil
}

//FindRegByID find a RegisterInfo
func (p *PostgresDB) FindRegByID(id int64) (rentities.RegisterInfo, error) {
	fmt.Println("# Querying By ID...")
	q := &rentities.ServiceQuality{}
	reg := rentities.RegisterInfo{Quality: q}
	err := p.db.QueryRow("SELECT type_name, instance_id, ip, version, load FROM registry WHERE id = $1", id).Scan(&reg.TName, &reg.IID, &reg.IP, &reg.Version, &reg.Quality.Load)
	return reg, err
}

//FindMinLoadSrv find a RegisterInfo min load
func (p *PostgresDB) FindMinLoadSrv(name string, version string) (rentities.RegisterInfo, error) {
	fmt.Println("# Querying Srv...")
	q := &rentities.ServiceQuality{}
	reg := rentities.RegisterInfo{Quality: q}
	statement := "SELECT type_name, instance_id, ip, version, load FROM registry WHERE type_name = $1 AND version = $2 and load=(SELECT MIN(load) FROM registry GROUP BY type_name)"
	err := p.db.QueryRow(statement, name, version).Scan(&reg.TName, &reg.IID, &reg.IP, &reg.Version, &reg.Quality.Load)
	//err := p.db.QueryRow("SELECT type_name, instance_id, ip, version, load FROM registry WHERE type_name = $1 AND version = $2", name, version).Scan(&reg.TName, &reg.IID, &reg.IP, &reg.Version, &reg.Quality.Load)
	return reg, err
}

//UpdateLoad updates load
func (p *PostgresDB) UpdateLoad(id int64, load float32) error {
	fmt.Println("# Updating load ...")
	_, err := p.db.Exec("UPDATE registry SET load = $2 WHERE id = $1", id, load)
	return err
}

//DeleteReg deletes reg
func (p *PostgresDB) DeleteReg(id int64) error {
	fmt.Println("# Deleting ...")
	_, err := p.db.Exec("DELETE FROM registry WHERE id = $1", id)
	return err
}
