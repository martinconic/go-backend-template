package main

import (
	"fmt"
	"os"

	"github.com/martinconic/go-backend-template/api"
	"github.com/martinconic/go-backend-template/config"
)

func main() {

	v, err := config.NewViper()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	api.StartServer(v)

}
