package mitigation

import (
	"go-risky/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getMitigations(context *gin.Context) {
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

	mitigationOutput, err := database.GetMitigations(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, mitigationOutput)
		return
	}

	context.JSON(http.StatusOK, mitigationOutput)
}

func getMitigation(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	mitigationId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	mitigationOutput, err := database.GetMitigation(mitigationId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, mitigationOutput)
		return
	}

	context.JSON(http.StatusOK, mitigationOutput)
}

func deleteMitigation(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	mitigationId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = database.DeleteMitigation(mitigationId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createMitigation(context *gin.Context) {
	mitigationInput := database.Mitigation{}
	err := context.ShouldBindJSON(&mitigationInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.CreateMitigation(mitigationInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateMitigation(context *gin.Context) {
	mitigationInput := database.Mitigation{}
	err := context.ShouldBindJSON(&mitigationInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.UpdateMitigation(mitigationInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func MitigationRoutes(router *gin.Engine) {
	router.GET("/mitigations", getMitigations)
	router.GET("/mitigation/:id", getMitigation)
	router.DELETE("/mitigation/:id", deleteMitigation)
	router.PATCH("/mitigation/:id", updateMitigation)
	router.POST("/mitigations", createMitigation)
}
