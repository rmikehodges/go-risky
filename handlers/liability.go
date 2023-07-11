package handlers

import (
	"go-risky/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller PublicController) GetLiabilities(context *gin.Context) {

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

	var liabilityOutputs types.Liabilities
	switch {
	case threatIdPresent:
		threatId, err := uuid.Parse(rawThreatId)
		if err != nil {
			log.Println("threatId is not a uuid")
			context.JSON(http.StatusNotFound, "Not found")
			return
		}
		liabilityOutputs, err = controller.DBManager.GetLiabilitiesByThreatId(businessId.String(), threatId.String())
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
		liabilityOutputs, err = controller.DBManager.GetLiabilitiesByMitigationId(businessId.String(), mitigationId.String())
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
		liabilityOutputs, err = controller.DBManager.GetLiabilitiesByImpactId(businessId.String(), impactId.String())
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusNotFound, "Not Found")
			return
		}
	default:
		liabilityOutputs, err = controller.DBManager.GetLiabilities(businessId.String())
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusNotFound, "Not Found")
			return
		}

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

	liabilityOutput, err := controller.DBManager.GetLiability(liabilityId.String())

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

	liabilityInput := types.Liability{}
	err := context.ShouldBindJSON(&liabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	liabilityId, err := controller.DBManager.CreateLiability(liabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, liabilityId)
}

func (controller PublicController) UpdateLiability(context *gin.Context) {

	liabilityInput := types.Liability{}
	err := context.ShouldBindJSON(&liabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = controller.DBManager.UpdateLiability(liabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
