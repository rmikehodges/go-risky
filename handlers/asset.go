package handlers

import (
	"log"
	"net/http"
	"time"

	"go-risky/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type AssetInput struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
}

type AssetOutput struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at"`
}

type AssetOutputs []AssetOutput

func (assetInput AssetInput) inputToModel() (assetModel database.AssetModel, err error) {
	//This is where you do input validation sanitization
	assetModel.ID = assetInput.ID
	assetModel.Name = assetInput.Name
	assetModel.Description = assetInput.Description
	assetModel.BusinessID = assetInput.BusinessID

	return

}

func (assetOutput *AssetOutput) modelToOutput(assetModel database.AssetModel) (err error) {
	//This is where you do input validation sanitization
	assetOutput.ID = assetModel.ID
	assetOutput.Name = assetModel.Name
	assetOutput.Description = assetModel.Description
	assetOutput.BusinessID = assetModel.BusinessID
	assetOutput.CreatedAt = assetModel.CreatedAt

	return
}

func assetModelsToOutput(assetModels []database.AssetModel) (assetOutputs AssetOutputs, err error) {
	//This is where you do input validation sanitization
	for _, model := range assetModels {
		assetOutput := AssetOutput{}
		err := assetOutput.modelToOutput(model)
		if err != nil {
			return nil, err
		}
		assetOutputs = append(assetOutputs, assetOutput)
	}

	return
}

func (controller PublicController) GetAssets(context *gin.Context) {

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

	assetModel, err := controller.DBManager.GetAssets(businessId.String())
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, assetModel)
		return
	}

	assetOutputs, err := assetModelsToOutput(assetModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, assetOutputs)
		return
	}

	context.JSON(http.StatusOK, assetOutputs)
}

func (controller PublicController) GetAsset(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	assetId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	assetModel, err := controller.DBManager.GetAsset(assetId.String())
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, assetModel)
		return
	}

	var assetOutput AssetOutput
	err = assetOutput.modelToOutput(assetModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, assetOutput)
		return
	}
	context.JSON(http.StatusOK, assetOutput)
}

func (controller PublicController) DeleteAsset(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	assetId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = controller.DBManager.DeleteAsset(assetId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateAsset(context *gin.Context) {

	assetInput := AssetInput{}
	err := context.ShouldBindJSON(&assetInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	assetModel, err := assetInput.inputToModel()
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	assetId, err := controller.DBManager.CreateAsset(assetModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, assetId)
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func (controller PublicController) UpdateAsset(context *gin.Context) {

	assetInput := AssetInput{}
	err := context.ShouldBindJSON(&assetInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	assetModel, err := assetInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	err = controller.DBManager.UpdateAsset(assetModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
