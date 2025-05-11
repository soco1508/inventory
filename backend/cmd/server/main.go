package main

import (
	"backend/config"
	"backend/internal/api/routes"
	"backend/pkg/db"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	app := gin.Default()
	config, err := config.NewParsedConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}

	dbConfig := db.DBConfig{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		Username: config.Database.Username,
		Password: config.Database.Password,
		Name:     config.Database.Name,
	}

	sqlxDb, err := db.SqlxInitDB(dbConfig)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer sqlxDb.Close()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowOrigins = []string{"*"}
	corsCfg.AllowCredentials = true
	corsCfg.AllowHeaders = []string{"*"}

	app.Use(cors.New(corsCfg))
	routes.RegisterDashboardRoutes(app, sqlxDb)

	if err = app.Run(config.BaseUrl + ":" + config.ServerPort); err != nil {
		log.Fatalf("Could not start the server: %v", err)
	}
}
