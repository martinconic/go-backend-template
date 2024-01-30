package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/martinconic/go-backend-template/config"
	"github.com/martinconic/go-backend-template/storage"
	"github.com/martinconic/go-backend-template/storage/psqldb"
	"github.com/martinconic/go-backend-template/utils/constants"
	"github.com/spf13/viper"
)

type Server struct {
	Router *gin.Engine
	db     storage.Database
}

var server *Server

func StartServer(v *viper.Viper) {
	server = &Server{}
	server.Initialize(v)
	server.Run(v.GetString(constants.ApiServer))
}

func (server *Server) Initialize(v *viper.Viper) {
	var err error
	server.Router = gin.Default()
	if err != nil {
		log.Println(err)
	}
	server.db, err = psqldb.NewDatabase(config.GetPostgresConfig(v))
	if err != nil {
		log.Println(err)
	}
	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	err := server.Router.Run(":" + addr)
	if err != nil {
		panic("unable to start server!")
	}
}
