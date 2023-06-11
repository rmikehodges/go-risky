package main

import (
	"context"
	"go-risky/database"
	"go-risky/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	controller := &handlers.PublicController{}
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}

	controller.DBManager = dbManager

	router := gin.Default()

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowHeaders = []string{"Sec-Fetch-Dest", "Sec-Fetch-Mode", "Sec-Fetch-Site"}
	corsConfig.AllowAllOrigins = true

	router.Use(cors.New(corsConfig))

	controller.RegisterRoutes(router)

	// router.Use(secure.Secure(secure.Options{
	// 	ContentSecurityPolicy: "http://localhost:3000",
	// 	IsDevelopment:         true,
	// }))

	router.Run(":8081")
}
