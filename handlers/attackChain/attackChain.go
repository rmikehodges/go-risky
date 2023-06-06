package attackChain

import (
	"log"
	"net/http"
	"time"

	"go-risky/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

//Create types AttackChainInput and AttackChainOutput that match the database model AttackChainModel

type AttackChainInput struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	ThreatID    uuid.UUID     `json:"assetId"`
	CreatedAt   time.Time     `json:"createdAt"`
}

// This is the output type that will be returned to the user
type AttackChainOutput struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	ThreatID    uuid.UUID     `json:"assetId"`
	CreatedAt   time.Time     `json:"createdAt"`
}

//Create functions modelToOutput, inputToModel, and modelsToOutput that convert between the database model and the input/output types

func inputToModel(attackChainInput AttackChainInput) (attackChainModel database.AttackChainModel, err error) {
	//This is where you do input validation sanitization
	attackChainModel.ID = attackChainInput.ID
	attackChainModel.Name = attackChainInput.Name
	attackChainModel.Description = attackChainInput.Description
	attackChainModel.BusinessID = attackChainInput.BusinessID
	attackChainModel.ThreatID = attackChainInput.ThreatID
	attackChainModel.CreatedAt = attackChainInput.CreatedAt

	return

}

func modelToOutput(attackChainModel database.AttackChainModel) (attackChainOutput AttackChainOutput, err error) {
	//This is where you do input validation sanitization
	attackChainOutput.ID = attackChainModel.ID
	attackChainOutput.Name = attackChainModel.Name
	attackChainOutput.Description = attackChainModel.Description
	attackChainOutput.BusinessID = attackChainModel.BusinessID
	attackChainOutput.ThreatID = attackChainModel.ThreatID
	attackChainOutput.CreatedAt = attackChainModel.CreatedAt
	return
}

func modelsToOutput(attackChainModels []database.AttackChainModel) (attackChainOutputs []AttackChainOutput, err error) {
	for _, attackChainModel := range attackChainModels {
		attackChainOutput, err := modelToOutput(attackChainModel)
		if err != nil {
			return attackChainOutputs, err
		}
		attackChainOutputs = append(attackChainOutputs, attackChainOutput)
	}
	return
}

func getAttackChains(context *gin.Context) {
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

	attackChainModel, err := db.GetAttackChains(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	attackChainOutput, err := modelsToOutput(attackChainModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, attackChainOutput)
}

func getAttackChain(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	attackChainId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	attackChainModel, err := db.GetAttackChain(attackChainId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	attackChainOutput, err := modelToOutput(attackChainModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, attackChainOutput)
		return
	}

	context.JSON(http.StatusOK, attackChainOutput)
}

func deleteAttackChain(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	attackChainId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = db.DeleteAttackChain(attackChainId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createAttackChain(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	attackChainInput := AttackChainInput{}
	err := context.ShouldBindJSON(&attackChainInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	attackChainModel, err := inputToModel(attackChainInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	attackChainId, err := db.CreateAttackChain(attackChainModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, attackChainId)
}

func updateAttackChain(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	attackChainInput := AttackChainInput{}
	err := context.ShouldBindJSON(&attackChainInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	attackChainModel, err := inputToModel(attackChainInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = db.UpdateAttackChain(attackChainModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func AttackChainRoutes(router *gin.Engine) {
	router.GET("/attackChains", getAttackChains)
	router.GET("/attackChain", getAttackChain)
	router.DELETE("/attackChain", deleteAttackChain)
	router.PATCH("/attackChain", updateAttackChain)
	router.POST("/attackChain", createAttackChain)
}
