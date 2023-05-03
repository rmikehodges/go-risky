package detection

import (
	"go-risky/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getDetections(context *gin.Context) {
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

	detectionOutput, err := database.GetDetections(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, detectionOutput)
		return
	}

	context.JSON(http.StatusOK, detectionOutput)
}

func getDetection(context *gin.Context) {
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

	detectionOutput, err := database.GetDetection(detectionId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, detectionOutput)
		return
	}

	context.JSON(http.StatusOK, detectionOutput)
}

func deleteDetection(context *gin.Context) {
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

	err = database.DeleteDetection(detectionId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createDetection(context *gin.Context) {
	detectionInput := database.Detection{}
	err := context.ShouldBindJSON(&detectionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.CreateDetection(detectionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateDetection(context *gin.Context) {
	detectionInput := database.Detection{}
	err := context.ShouldBindJSON(&detectionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.UpdateDetection(detectionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func DetectionRoutes(router *gin.Engine) {
	router.GET("/detections", getDetections)
	router.GET("/detection", getDetection)
	router.DELETE("/detection", deleteDetection)
	router.PATCH("/detection", updateDetection)
	router.POST("/detections", createDetection)
}
