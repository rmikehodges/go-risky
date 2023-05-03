package mitigation

import (
	"go-risky/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type MitigationInput struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	ActionID    uuid.UUID     `json:"actionId"`
	Implemented bool          `json:"implemented"`
	CreatedAt   time.Time     `json:"createdAt"`
}

type MitigationOutput struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	ActionID    uuid.UUID     `json:"actionId"`
	Implemented bool          `json:"implemented"`
	CreatedAt   time.Time     `json:"createdAt"`
}

func inputToModel(mitigationInput MitigationInput) (mitigationModel database.MitigationModel, err error) {
	mitigationModel.ID = mitigationInput.ID
	mitigationModel.Name = mitigationInput.Name
	mitigationModel.Description = mitigationInput.Description
	mitigationModel.BusinessID = mitigationInput.BusinessID
	mitigationModel.ActionID = mitigationInput.ActionID
	mitigationModel.Implemented = mitigationInput.Implemented
	mitigationModel.CreatedAt = mitigationInput.CreatedAt
	return
}

func modelToOutput(mitigationModel database.MitigationModel) (mitigationOutput MitigationOutput, err error) {
	mitigationOutput.ID = mitigationModel.ID
	mitigationOutput.Name = mitigationModel.Name
	mitigationOutput.Description = mitigationModel.Description
	mitigationOutput.BusinessID = mitigationModel.BusinessID
	mitigationOutput.ActionID = mitigationModel.ActionID
	mitigationOutput.Implemented = mitigationModel.Implemented
	mitigationOutput.CreatedAt = mitigationModel.CreatedAt
	return
}

func modelsToOutput(mitigationModels []database.MitigationModel) (mitigationOutputs []MitigationOutput, err error) {
	for _, mitigationModel := range mitigationModels {
		mitigationOutput, err := modelToOutput(mitigationModel)
		if err != nil {
			return mitigationOutputs, err
		}
		mitigationOutputs = append(mitigationOutputs, mitigationOutput)
	}
	return
}

func getMitigations(context *gin.Context) {
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

	mitigationModels, err := database.GetMitigations(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	mitigationOutput, err := modelsToOutput(mitigationModels)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	context.JSON(http.StatusOK, mitigationOutput)
}

func getMitigation(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	mitigationId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	mitigationModel, err := database.GetMitigation(mitigationId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	mitigationOutput, err := modelToOutput(mitigationModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	context.JSON(http.StatusOK, mitigationOutput)
}

func deleteMitigation(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	mitigationId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = database.DeleteMitigation(mitigationId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createMitigation(context *gin.Context) {
	mitigationInput := MitigationInput{}
	err := context.ShouldBindJSON(&mitigationInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	mitigationModel, err := inputToModel(mitigationInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = database.CreateMitigation(mitigationModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateMitigation(context *gin.Context) {
	mitigationInput := MitigationInput{}
	err := context.ShouldBindJSON(&mitigationInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	mitigationModel, err := inputToModel(mitigationInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = database.UpdateMitigation(mitigationModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func MitigationRoutes(router *gin.Engine) {
	router.GET("/mitigations", getMitigations)
	router.GET("/mitigation/:id", getMitigation)
	router.DELETE("/mitigation/:id", deleteMitigation)
	router.PATCH("/mitigation/:id", updateMitigation)
	router.POST("/mitigations", createMitigation)
}
