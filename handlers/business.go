package handlers

import (
	"fmt"
	"go-risky/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller PublicController) GetBusinesses(context *gin.Context) {

	businessOutputs, err := controller.DBManager.GetBusinesses()

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, businessOutputs)
		return
	}

	context.JSON(http.StatusOK, businessOutputs)
}

func (controller PublicController) GetBusiness(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	businessId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	businessOutput, err := controller.DBManager.GetBusiness(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, businessOutput)
		return
	}

	context.JSON(http.StatusOK, businessOutput)
}

func (controller PublicController) DeleteBusiness(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	businessId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = controller.DBManager.DeleteBusiness(businessId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, businessId.String()+" deleted")
}

func (controller PublicController) CreateBusiness(context *gin.Context) {

	businessInput := types.Business{}
	err := context.ShouldBindJSON(&businessInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	businessOutput, err := controller.DBManager.CreateBusiness(businessInput)
	fmt.Println("returned from create business")
	if err != nil {
		log.Println(err)
		fmt.Println("Error creating business")
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, businessOutput)
}

func (controller PublicController) UpdateBusiness(context *gin.Context) {

	businessInput := types.Business{}
	err := context.ShouldBindJSON(&businessInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	err = controller.DBManager.UpdateBusiness(businessInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, businessInput.ID.String()+" updated")
}
