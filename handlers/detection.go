package handlers

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
	Implemented bool          `json:"complexity"`
	CreatedAt   time.Time     `json:"createdAt"`
}

type DetectionOutput struct { //This is the output type that will be returned to the user
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	Implemented bool          `json:"complexity"`
	CreatedAt   time.Time     `json:"createdAt"`
}

type DetectionOutputs []DetectionOutput

func (detectionInput DetectionInput) inputToModel() (detectionModel database.DetectionModel, err error) {
	//This is where you do input validation sanitization
	detectionModel.ID = detectionInput.ID
	detectionModel.Name = detectionInput.Name
	detectionModel.Description = detectionInput.Description
	detectionModel.BusinessID = detectionInput.BusinessID
	detectionModel.Implemented = detectionInput.Implemented
	detectionModel.CreatedAt = detectionInput.CreatedAt
	return
}

func (detectionOutput *DetectionOutput) modelToOutput(detectionModel database.DetectionModel) (err error) {
	//This is where you do input validation sanitization
	detectionOutput.ID = detectionModel.ID
	detectionOutput.Name = detectionModel.Name
	detectionOutput.Description = detectionModel.Description
	detectionOutput.BusinessID = detectionModel.BusinessID
	detectionOutput.Implemented = detectionModel.Implemented
	detectionOutput.CreatedAt = detectionModel.CreatedAt
	return
}

func detectionModelsToOutput(detectionModels []database.DetectionModel) (detectionOutputs DetectionOutputs, err error) {
	for _, detectionModel := range detectionModels {
		detectionOutput := DetectionOutput{}
		err := detectionOutput.modelToOutput(detectionModel)
		if err != nil {
			return nil, err
		}
		detectionOutputs = append(detectionOutputs, detectionOutput)
	}
	return
}

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

	detectionModels, err := controller.DBManager.GetDetections(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, detectionModels)
		return
	}

	detectionOutput, err := detectionModelsToOutput(detectionModels)

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

	detectionModel, err := controller.DBManager.GetDetection(detectionId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, detectionModel)
		return
	}
	var detectionOutput DetectionOutput
	err = detectionOutput.modelToOutput(detectionModel)

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

	detectionInput := DetectionInput{}
	err := context.ShouldBindJSON(&detectionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	detectionModel, err := detectionInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	detectionId, err := controller.DBManager.CreateDetection(detectionModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, detectionId)
}

func (controller PublicController) UpdateDetection(context *gin.Context) {

	detectionInput := DetectionInput{}
	err := context.ShouldBindJSON(&detectionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	detectionModel, err := detectionInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = controller.DBManager.UpdateDetection(detectionModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
