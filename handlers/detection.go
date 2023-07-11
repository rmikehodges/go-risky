package handlers

import (
	"go-risky/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//Create types types.Detection and DetectionOutput that match the database model DetectionModel

func (controller PublicController) GetDetections(context *gin.Context) {

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

	detectionOutput, err := controller.DBManager.GetDetections(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, detectionOutput)
		return
	}

	context.JSON(http.StatusOK, detectionOutput)
}

func (controller PublicController) GetDetection(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	detectionId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	detectionOutput, err := controller.DBManager.GetDetection(detectionId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, detectionOutput)
		return
	}

	context.JSON(http.StatusOK, detectionOutput)
}

func (controller PublicController) DeleteDetection(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	detectionId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = controller.DBManager.DeleteDetection(detectionId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateDetection(context *gin.Context) {

	detectionInput := types.Detection{}
	err := context.ShouldBindJSON(&detectionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	detectionId, err := controller.DBManager.CreateDetection(detectionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, detectionId)
}

func (controller PublicController) UpdateDetection(context *gin.Context) {

	detectionInput := types.Detection{}
	err := context.ShouldBindJSON(&detectionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	err = controller.DBManager.UpdateDetection(detectionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
