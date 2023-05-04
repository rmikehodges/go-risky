package AttackChainStep

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
	AttackChainID uuid.UUID `json:"attackChainId"`
	Postion       int       `json:"position"`
	CreatedAt     time.Time `json:"createdAt"`
}

type AttackChainStepOutput struct {
	BusinessID    uuid.UUID `json:"businessId"`
	ActionID      uuid.UUID `json:"actionId"`
	AttackChainID uuid.UUID `json:"attackChainId"`
	Position      int       `json:"position"`
	CreatedAt     time.Time `json:"createdAt"`
}

func inputToModel(attackChainStepInput AttackChainStepInput) (attackChainStepModel database.attackChainStepModel, err error) {
	attackChainStepModel.BusinessID = AttackChainStepInput.BusinessID
	attackChainStepModel.ActionID = AttackChainStepInput.ActionID
	attackChainStepModel.AttackChainID = AttackChainStepInput.AttackChainID
	attackChainStepModel.Position = AttackChainStepInput.Postion
	attackChainStepModel.CreatedAt = AttackChainStepInput.CreatedAt

	return

}

func modelToOutput(attackChainStepModel database.attackChainStepModel) (attackChainStepOutput AttackChainStepOutput, err error) {
	//This is where you do input validation sanitization
	attackChainStepOutput.BusinessID = AttackChainStepModel.BusinessID
	attackChainStepOutput.ActionID = AttackChainStepModel.ActionID
	attackChainStepOutput.AttackChainID = AttackChainStepModel.AttackChainID
	attackChainStepOutput.Position = AttackChainStepModel.Postion
	attackChainStepOutput.CreatedAt = AttackChainStepModel.CreatedAt
	return
}

func modelsToOutput(attackChainStepModels []database.attackChainStepModel) (attackChainStepOutput []attackChainStepOutput, err error) {
	//This is where you do input validation sanitization
	for _, model := range AttackChainStepModels {
		output, err := modelToOutput(model)
		if err != nil {
			return nil, err
		}
		attackChainStepOutput = append(attackChainStepOutput, output)
	}
	return
}

func getAttackChainSteps(c *gin.Context) {
	businessID := c.Query("businessId")
	attackChainStepModels, err := database.GetAttackChainSteps(businessID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepOutput, err := modelsToOutput(attackChainStepModels)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, AttackChainStepOutput)
}

func getAttackChainStep(c *gin.Context) {
	id := c.Param("id")
	attackChainStepModel, err := database.GetAttackChainStep(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepOutput, err := modelToOutput(attackChainStepModel)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, AttackChainStepOutput)
}

func deleteAttackChainStep(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteAttackChainStep(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}

func updateAttackChainStep(c *gin.Context) {
	var AttackChainStepInput AttackChainStepInput
	err := c.ShouldBindJSON(&attackChainStepInput)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepModel, err := inputToModel(attackChainStepInput)
	if err != nil {

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = database.UpdateAttackChainStep(attackChainStepModel)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepOutput, err := modelToOutput(attackChainStepModel)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, AttackChainStepOutput)
}

func createAttackChainStep(c *gin.Context) {
	var AttackChainStepInput AttackChainStepInput
	err := c.ShouldBindJSON(&attackChainStepInput)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepModel, err := inputToModel(attackChainStepInput)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = database.CreateAttackChainStep(attackChainStepModel)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepOutput, err := modelToOutput(attackChainStepModel)
	if err != nil {

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, AttackChainStepOutput)
}

//Create the handlers for the AttackChainStep that matches the format of AttackChain handlers

func AttackChainStepRoutes(router *gin.Engine) {
	router.GET("/attackChainSteps", getAttackChainSteps)
	router.GET("/attackChainStep/:id", getAttackChainStep)
	router.DELETE("/attackChainStep/:id", deleteAttackChainStep)
	router.PATCH("/attackChainStep/:id", updateAttackChainStep)
	router.POST("/attackChainSteps", createAttackChainStep)
}
