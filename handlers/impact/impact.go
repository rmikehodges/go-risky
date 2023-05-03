package impact

import (
	"go-risky/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getImpacts(context *gin.Context) {
	id, ok := context.GetQuery("businessId")
	if !ok {
		log.Println("Parameter businessId not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	businessId, err := uuid.Parse(id)
	if err != nil {
		log.Println("businessId is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	impactOutput, err := database.GetImpacts(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, impactOutput)
		return
	}

	context.JSON(http.StatusOK, impactOutput)
}

func getImpact(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	impactId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	impactOutput, err := database.GetImpact(impactId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, impactOutput)
		return
	}

	context.JSON(http.StatusOK, impactOutput)
}

func deleteImpact(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	impactId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = database.DeleteImpact(impactId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createImpact(context *gin.Context) {
	impactInput := database.Impact{}
	err := context.ShouldBindJSON(&impactInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.CreateImpact(impactInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateImpact(context *gin.Context) {
	impactInput := database.Impact{}
	err := context.ShouldBindJSON(&impactInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.UpdateImpact(impactInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func ImpactRoutes(router *gin.Engine) {
	router.GET("/impacts", getImpacts)
	router.GET("/impact/:id", getImpact)
	router.DELETE("/impact/:id", deleteImpact)
	router.PATCH("/impact/:id", updateImpact)
	router.POST("/impacts", createImpact)
}
