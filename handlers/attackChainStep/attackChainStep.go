package attackChainStep

import (
	"go-risky/database"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//Create types and functions for AttackChainStep that match the handler attackChain

type AttackChainStepInput struct {
	BusinessID    uuid.UUID `json:"businessId"`
	ActionID      uuid.UUID `json:"actionId"`
	AssetID       uuid.UUID `json:"assetId"`
	AttackChainID uuid.UUID `json:"attackChainId"`
	Postion       int       `json:"position"`
	CreatedAt     time.Time `json:"createdAt"`
}

type AttackChainStepOutput struct {
	BusinessID    uuid.UUID `json:"businessId"`
	ActionID      uuid.UUID `json:"actionId"`
	AssetID       uuid.UUID `json:"assetId"`
	AttackChainID uuid.UUID `json:"attackChainId"`
	Position      int       `json:"position"`
	CreatedAt     time.Time `json:"createdAt"`
}

func inputToModel(attackChainStepInput AttackChainStepInput) (attackChainStepModel database.AttackChainStepModel, err error) {
	attackChainStepModel.BusinessID = attackChainStepInput.BusinessID
	attackChainStepModel.ActionID = attackChainStepInput.ActionID
	attackChainStepModel.AttackChainID = attackChainStepInput.AttackChainID
	attackChainStepModel.Position = attackChainStepInput.Postion
	attackChainStepModel.CreatedAt = attackChainStepInput.CreatedAt

	return

}

func modelToOutput(attackChainStepModel database.AttackChainStepModel) (attackChainStepOutput AttackChainStepOutput, err error) {
	//This is where you do input validation sanitization
	attackChainStepOutput.BusinessID = attackChainStepModel.BusinessID
	attackChainStepOutput.ActionID = attackChainStepModel.ActionID
	attackChainStepOutput.AttackChainID = attackChainStepModel.AttackChainID
	attackChainStepOutput.Position = attackChainStepModel.Position
	attackChainStepOutput.CreatedAt = attackChainStepModel.CreatedAt
	return
}

func modelsToOutput(attackChainStepModels []database.AttackChainStepModel) (attackChainStepOutput []AttackChainStepOutput, err error) {
	//This is where you do input validation sanitization
	for _, model := range attackChainStepModels {
		output, err := modelToOutput(model)
		if err != nil {
			return nil, err
		}
		attackChainStepOutput = append(attackChainStepOutput, output)
	}
	return
}

func getAttackChainSteps(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)
	businessID := context.Query("businessId")
	attackChainId := context.Query("attackChainId")
	attackChainStepModels, err := db.GetAttackChainSteps(businessID, attackChainId)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepOutput, err := modelsToOutput(attackChainStepModels)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, attackChainStepOutput)
}

func getAttackChainStep(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)
	id := context.Param("id")
	attackChainStepModel, err := db.GetAttackChainStep(id)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepOutput, err := modelToOutput(attackChainStepModel)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, attackChainStepOutput)
}

func deleteAttackChainStep(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)
	id := context.Param("id")
	err := db.DeleteAttackChainStep(id)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{})
}

func updateAttackChainStep(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)
	var attackChainStepInput AttackChainStepInput
	err := context.ShouldBindJSON(&attackChainStepInput)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepModel, err := inputToModel(attackChainStepInput)
	if err != nil {

		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = db.UpdateAttackChainStep(attackChainStepModel)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepOutput, err := modelToOutput(attackChainStepModel)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, attackChainStepOutput)
}

func createAttackChainStep(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)
	var attackChainStepInput AttackChainStepInput
	err := context.ShouldBindJSON(&attackChainStepInput)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepModel, err := inputToModel(attackChainStepInput)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepModel, err = db.CreateAttackChainStep(attackChainStepModel)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepOutput, err := modelToOutput(attackChainStepModel)
	if err != nil {

		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, attackChainStepOutput)
}

//Create the handlers for the AttackChainStep that matches the format of AttackChain handlers

func AttackChainStepRoutes(router *gin.Engine) {
	router.GET("/attackChainSteps", getAttackChainSteps)
	router.GET("/attackChainStep/:id", getAttackChainStep)
	router.DELETE("/attackChainStep/:id", deleteAttackChainStep)
	router.PATCH("/attackChainStep/:id", updateAttackChainStep)
	router.POST("/attackChainSteps", createAttackChainStep)
}
