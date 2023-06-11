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

//Create types CapabilityInput and CapabilityOutput that match the database model CapabilityModel

type CapabilityInput struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	CreatedAt   time.Time     `json:"createdAt"`
}

type CapabilityOutput struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	CreatedAt   time.Time     `json:"createdAt"`
}

type CapabilityOutputs []CapabilityOutput

func (capabilityInput CapabilityInput) inputToModel() (capabilityModel database.CapabilityModel, err error) {
	//This is where you do input validation sanitization
	capabilityModel.ID = capabilityInput.ID
	capabilityModel.Name = capabilityInput.Name
	capabilityModel.Description = capabilityInput.Description
	capabilityModel.BusinessID = capabilityInput.BusinessID
	capabilityModel.CreatedAt = capabilityInput.CreatedAt
	return
}

func (capabilityOutput *CapabilityOutput) modelToOutput(capabilityModel database.CapabilityModel) (err error) {
	//This is where you do input validation sanitization
	capabilityOutput.ID = capabilityModel.ID
	capabilityOutput.Name = capabilityModel.Name
	capabilityOutput.Description = capabilityModel.Description
	capabilityOutput.BusinessID = capabilityModel.BusinessID
	capabilityOutput.CreatedAt = capabilityModel.CreatedAt
	return
}

func capabilityModelsToOutput(capabilityModels []database.CapabilityModel) (capabilityOutputs CapabilityOutputs, err error) {
	for _, capabilityModel := range capabilityModels {
		capabilityOutput := CapabilityOutput{}
		err = capabilityOutput.modelToOutput(capabilityModel)
		if err != nil {
			return nil, err
		}
		capabilityOutputs = append(capabilityOutputs, capabilityOutput)
	}
	return
}

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

	capabilityModel, err := controller.DBManager.GetCapabilities(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, capabilityModel)
		return
	}
	capabilityOutputs, err := capabilityModelsToOutput(capabilityModel)

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

	capabilityInput := CapabilityInput{}
	err := context.ShouldBindJSON(&capabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	capabilityModel, err := capabilityInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	capabilityId, err := controller.DBManager.CreateCapability(capabilityModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, capabilityId)
}

func (controller PublicController) UpdateCapability(context *gin.Context) {

	capabilityInput := CapabilityInput{}
	err := context.ShouldBindJSON(&capabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	capabilityModel, err := capabilityInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}
	err = controller.DBManager.UpdateCapability(capabilityModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
