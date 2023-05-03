package attackChainAction

import (
	"go-risky/database"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//Create types and functions for attackChainAction that match the handler attackChain

type AttackChainActionInput struct {
	BusinessID    uuid.UUID `json:"businessId"`
	ActionID      uuid.UUID `json:"actionId"`
	AttackChainID uuid.UUID `json:"attackChainId"`
	Postion       int       `json:"position"`
	CreatedAt     time.Time `json:"createdAt"`
}

type AttackChainActionOutput struct {
	BusinessID    uuid.UUID `json:"businessId"`
	ActionID      uuid.UUID `json:"actionId"`
	AttackChainID uuid.UUID `json:"attackChainId"`
	Position      int       `json:"position"`
	CreatedAt     time.Time `json:"createdAt"`
}

func inputToModel(attackChainActionInput AttackChainActionInput) (attackChainActionModel database.AttackChainActionModel, err error) {
	attackChainActionModel.BusinessID = attackChainActionInput.BusinessID
	attackChainActionModel.ActionID = attackChainActionInput.ActionID
	attackChainActionModel.AttackChainID = attackChainActionInput.AttackChainID
	attackChainActionModel.Position = attackChainActionInput.Postion
	attackChainActionModel.CreatedAt = attackChainActionInput.CreatedAt

	return

}

func modelToOutput(attackChainActionModel database.AttackChainActionModel) (attackChainActionOutput AttackChainActionOutput, err error) {
	//This is where you do input validation sanitization
	attackChainActionOutput.BusinessID = attackChainActionModel.BusinessID
	attackChainActionOutput.ActionID = attackChainActionModel.ActionID
	attackChainActionOutput.AttackChainID = attackChainActionModel.AttackChainID
	attackChainActionOutput.Position = attackChainActionModel.Postion
	attackChainActionOutput.CreatedAt = attackChainActionModel.CreatedAt
	return
}

func modelsToOutput(attackChainActionModels []database.AttackChainActionModel) (attackChainActionOutput []AttackChainActionOutput, err error) {
	//This is where you do input validation sanitization
	for _, model := range attackChainActionModels {
		output, err := modelToOutput(model)
		if err != nil {
			return nil, err
		}
		attackChainActionOutput = append(attackChainActionOutput, output)
	}
	return
}

func getAttackChainActions(c *gin.Context) {
	businessID := c.Query("businessId")
	attackChainActionModels, err := database.GetAttackChainActions(businessID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainActionOutput, err := modelsToOutput(attackChainActionModels)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, attackChainActionOutput)
}

func getAttackChainAction(c *gin.Context) {
	id := c.Param("id")
	attackChainActionModel, err := database.GetAttackChainAction(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainActionOutput, err := modelToOutput(attackChainActionModel)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, attackChainActionOutput)
}

func deleteAttackChainAction(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteAttackChainAction(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}

func updateAttackChainAction(c *gin.Context) {
	var attackChainActionInput AttackChainActionInput
	err := c.ShouldBindJSON(&attackChainActionInput)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainActionModel, err := inputToModel(attackChainActionInput)
	if err != nil {

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = database.UpdateAttackChainAction(attackChainActionModel)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainActionOutput, err := modelToOutput(attackChainActionModel)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, attackChainActionOutput)
}

func createAttackChainAction(c *gin.Context) {
	var attackChainActionInput AttackChainActionInput
	err := c.ShouldBindJSON(&attackChainActionInput)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainActionModel, err := inputToModel(attackChainActionInput)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = database.CreateAttackChainAction(attackChainActionModel)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainActionOutput, err := modelToOutput(attackChainActionModel)
	if err != nil {

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, attackChainActionOutput)
}

//Create the handlers for the attackChainAction that matches the format of AttackChain handlers

func AttackChainActionRoutes(router *gin.Engine) {
	router.GET("/attackChainActions", getAttackChainActions)
	router.GET("/attackChainAction/:id", getAttackChainAction)
	router.DELETE("/attackChainAction/:id", deleteAttackChainAction)
	router.PATCH("/attackChainAction/:id", updateAttackChainAction)
	router.POST("/attackChainActions", createAttackChainAction)
}
