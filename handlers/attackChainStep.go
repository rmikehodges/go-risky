package handlers

import (
	"go-risky/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Create types and functions for AttackChainStep that match the handler attackChain

func (controller PublicController) GetAttackChainSteps(context *gin.Context) {

	businessID := context.Query("businessId")
	attackChainId := context.Query("attackChainId")
	attackChainStepOutputs, err := controller.DBManager.GetAttackChainSteps(businessID, attackChainId)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, attackChainStepOutputs)
}

func (controller PublicController) GetAttackChainStep(context *gin.Context) {

	id := context.Param("id")
	attackChainStepOutput, err := controller.DBManager.GetAttackChainStep(id)
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

	var attackChainStepInput types.AttackChainStep
	err := context.ShouldBindJSON(&attackChainStepInput)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = controller.DBManager.UpdateAttackChainStep(attackChainStepInput)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, attackChainStepInput)
}

func (controller PublicController) CreateAttackChainStep(context *gin.Context) {

	var attackChainStepInput types.AttackChainStep
	err := context.ShouldBindJSON(&attackChainStepInput)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	attackChainStepOutput, err := controller.DBManager.CreateAttackChainStep(attackChainStepInput)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, attackChainStepOutput)
}

//Create the handlers for the AttackChainStep that matches the format of AttackChain handlers
