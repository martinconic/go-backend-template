package storage

import "github.com/martinconic/go-backend-template/data"

type Database interface {
	Get() ([]data.Data, error)
	Insert(data string) (int64, error)
	Update(data *data.Data) (string, error)
	Delete(data *data.Data) (string, error)

	GetDBStatus() error
}
