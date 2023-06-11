package handlers

import (
	"go-risky/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ImpactInput struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Description      *string   `json:"description"`
	BusinessID       uuid.UUID `json:"businessId"`
	ThreatID         uuid.UUID `json:"threatId"`
	ExploitationCost *float32  `json:"exploitationCost"`
	MitigationCost   *float32  `json:"mitigationCost"`
	CreatedAt        time.Time `json:"createdAt"`
}

type ImpactOutput struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Description      *string   `json:"description"`
	BusinessID       uuid.UUID `json:"businessId"`
	ThreatID         uuid.UUID `json:"threatId"`
	ExploitationCost *float32  `json:"exploitationCost"`
	MitigationCost   *float32  `json:"mitigationCost"`
	CreatedAt        time.Time `json:"createdAt"`
}

type ImpactOutputs []ImpactOutput

func (impactInput ImpactInput) inputToModel() (impactModel database.ImpactModel, err error) {
	impactModel.ID = impactInput.ID
	impactModel.Name = impactInput.Name
	impactModel.Description = impactInput.Description
	impactModel.BusinessID = impactInput.BusinessID
	impactModel.ThreatID = impactInput.ThreatID
	impactModel.ExploitationCost = impactInput.ExploitationCost
	impactModel.MitigationCost = impactInput.MitigationCost
	impactModel.CreatedAt = impactInput.CreatedAt
	return
}

func (impactOutput *ImpactOutput) modelToOutput(impactModel database.ImpactModel) (err error) {
	impactOutput.ID = impactModel.ID
	impactOutput.Name = impactModel.Name
	impactOutput.Description = impactModel.Description
	impactOutput.BusinessID = impactModel.BusinessID
	impactOutput.ThreatID = impactModel.ThreatID
	impactOutput.ExploitationCost = impactModel.ExploitationCost
	impactOutput.MitigationCost = impactModel.MitigationCost
	impactOutput.CreatedAt = impactModel.CreatedAt
	return
}

func impactModelsToOutput(impactModels []database.ImpactModel) (impactOutputs ImpactOutputs, err error) {
	for _, impactModel := range impactModels {
		impactOutput := ImpactOutput{}
		err := impactOutput.modelToOutput(impactModel)
		if err != nil {
			return nil, err
		}
		impactOutputs = append(impactOutputs, impactOutput)
	}
	return
}

func (controller PublicController) GetImpacts(context *gin.Context) {

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

	impactModel, err := controller.DBManager.GetImpacts(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, impactModel)
		return
	}

	impactOutputs, err := impactModelsToOutput(impactModel)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, impactOutputs)
		return
	}

	context.JSON(http.StatusOK, impactOutputs)
}

func (controller PublicController) GetImpact(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	impactId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	impactModel, err := controller.DBManager.GetImpact(impactId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	var impactOutput ImpactOutput
	err = impactOutput.modelToOutput(impactModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	context.JSON(http.StatusOK, impactOutput)
}

func (controller PublicController) DeleteImpact(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	impactId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = controller.DBManager.DeleteImpact(impactId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateImpact(context *gin.Context) {

	impactInput := ImpactInput{}
	err := context.ShouldBindJSON(&impactInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	impactModel, err := impactInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	impactId, err := controller.DBManager.CreateImpact(impactModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, impactId)
}

func (controller PublicController) UpdateImpact(context *gin.Context) {

	impactInput := ImpactInput{}
	err := context.ShouldBindJSON(&impactInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	impactModel, err := impactInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = controller.DBManager.UpdateImpact(impactModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
