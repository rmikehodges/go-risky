package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go-risky/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type ActionInput struct {
	ID              uuid.UUID     `json:"id"`
	Name            string        `json:"name" binding:"required"`
	Description     zeronull.Text `json:"description"`
	CapabilityID    *uuid.UUID    `json:"capabilityId"`
	VulnerabilityID *uuid.UUID    `json:"vulnerabilityId"`
	BusinessID      uuid.UUID     `json:"businessId" binding:"required"`
	Complexity      zeronull.Text `json:"complexity"`
	AssetID         *uuid.UUID    `json:"assetId"`
	CreatedAt       time.Time     `json:"createdAt"`
}

type ActionOutput struct {
	ID              uuid.UUID     `json:"id"`
	Name            string        `json:"name"`
	Description     zeronull.Text `json:"description"`
	CapabilityID    *uuid.UUID    `json:"capabilityId" db:"capability_id"`
	VulnerabilityID *uuid.UUID    `json:"vulnerabilityId" db:"vulnerability_id"`
	BusinessID      uuid.UUID     `json:"businessId" db:"business_id"`
	Complexity      zeronull.Text `json:"complexity"`
	AssetID         *uuid.UUID    `json:"assetId" db:"asset_id"`
	CreatedAt       time.Time     `json:"createdAt" db:"created_at"`
}

type ActionOutputs []ActionOutput

func (actionInput ActionInput) inputToModel() (actionModel database.ActionModel, err error) {
	//This is where you do input validation sanitization
	actionModel.ID = actionInput.ID
	actionModel.Name = actionInput.Name
	actionModel.Description = actionInput.Description
	actionModel.BusinessID = actionInput.BusinessID
	actionModel.Complexity = actionInput.Complexity
	actionModel.CreatedAt = actionInput.CreatedAt

	actionModel.CapabilityID = actionInput.CapabilityID
	actionModel.VulnerabilityID = actionInput.VulnerabilityID
	actionModel.AssetID = actionInput.AssetID

	return

}

func (actionOutput *ActionOutput) modelToOutput(actionModel database.ActionModel) (err error) {
	//This is where you do input validation sanitization
	actionOutput.ID = actionModel.ID
	actionOutput.Name = actionModel.Name
	actionOutput.Description = actionModel.Description
	actionOutput.CapabilityID = actionModel.CapabilityID
	actionOutput.VulnerabilityID = actionModel.VulnerabilityID
	actionOutput.BusinessID = actionModel.BusinessID
	actionOutput.Complexity = actionModel.Complexity
	actionOutput.AssetID = actionModel.AssetID
	actionOutput.CreatedAt = actionModel.CreatedAt

	return
}

func actionModelsToOutput(actionModels []database.ActionModel) (actionOutputs ActionOutputs, err error) {
	//This is where you do input validation sanitization
	for _, model := range actionModels {
		actionOutput := ActionOutput{}
		err = actionOutput.modelToOutput(model)
		if err != nil {
			return
		}
		fmt.Println(actionOutput)
		actionOutputs = append(actionOutputs, actionOutput)
	}

	return
}

func (controller PublicController) GetActions(context *gin.Context) {
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

	actionmodel, err := controller.DBManager.GetActions(businessId.String())

	if err != nil {
		log.Printf("Get Actions Error %s", err)
		context.JSON(http.StatusNotFound, actionmodel)
		return
	}

	actionOutputs, err := actionModelsToOutput(actionmodel)

	if err != nil {
		context.JSON(http.StatusNotFound, actionOutputs)
		return
	}

	context.JSON(http.StatusOK, actionOutputs)
}

func (controller PublicController) GetAction(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	actionId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	actionOutput, err := controller.DBManager.GetAction(actionId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, actionOutput)
		return
	}

	context.JSON(http.StatusOK, actionOutput)
}

func (controller PublicController) DeleteAction(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	actionId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = controller.DBManager.DeleteAction(actionId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateAction(context *gin.Context) {

	actionInput := ActionInput{}
	err := context.ShouldBindJSON(&actionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}
	actionModel, err := actionInput.inputToModel()
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	_, err = controller.DBManager.CreateAction(actionModel)
	fmt.Println("returned from create action")
	if err != nil {
		log.Println(err)
		fmt.Println("Error creating action")
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func (controller PublicController) UpdateAction(context *gin.Context) {

	actionInput := ActionInput{}
	err := context.ShouldBindJSON(&actionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	actionModel, err := actionInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	err = controller.DBManager.UpdateAction(actionModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
