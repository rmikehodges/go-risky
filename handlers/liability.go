package handlers

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
	Type         string        `json:"type"`
	ResourceType string        `json:"resourceType"`
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
	Type         string        `json:"type"`
	ResourceType string        `json:"resourceType"`
	Cost         float32       `json:"cost"`
	BusinessID   uuid.UUID     `json:"businessId"`
	DetectionID  *uuid.UUID    `json:"detectionId"`
	MitigationID *uuid.UUID    `json:"mitigationId"`
	ResourceID   *uuid.UUID    `json:"resourceId"`
	ThreatID     *uuid.UUID    `json:"threatId"`
	ImpactID     *uuid.UUID    `json:"impactId"`
	CreatedAt    time.Time     `json:"createdAt"`
}
type LiabilityOutputs []LiabilityOutput

func (liabilityInput LiabilityInput) inputToModel() (liabilityModel database.LiabilityModel, err error) {
	liabilityModel.ID = liabilityInput.ID
	liabilityModel.Name = liabilityInput.Name
	liabilityModel.Description = liabilityInput.Description
	liabilityModel.Quantity = liabilityInput.Quantity
	liabilityModel.Type = liabilityInput.Type
	liabilityModel.ResourceType = liabilityInput.ResourceType
	liabilityModel.Cost = liabilityInput.Cost
	liabilityModel.BusinessID = liabilityInput.BusinessID
	liabilityModel.MitigationID = liabilityInput.MitigationID
	liabilityModel.ResourceID = liabilityInput.ResourceID
	liabilityModel.ThreatID = liabilityInput.ThreatID
	liabilityModel.ImpactID = liabilityInput.ImpactID
	liabilityModel.CreatedAt = liabilityInput.CreatedAt
	return
}

func (liabilityOutput *LiabilityOutput) modelToOutput(liabilityModel database.LiabilityModel) (err error) {
	liabilityOutput.ID = liabilityModel.ID
	liabilityOutput.Name = liabilityModel.Name
	liabilityOutput.Description = liabilityModel.Description
	liabilityOutput.Quantity = liabilityModel.Quantity
	liabilityOutput.Type = liabilityModel.Type
	liabilityOutput.ResourceType = liabilityModel.ResourceType
	liabilityOutput.Cost = liabilityModel.Cost
	liabilityOutput.BusinessID = liabilityModel.BusinessID
	liabilityOutput.MitigationID = liabilityModel.MitigationID
	liabilityOutput.ResourceID = liabilityModel.ResourceID
	liabilityOutput.ThreatID = liabilityModel.ThreatID
	liabilityOutput.ImpactID = liabilityModel.ImpactID
	liabilityOutput.CreatedAt = liabilityModel.CreatedAt
	return
}

func liabilityModelsToOutputs(liabilityModels []database.LiabilityModel) (liabilityOutputs LiabilityOutputs, err error) {
	for _, liabilityModel := range liabilityModels {
		liabilityOutput := LiabilityOutput{}
		err := liabilityOutput.modelToOutput(liabilityModel)
		if err != nil {
			return nil, err
		}
		liabilityOutputs = append(liabilityOutputs, liabilityOutput)
	}
	return
}

func (controller PublicController) GetLiabilities(context *gin.Context) {
	var liabilityModels []database.LiabilityModel

	rawBusinessId, ok := context.GetQuery("businessId")
	if !ok {
		log.Println("Parameter businessId not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	businessId, err := uuid.Parse(rawBusinessId)
	if err != nil {
		log.Println("businessId is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	rawImpactId, impactIdPresent := context.GetQuery("impactId")
	rawMitigationId, mitigationIdPresent := context.GetQuery("mitigationId")
	rawThreatId, threatIdPresent := context.GetQuery("threatId")

	switch {
	case threatIdPresent:
		threatId, err := uuid.Parse(rawThreatId)
		if err != nil {
			log.Println("threatId is not a uuid")
			context.JSON(http.StatusNotFound, "Not found")
			return
		}
		liabilityModels, err = controller.DBManager.GetLiabilitiesByThreatId(businessId.String(), threatId.String())
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusNotFound, "Not Found")
			return
		}
	case mitigationIdPresent:
		mitigationId, err := uuid.Parse(rawMitigationId)
		if err != nil {
			log.Println("mitigationId is not a uuid")
			context.JSON(http.StatusNotFound, "Not found")
			return
		}
		liabilityModels, err = controller.DBManager.GetLiabilitiesByMitigationId(businessId.String(), mitigationId.String())
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusNotFound, "Not Found")
			return
		}
	case impactIdPresent:
		impactId, err := uuid.Parse(rawImpactId)
		if err != nil {
			log.Println("impactId is not a uuid")
			context.JSON(http.StatusNotFound, "Not found")
			return
		}
		liabilityModels, err = controller.DBManager.GetLiabilitiesByImpactId(businessId.String(), impactId.String())
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusNotFound, "Not Found")
			return
		}
	default:
		liabilityModels, err = controller.DBManager.GetLiabilities(businessId.String())
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusNotFound, "Not Found")
			return
		}

	}

	liabilityOutputs, err := liabilityModelsToOutputs(liabilityModels)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}

	context.JSON(http.StatusOK, liabilityOutputs)
}

func (controller PublicController) GetLiability(context *gin.Context) {

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

	liabilityModel, err := controller.DBManager.GetLiability(liabilityId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}
	var liabilityOutput LiabilityOutput

	err = liabilityOutput.modelToOutput(liabilityModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not Found")
		return
	}
	context.JSON(http.StatusOK, liabilityOutput)
}

// func (controller PublicController) GetLiabilityByImpactId(context *gin.Context) {

// 	id, ok := context.GetQuery("impactId")
// 	if !ok {
// 		log.Println("Parameter impactId not found")
// 		context.JSON(http.StatusNotFound, "Not found")
// 		return
// 	}

// 	liabilityId, err := uuid.Parse(id)
// 	if err != nil {
// 		log.Println("impactId is not a uuid")
// 		context.JSON(http.StatusNotFound, "Not found")
// 		return
// 	}

// 	liabilityModel, err := controller.DBManager.GetLiabilityByImpactId(liabilityId.String())

// 	if err != nil {
// 		log.Println(err)
// 		context.JSON(http.StatusNotFound, "Not Found")
// 		return
// 	}
// 	var liabilityOutputs LiabilityOutputs
// 	err = liabilityOutputs.modelsToOutputs(liabilityModel)

// 	if err != nil {
// 		log.Println(err)
// 		context.JSON(http.StatusNotFound, "Not Found")
// 		return
// 	}
// 	context.JSON(http.StatusOK, liabilityOutputs)
// }

func (controller PublicController) DeleteLiability(context *gin.Context) {

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

	err = controller.DBManager.DeleteLiability(liabilityId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateLiability(context *gin.Context) {

	liabilityInput := LiabilityInput{}
	err := context.ShouldBindJSON(&liabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	liabilityModel, err := liabilityInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	liabilityId, err := controller.DBManager.CreateLiability(liabilityModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, liabilityId)
}

func (controller PublicController) UpdateLiability(context *gin.Context) {

	liabilityInput := LiabilityInput{}
	err := context.ShouldBindJSON(&liabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	liabilityModel, err := liabilityInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = controller.DBManager.UpdateLiability(liabilityModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
