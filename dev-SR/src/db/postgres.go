package postgresql

import (
	"database/sql"
	"fmt"
	"rentities"

	_ "github.com/lib/pq"
)

//DBUser user of db
const DBUser string = "postgres"

//DBPassword password of db
const DBPassword string = "postgres"

//DBName name of db
const DBName string = "registry"

//PostgresDB is repository of postgresql db
type PostgresDB struct {
	db *sql.DB
}

//NewPostgres inits db
func NewPostgres(user string, password string, name string) (*PostgresDB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, name)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}
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
	rows, err := p.db.Query("SELECT * FROM registry")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//Parse all rows into a slice of RegisterInfo
	regs := []rentities.RegisterInfo{}
	for rows.Next() {
		reg := rentities.RegisterInfo{}
		if err = rows.Scan(&reg.TName, &reg.IID, &reg.IP, &reg.Version, &reg.Quality.Load); err == nil {
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
	fmt.Println("# Querying ...")
	reg := rentities.RegisterInfo{}
	err := p.db.Query("SELECT type_name, instance_id, ip, version, load FROM registry WHERE id = $1", id).Scan(&reg.TName, &reg.IID, &reg.IP, &reg.Version, &reg.Quality.Load)
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return reg, nil
}

//UpdateLoad updates load
func (p *PostgresDB) UpdateLoad(id string, load float32) error {
	fmt.Println("# Updating load ...")
	_, err := p.db.Exec("UPDATE registry SET load = $2 WHERE id = $1", id, load)
	return err
}

//DeleteReg deletes reg
func (p *PostgresDB) DeleteReg(id int64) error {
	fmt.Println("# Deleting ...")
	_, err = p.db.Exec("DELETE FROM registry WHERE id = $1", id)
	return err
}
