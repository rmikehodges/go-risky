package handlers

import (
	"log"
	"net/http"

	"go-risky/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

	assetOutputs, err := controller.DBManager.GetAssets(businessId.String())
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

	assetOutput, err := controller.DBManager.GetAsset(assetId.String())
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

	assetInput := types.Asset{}
	err := context.ShouldBindJSON(&assetInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	assetId, err := controller.DBManager.CreateAsset(assetInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, assetId)
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func (controller PublicController) UpdateAsset(context *gin.Context) {

	assetInput := types.Asset{}
	err := context.ShouldBindJSON(&assetInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	err = controller.DBManager.UpdateAsset(assetInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
