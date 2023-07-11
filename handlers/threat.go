package handlers

import (
	"go-risky/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller PublicController) GetThreats(context *gin.Context) {

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

	threatOutputs, err := controller.DBManager.GetThreats(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, threatOutputs)
		return
	}

	context.JSON(http.StatusOK, threatOutputs)
}

func (controller PublicController) GetThreat(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	threatId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	threatOutput, err := controller.DBManager.GetThreat(threatId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, nil)
		return
	}

	context.JSON(http.StatusOK, threatOutput)
}

func (controller PublicController) DeleteThreat(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	threatId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = controller.DBManager.DeleteThreat(threatId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateThreat(context *gin.Context) {

	threatInput := types.Threat{}
	err := context.ShouldBindJSON(&threatInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	threatId, err := controller.DBManager.CreateThreat(threatInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, threatId)
}

func (controller PublicController) UpdateThreat(context *gin.Context) {

	threatInput := types.Threat{}
	err := context.ShouldBindJSON(&threatInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	err = controller.DBManager.UpdateThreat(threatInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
