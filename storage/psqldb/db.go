package psqldb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/martinconic/go-backend-template/config"
	"github.com/martinconic/go-backend-template/data"
	"github.com/martinconic/go-backend-template/storage"
)

type PostgresDB struct {
	Config *config.PostgresConfig
	DB     *sql.DB
}

func NewDatabase(config *config.PostgresConfig) (storage.Database, error) {
	db := &PostgresDB{
		Config: config,
	}
	error := db.getDBConnection()

	return db, error
}

func (d *PostgresDB) getDBConnection() error {
	var err error

	if d.DB != nil {
		return err
	}

	dataSource := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", d.Config.Host, d.Config.Port, d.Config.User,
		d.Config.Password, d.Config.Name)
	d.DB, err = sql.Open("postgres", dataSource)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (p *PostgresDB) GetDBStatus() error {
	var err error
	if err = p.DB.Ping(); err != nil {
		p.DB.Close()
	}
	return err
}

func (p *PostgresDB) Get() ([]data.Data, error) {
	sc := []data.Data{}
	rows, err := p.DB.Query("SELECT id, description FROM data")
	if err != nil {
		return sc, err
	}

	for rows.Next() {
		var c data.Data
		err = rows.Scan(
			&c.ID,
			&c.Description,
		)
		if err != nil {
			log.Println(err)
		}

		sc = append(sc, c)
	}

	return sc, nil
}

func (p *PostgresDB) Insert(data string) (int64, error) {
	sql := ` INSERT into data (description) VALUES($1)`
	result, err := p.DB.Exec(sql, data)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return result.LastInsertId()

}

func (p *PostgresDB) Update(data *data.Data) (string, error) {
	return "Success", nil
}

func (p *PostgresDB) Delete(data *data.Data) (string, error) {
	return "Success", nil
}
