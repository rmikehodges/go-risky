package main

import (
	"context"
	"go-risky/database"
	"go-risky/riskyrouter"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Initialize PGX pool
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}

	router := gin.Default()
	router.Use(dbManager.Handle())
	router.Use(cors.Default())
	riskyrouter.InitializeRouter(router)

	router.Use(secure.Secure(secure.Options{
		ContentSecurityPolicy: "default-src 'self' http://localhost:3000",
	}))

	router.Run(":8081")
}
