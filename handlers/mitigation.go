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

type MitigationInput struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	ActionID    uuid.UUID     `json:"actionId"`
	Implemented bool          `json:"implemented"`
	CreatedAt   time.Time     `json:"createdAt"`
}

type MitigationOutput struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	Implemented bool          `json:"implemented"`
	CreatedAt   time.Time     `json:"createdAt"`
}
type MitigationOutputs []MitigationOutput

func (mitigationInput MitigationInput) inputToModel() (mitigationModel database.MitigationModel, err error) {
	mitigationModel.ID = mitigationInput.ID
	mitigationModel.Name = mitigationInput.Name
	mitigationModel.Description = mitigationInput.Description
	mitigationModel.BusinessID = mitigationInput.BusinessID
	mitigationModel.Implemented = mitigationInput.Implemented
	mitigationModel.CreatedAt = mitigationInput.CreatedAt
	return
}

func (mitigationOutput *MitigationOutput) modelToOutput(mitigationModel database.MitigationModel) (err error) {
	mitigationOutput.ID = mitigationModel.ID
	mitigationOutput.Name = mitigationModel.Name
	mitigationOutput.Description = mitigationModel.Description
	mitigationOutput.BusinessID = mitigationModel.BusinessID
	mitigationOutput.Implemented = mitigationModel.Implemented
	mitigationOutput.CreatedAt = mitigationModel.CreatedAt
	return
}

func mitigationModelsToOutput(mitigationModels []database.MitigationModel) (mitigationOutputs MitigationOutputs, err error) {
	for _, mitigationModel := range mitigationModels {
		mitigationOutput := MitigationOutput{}
		err := mitigationOutput.modelToOutput(mitigationModel)
		if err != nil {
			return nil, err
		}
		mitigationOutputs = append(mitigationOutputs, mitigationOutput)
	}
	return
}

func (controller PublicController) GetMitigations(context *gin.Context) {

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

	mitigationModels, err := controller.DBManager.GetMitigations(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	mitigationOutputs, err := mitigationModelsToOutput(mitigationModels)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	context.JSON(http.StatusOK, mitigationOutputs)
}

func (controller PublicController) GetMitigation(context *gin.Context) {

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

	mitigationModel, err := controller.DBManager.GetMitigation(mitigationId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	var mitigationOutput MitigationOutput
	err = mitigationOutput.modelToOutput(mitigationModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	context.JSON(http.StatusOK, mitigationOutput)
}

func (controller PublicController) DeleteMitigation(context *gin.Context) {

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

	err = controller.DBManager.DeleteMitigation(mitigationId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateMitigation(context *gin.Context) {

	mitigationInput := MitigationInput{}
	err := context.ShouldBindJSON(&mitigationInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	mitigationModel, err := mitigationInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	mitigationId, err := controller.DBManager.CreateMitigation(mitigationModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, mitigationId)
}

func (controller PublicController) UpdateMitigation(context *gin.Context) {

	mitigationInput := MitigationInput{}
	err := context.ShouldBindJSON(&mitigationInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	mitigationModel, err := mitigationInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = controller.DBManager.UpdateMitigation(mitigationModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
