package action

import (
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
	Name            string        `json:"name"`
	Description     zeronull.Text `json:"description"`
	CapabilityID    uuid.UUID     `json:"capabilityId" db:"capability_id"`
	VulnerabilityID uuid.UUID     `json:"vulnerabilityId" db:"vulnerability_id"`
	BusinessID      uuid.UUID     `json:"businessId" db:"business_id"`
	Complexity      zeronull.Text `json:"complexity"`
	AssetID         uuid.UUID     `json:"assetId" db:"asset_id"`
	CreatedAt       time.Time     `json:"createdAt" db:"created_at"`
}

type ActionOutput struct {
	ID              uuid.UUID     `json:"id"`
	Name            string        `json:"name"`
	Description     zeronull.Text `json:"description"`
	CapabilityID    uuid.UUID     `json:"capabilityId" db:"capability_id"`
	VulnerabilityID uuid.UUID     `json:"vulnerabilityId" db:"vulnerability_id"`
	BusinessID      uuid.UUID     `json:"businessId" db:"business_id"`
	Complexity      zeronull.Text `json:"complexity"`
	AssetID         uuid.UUID     `json:"assetId" db:"asset_id"`
	CreatedAt       time.Time     `json:"createdAt" db:"created_at"`
}

func inputToModel(actionInput ActionInput) (actionModel database.ActionModel, err error) {
	//This is where you do input validation sanitization
	actionModel.ID = actionInput.ID
	actionModel.Name = actionInput.Name
	actionModel.Description = actionInput.Description
	actionModel.CapabilityID = actionInput.CapabilityID
	actionModel.VulnerabilityID = actionInput.VulnerabilityID
	actionModel.BusinessID = actionInput.BusinessID
	actionModel.Complexity = actionInput.Complexity
	actionModel.AssetID = actionInput.AssetID
	actionModel.CreatedAt = actionInput.CreatedAt

	return

}


func modelToOutput(actionModel database.ActionModel) (actionOutput ActionOutput, err error) {
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

func modelsToOutput(actionModels []database.ActionModel) (actionOutput []ActionOutput, err error) {
	//This is where you do input validation sanitization
	for _, model := range actionModels {
		output, err := modelToOutput(model)
		if err != nil {
			return []ActionOutput{}, err
		}
		actionOutput = append(actionOutput, output)
	}

	return
}


func getActions(context *gin.Context) {
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

	actionmodel, err := database.GetActions(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, actionmodel)
		return
	}

	actionOutput, err := modelsToOutput(actionmodel)

	if err != nil {
		context.JSON(http.StatusNotFound, actionOutput)
		return
	}

	context.JSON(http.StatusOK, actionOutput)
}

func getAction(context *gin.Context) {
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

	actionOutput, err := database.GetAction(actionId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, actionOutput)
		return
	}

	context.JSON(http.StatusOK, actionOutput)
}

func deleteAction(context *gin.Context) {
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

	err = database.DeleteAction(actionId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createAction(context *gin.Context) {
	actionInput := ActionInput{}
	err := context.ShouldBindJSON(&actionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	actionModel, err := inputToModel(actionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	err = database.CreateAction(actionModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateAction(context *gin.Context) {
	actionInput := ActionInput{}
	err := context.ShouldBindJSON(&actionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	actionModel, err := inputToModel(actionInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	err = database.UpdateAction(actionModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func ActionRoutes(router *gin.Engine) {
	router.GET("/actions", getActions)
	router.GET("/action", getAction)
	router.DELETE("/action", deleteAction)
	router.PATCH("/action", updateAction)
	router.POST("/action", createAction)
}
