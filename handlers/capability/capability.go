package capability

import (
	"go-risky/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getCapabilities(context *gin.Context) {
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

	capabilityOutput, err := database.GetCapabilities(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, capabilityOutput)
		return
	}

	context.JSON(http.StatusOK, capabilityOutput)
}

func getCapability(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	capabilityId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	capabilityOutput, err := database.GetCapability(capabilityId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, capabilityOutput)
		return
	}

	context.JSON(http.StatusOK, capabilityOutput)
}

func deleteCapability(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	capabilityId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = database.DeleteCapability(capabilityId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createCapability(context *gin.Context) {
	capabilityInput := database.Capability{}
	err := context.ShouldBindJSON(&capabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.CreateCapability(capabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateCapability(context *gin.Context) {
	capabilityInput := database.Capability{}
	err := context.ShouldBindJSON(&capabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.UpdateCapability(capabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func CapabilityRoutes(router *gin.Engine) {
	router.GET("/capabilities", getCapabilities)
	router.GET("/capability", getCapability)
	router.DELETE("/capability", deleteCapability)
	router.PATCH("/capability", updateCapability)
	router.POST("/capabilities", createCapability)
}
