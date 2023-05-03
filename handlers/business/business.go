package business

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type business struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Revenue   float32   `json:"revenue"`
	CreatedAt time.Time `json:"createdAt"`
}

var businesses = []business{{}}

func getBusinesses(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, businesses)
}

func getBusiness(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, businesses)
}

func deleteBusiness(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, businesses)
}

func createBusiness(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, businesses)
}

func updateBusiness(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, businesses)
}

func BusinessRoutes(router *gin.Engine) {
	router.GET("/businesses", getBusinesses)
	router.GET("/business/:id", getBusiness)
	router.DELETE("/business/:id", deleteBusiness)
	router.PATCH("/business/:id", updateBusiness)
	router.POST("/businesses", createBusiness)
}
