package main

import (
	"fmt"

	"github.com/NaphatBuranathaworn/go-gin/config"
	"github.com/NaphatBuranathaworn/go-gin/models"
	"github.com/NaphatBuranathaworn/go-gin/routes"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)


var err error

func main() {
    var dsn = config.DSNPostgres(config.BuildDBConfig())
    fmt.Println("db url : ", dsn)

    config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	// defer config.DB.Close()

	config.DB.AutoMigrate(&models.Todo{})

	r := routes.SetupRouter()

	r.Run()
}