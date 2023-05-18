package asset

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

func inputToModel(assetInput AssetInput) (assetModel database.AssetModel, err error) {
	//This is where you do input validation sanitization
	assetModel.ID = assetInput.ID
	assetModel.Name = assetInput.Name
	assetModel.Description = assetInput.Description
	assetModel.BusinessID = assetInput.BusinessID

	return

}

func modelToOutput(assetModel database.AssetModel) (assetOutput AssetOutput, err error) {
	//This is where you do input validation sanitization
	assetOutput.ID = assetModel.ID
	assetOutput.Name = assetModel.Name
	assetOutput.Description = assetModel.Description
	assetOutput.BusinessID = assetModel.BusinessID
	assetOutput.CreatedAt = assetModel.CreatedAt

	return
}

func modelsToOutput(assetModels []database.AssetModel) (assetOutput []AssetOutput, err error) {
	//This is where you do input validation sanitization
	for _, model := range assetModels {
		output, err := modelToOutput(model)
		if err != nil {
			return []AssetOutput{}, err
		}
		assetOutput = append(assetOutput, output)
	}

	return
}

func getAssets(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

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

	assetModel, err := db.GetAssets(businessId.String())
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, assetModel)
		return
	}

	assetOutput, err := modelsToOutput(assetModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, assetOutput)
		return
	}

	context.JSON(http.StatusOK, assetOutput)
}

func getAsset(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

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

	assetModel, err := db.GetAsset(assetId.String())
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, assetModel)
		return
	}

	assetOutput, err := modelToOutput(assetModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, assetOutput)
		return
	}
	context.JSON(http.StatusOK, assetOutput)
}

func deleteAsset(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

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

	err = db.DeleteAsset(assetId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createAsset(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	assetInput := AssetInput{}
	err := context.ShouldBindJSON(&assetInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	assetModel, err := inputToModel(assetInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	err = db.CreateAsset(assetModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateAsset(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	assetInput := AssetInput{}
	err := context.ShouldBindJSON(&assetInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	assetModel, err := inputToModel(assetInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	err = db.UpdateAsset(assetModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
func AssetRoutes(router *gin.Engine) {
	router.GET("/assets", getAssets)
	router.GET("/asset", getAsset)
	router.DELETE("/asset", deleteAsset)
	router.PUT("/asset", updateAsset)
	router.POST("/assets", createAsset)
}
