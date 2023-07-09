package handlers

import (
	"go-risky/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//Create types and functions for AttackChainStep that match the handler attackChain

type AttackChainStepInput struct {
	ID            uuid.UUID  `json:"id"`
	BusinessID    uuid.UUID  `json:"businessId"`
	ActionID      uuid.UUID  `json:"actionId"`
	AssetID       *uuid.UUID `json:"assetId"`
	AttackChainID uuid.UUID  `json:"attackChainId"`
	Postion       int        `json:"position"`
	CreatedAt     time.Time  `json:"createdAt"`
}

type AttackChainStepOutput struct {
	ID            uuid.UUID  `json:"id"`
	BusinessID    uuid.UUID  `json:"businessId"`
	ActionID      uuid.UUID  `json:"actionId"`
	AssetID       *uuid.UUID `json:"assetId"`
	AttackChainID uuid.UUID  `json:"attackChainId"`
	Position      int        `json:"position"`
	CreatedAt     time.Time  `json:"createdAt"`
}

type AttackChainStepOutputs []AttackChainStepOutput

func (attackChainStepInput AttackChainStepInput) inputToModel() (attackChainStepModel database.AttackChainStepModel, err error) {
	attackChainStepModel.BusinessID = attackChainStepInput.BusinessID
	attackChainStepModel.ActionID = attackChainStepInput.ActionID
	attackChainStepModel.AssetID = attackChainStepInput.AssetID
	attackChainStepModel.AttackChainID = attackChainStepInput.AttackChainID
	attackChainStepModel.Position = attackChainStepInput.Postion
	attackChainStepModel.CreatedAt = attackChainStepInput.CreatedAt

	return

}

func (attackChainStepOutput *AttackChainStepOutput) modelToOutput(attackChainStepModel database.AttackChainStepModel) (err error) {
	//This is where you do input validation sanitization
	attackChainStepOutput.BusinessID = attackChainStepModel.BusinessID
	attackChainStepOutput.ActionID = attackChainStepModel.ActionID
	attackChainStepOutput.AssetID = attackChainStepModel.AssetID
	attackChainStepOutput.AttackChainID = attackChainStepModel.AttackChainID
	attackChainStepOutput.Position = attackChainStepModel.Position
	attackChainStepOutput.CreatedAt = attackChainStepModel.CreatedAt
	return
}

func attackChainStepModelsToOutput(attackChainStepModels []database.AttackChainStepModel) (attackChainStepOutputs AttackChainStepOutputs, err error) {
	//This is where you do input validation sanitization
	for _, model := range attackChainStepModels {
		attackChainStepOutput := AttackChainStepOutput{}
		err := attackChainStepOutput.modelToOutput(model)
		if err != nil {
			return nil, err
		}
		attackChainStepOutputs = append(attackChainStepOutputs, attackChainStepOutput)
	}
	return
}

func (controller PublicController) GetAttackChainSteps(context *gin.Context) {

	businessID := context.Query("businessId")
	attackChainId := context.Query("attackChainId")
	attackChainStepModels, err := controller.DBManager.GetAttackChainSteps(businessID, attackChainId)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	attackChainStepOutputs, err := attackChainStepModelsToOutput(attackChainStepModels)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, attackChainStepOutputs)
}

func (controller PublicController) GetAttackChainStep(context *gin.Context) {

	id := context.Param("id")
	attackChainStepModel, err := controller.DBManager.GetAttackChainStep(id)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var attackChainStepOutput AttackChainStepOutput
	err = attackChainStepOutput.modelToOutput(attackChainStepModel)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, attackChainStepOutput)
}

func (controller PublicController) DeleteAttackChainStep(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err := controller.DBManager.DeleteAttackChainStep(id)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{})
}

func (controller PublicController) UpdateAttackChainStep(context *gin.Context) {

	var attackChainStepInput AttackChainStepInput
	err := context.ShouldBindJSON(&attackChainStepInput)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepModel, err := attackChainStepInput.inputToModel()
	if err != nil {

		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = controller.DBManager.UpdateAttackChainStep(attackChainStepModel)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var attackChainStepOutput AttackChainStepOutput
	err = attackChainStepOutput.modelToOutput(attackChainStepModel)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, attackChainStepOutput)
}

func (controller PublicController) CreateAttackChainStep(context *gin.Context) {

	var attackChainStepInput AttackChainStepInput
	err := context.ShouldBindJSON(&attackChainStepInput)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepModel, err := attackChainStepInput.inputToModel()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepOutput, err := controller.DBManager.CreateAttackChainStep(attackChainStepModel)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, attackChainStepOutput)
}

//Create the handlers for the AttackChainStep that matches the format of AttackChain handlers
