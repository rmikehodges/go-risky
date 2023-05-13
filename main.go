package main

import (
	"go-risky/riskyrouter"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	riskyrouter.InitializeRouter(router)

	router.Use(secure.Secure(secure.Options{
		ContentSecurityPolicy: "default-src 'self' http://localhost:3000",
	}))

	router.Run(":8081")
}
