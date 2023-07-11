package handlers

import (
	"go-risky/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller PublicController) GetCapabilities(context *gin.Context) {

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

	capabilityOutputs, err := controller.DBManager.GetCapabilities(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, capabilityOutputs)
		return
	}

	context.JSON(http.StatusOK, capabilityOutputs)
}

func (controller PublicController) GetCapability(context *gin.Context) {

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

	capabilityOutput, err := controller.DBManager.GetCapability(capabilityId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, capabilityOutput)
		return
	}

	context.JSON(http.StatusOK, capabilityOutput)
}

func (controller PublicController) DeleteCapability(context *gin.Context) {

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

	err = controller.DBManager.DeleteCapability(capabilityId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateCapability(context *gin.Context) {

	capabilityInput := types.Capability{}
	err := context.ShouldBindJSON(&capabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	capabilityId, err := controller.DBManager.CreateCapability(capabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, capabilityId)
}

func (controller PublicController) UpdateCapability(context *gin.Context) {

	capabilityInput := types.Capability{}
	err := context.ShouldBindJSON(&capabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	err = controller.DBManager.UpdateCapability(capabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
