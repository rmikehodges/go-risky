package riskyrouter

import (
	"context"
	"go-risky/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeRouter() (router *gin.Engine) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}

	router = gin.Default()
	router.Use(dbManager.Handle())
	router.Use(cors.Default())

	router.Use(secure.Secure(secure.Options{
		ContentSecurityPolicy: "default-src 'self' http://localhost:3000",
	}))
	return
}
