package detection

import (
	"go-risky/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

//Create types DetectionInput and DetectionOutput that match the database model DetectionModel

type DetectionInput struct { //This is the input type that will be received from the user in the POST request body (create) and PUT request body (update)
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	ActionID    uuid.UUID     `json:"actionId" db:"action_id"`
	Implemented bool          `json:"complexity"`
	CreatedAt   time.Time     `json:"createdAt"`
}

type DetectionOutput struct { //This is the output type that will be returned to the user
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	ActionID    uuid.UUID     `json:"actionId" db:"action_id"`
	Implemented bool          `json:"complexity"`
	CreatedAt   time.Time     `json:"createdAt"`
}

//Create functions modelToOutput, inputToModel, and modelsToOutput that convert between the database model and the input/output types

func inputToModel(detectionInput DetectionInput) (detectionModel database.DetectionModel, err error) {
	//This is where you do input validation sanitization
	detectionModel.ID = detectionInput.ID
	detectionModel.Name = detectionInput.Name
	detectionModel.Description = detectionInput.Description
	detectionModel.BusinessID = detectionInput.BusinessID
	detectionModel.ActionID = detectionInput.ActionID
	detectionModel.Implemented = detectionInput.Implemented
	detectionModel.CreatedAt = detectionInput.CreatedAt
	return
}

func modelToOutput(detectionModel database.DetectionModel) (detectionOutput DetectionOutput, err error) {
	//This is where you do input validation sanitization
	detectionOutput.ID = detectionModel.ID
	detectionOutput.Name = detectionModel.Name
	detectionOutput.Description = detectionModel.Description
	detectionOutput.BusinessID = detectionModel.BusinessID
	detectionOutput.ActionID = detectionModel.ActionID
	detectionOutput.Implemented = detectionModel.Implemented
	detectionOutput.CreatedAt = detectionModel.CreatedAt
	return
}

func modelsToOutput(detectionModels []database.DetectionModel) (detectionOutputs []DetectionOutput, err error) {
	for _, detectionModel := range detectionModels {
		detectionOutput, err := modelToOutput(detectionModel)
		if err != nil {
			return detectionOutputs, err
		}
		detectionOutputs = append(detectionOutputs, detectionOutput)
	}
	return
}

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

	detectionModels, err := database.GetDetections(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, detectionModels)
		return
	}

	detectionOutput, err := modelsToOutput(detectionModels)

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

	detectionModel, err := database.GetDetection(detectionId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, detectionModel)
		return
	}

	detectionOutput, err := modelToOutput(detectionModel)

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
	detectionInput := DetectionInput{}
	err := context.ShouldBindJSON(&detectionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	detectionModel, err := inputToModel(detectionInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = database.CreateDetection(detectionModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateDetection(context *gin.Context) {
	detectionInput := DetectionInput{}
	err := context.ShouldBindJSON(&detectionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	detectionModel, err := inputToModel(detectionInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = database.UpdateDetection(detectionModel)
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
