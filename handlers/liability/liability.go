package liability

import (
	"go-risky/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type LiabilityInput struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  zeronull.Text `json:"description"`
	Quantity     float32       `json:"quantity"`
	Cost         float32       `json:"cost"`
	BusinessID   uuid.UUID     `json:"businessId"`
	DetectionID  *uuid.UUID    `json:"detectionId"`
	MitigationID *uuid.UUID    `json:"mitigationId"`
	ResourceID   *uuid.UUID    `json:"resourceId"`
	ThreatID     *uuid.UUID    `json:"threatId"`
	ImpactID     *uuid.UUID    `json:"impactId"`
	CreatedAt    time.Time     `json:"createdAt"`
}

type LiabilityOutput struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  zeronull.Text `json:"description"`
	Quantity     float32       `json:"quantity"`
	Cost         float32       `json:"cost"`
	BusinessID   uuid.UUID     `json:"businessId"`
	DetectionID  *uuid.UUID    `json:"detectionId"`
	MitigationID *uuid.UUID    `json:"mitigationId"`
	ResourceID   *uuid.UUID    `json:"resourceId"`
	ThreatID     *uuid.UUID    `json:"threatId"`
	ImpactID     *uuid.UUID    `json:"impactId"`
	CreatedAt    time.Time     `json:"createdAt"`
}

func inputToModel(liabilityInput LiabilityInput) (liabilityModel database.LiabilityModel, err error) {
	liabilityModel.ID = liabilityInput.ID
	liabilityModel.Name = liabilityInput.Name
	liabilityModel.Description = liabilityInput.Description
	liabilityModel.Quantity = liabilityInput.Quantity
	liabilityModel.Cost = liabilityInput.Cost
	liabilityModel.BusinessID = liabilityInput.BusinessID
	liabilityModel.MitigationID = liabilityInput.MitigationID
	liabilityModel.ResourceID = liabilityInput.ResourceID
	liabilityModel.ThreatID = liabilityInput.ThreatID
	liabilityModel.ImpactID = liabilityInput.ImpactID
	liabilityModel.CreatedAt = liabilityInput.CreatedAt
	return
}

func modelToOutput(liabilityModel database.LiabilityModel) (liabilityOutput LiabilityOutput, err error) {
	liabilityOutput.ID = liabilityModel.ID
	liabilityOutput.Name = liabilityModel.Name
	liabilityOutput.Description = liabilityModel.Description
	liabilityOutput.Quantity = liabilityModel.Quantity
	liabilityOutput.Cost = liabilityModel.Cost
	liabilityOutput.BusinessID = liabilityModel.BusinessID
	liabilityOutput.MitigationID = liabilityModel.MitigationID
	liabilityOutput.ResourceID = liabilityModel.ResourceID
	liabilityOutput.ThreatID = liabilityModel.ThreatID
	liabilityOutput.ImpactID = liabilityModel.ImpactID
	liabilityOutput.CreatedAt = liabilityModel.CreatedAt
	return
}

func modelsToOutputs(liabilityModels []database.LiabilityModel) (liabilityOutputs []LiabilityOutput, err error) {
	for _, liabilityModel := range liabilityModels {
		liabilityOutput, err := modelToOutput(liabilityModel)
		if err != nil {
			return nil, err
		}
		liabilityOutputs = append(liabilityOutputs, liabilityOutput)
	}
	return
}

func getLiabilities(context *gin.Context) {
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

	liabilityModels, err := db.GetLiabilities(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	liabilityOutput, err := modelsToOutputs(liabilityModels)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	context.JSON(http.StatusOK, liabilityOutput)
}

func getLiability(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	liabilityId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	liabilityModel, err := db.GetLiability(liabilityId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	liabilityOutput, err := modelToOutput(liabilityModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}
	context.JSON(http.StatusOK, liabilityOutput)
}

func getLiabilityByImpactId(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	id, ok := context.GetQuery("impactId")
	if !ok {
		log.Println("Parameter impactId not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	liabilityId, err := uuid.Parse(id)
	if err != nil {
		log.Println("impactId is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	liabilityModel, err := db.GetLiabilityByImpactId(liabilityId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	liabilityOutput, err := modelsToOutputs(liabilityModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}
	context.JSON(http.StatusOK, liabilityOutput)
}

func deleteLiability(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	liabilityId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = db.DeleteLiability(liabilityId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createLiability(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	liabilityInput := LiabilityInput{}
	err := context.ShouldBindJSON(&liabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	liabilityModel, err := inputToModel(liabilityInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	liabilityId, err := db.CreateLiability(liabilityModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, liabilityId)
}

func updateLiability(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	liabilityInput := LiabilityInput{}
	err := context.ShouldBindJSON(&liabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	liabilityModel, err := inputToModel(liabilityInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = db.UpdateLiability(liabilityModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
func LiabilityRoutes(router *gin.Engine) {
	router.GET("/liabilities", getLiabilities)
	router.GET("/liability", getLiability)
	router.GET("/liabilityByImpactId", getLiabilityByImpactId)
	router.DELETE("/liability", deleteLiability)
	router.PATCH("/liability", updateLiability)
	router.POST("/liabilities", createLiability)
}
