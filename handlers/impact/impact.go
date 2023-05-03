package impact

import (
	"go-risky/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type ImpactInput struct {
	ID               uuid.UUID       `json:"id"`
	Name             string          `json:"name"`
	Description      zeronull.Text   `json:"description"`
	BusinessID       uuid.UUID       `json:"businessId"`
	ThreatID         uuid.UUID       `json:"threatId"`
	ExploitationCost zeronull.Float8 `json:"exploitationCost"`
	MitigationCost   zeronull.Float8 `json:"mitigationCost"`
	CreatedAt        time.Time       `json:"createdAt"`
}

type ImpactOutput struct {
	ID               uuid.UUID       `json:"id"`
	Name             string          `json:"name"`
	Description      zeronull.Text   `json:"description"`
	BusinessID       uuid.UUID       `json:"businessId"`
	ThreatID         uuid.UUID       `json:"threatId"`
	ExploitationCost zeronull.Float8 `json:"exploitationCost"`
	MitigationCost   zeronull.Float8 `json:"mitigationCost"`
	CreatedAt        time.Time       `json:"createdAt"`
}

func inputToModel(impactInput ImpactInput) (impactModel database.ImpactModel, err error) {
	impactModel.ID = impactInput.ID
	impactModel.Name = impactInput.Name
	impactModel.Description = impactInput.Description
	impactModel.BusinessID = impactInput.BusinessID
	impactModel.ThreatID = impactInput.ThreatID
	impactModel.ExploitationCost = impactInput.ExploitationCost
	impactModel.MitigationCost = impactInput.MitigationCost
	impactModel.CreatedAt = impactInput.CreatedAt
	return
}

func modelToOutput(impactModel database.ImpactModel) (impactOutput ImpactOutput, err error) {
	impactOutput.ID = impactModel.ID
	impactOutput.Name = impactModel.Name
	impactOutput.Description = impactModel.Description
	impactOutput.BusinessID = impactModel.BusinessID
	impactOutput.ThreatID = impactModel.ThreatID
	impactOutput.ExploitationCost = impactModel.ExploitationCost
	impactOutput.MitigationCost = impactModel.MitigationCost
	impactOutput.CreatedAt = impactModel.CreatedAt
	return
}

func modelsToOutput(impactModels []database.ImpactModel) (impactOutputs []ImpactOutput, err error) {
	for _, impactModel := range impactModels {
		impactOutput, err := modelToOutput(impactModel)
		if err != nil {
			return nil, err
		}
		impactOutputs = append(impactOutputs, impactOutput)
	}
	return
}

func getImpacts(context *gin.Context) {
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

	impactModel, err := database.GetImpacts(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, impactModel)
		return
	}

	impactOutput, err := modelsToOutput(impactModel)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, impactOutput)
		return
	}

	context.JSON(http.StatusOK, impactOutput)
}

func getImpact(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	impactId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	impactModel, err := database.GetImpact(impactId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	impactOutput, err := modelToOutput(impactModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	context.JSON(http.StatusOK, impactOutput)
}

func deleteImpact(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	impactId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = database.DeleteImpact(impactId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createImpact(context *gin.Context) {
	impactInput := ImpactInput{}
	err := context.ShouldBindJSON(&impactInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	impactModel, err := inputToModel(impactInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = database.CreateImpact(impactModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateImpact(context *gin.Context) {
	impactInput := ImpactInput{}
	err := context.ShouldBindJSON(&impactInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	impactModel, err := inputToModel(impactInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = database.UpdateImpact(impactModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func ImpactRoutes(router *gin.Engine) {
	router.GET("/impacts", getImpacts)
	router.GET("/impact/:id", getImpact)
	router.DELETE("/impact/:id", deleteImpact)
	router.PATCH("/impact/:id", updateImpact)
	router.POST("/impacts", createImpact)
}
