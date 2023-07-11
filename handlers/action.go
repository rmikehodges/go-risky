package handlers

import (
	"fmt"
	"go-risky/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller PublicController) GetActions(context *gin.Context) {
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

	actionOutputs, err := controller.DBManager.GetActions(businessId.String())

	if err != nil {
		log.Printf("Get Actions Error %s", err)
		context.JSON(http.StatusNotFound, actionOutputs)
		return
	}

	context.JSON(http.StatusOK, actionOutputs)
}

func (controller PublicController) GetAction(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	actionId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	actionOutput, err := controller.DBManager.GetAction(actionId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, actionOutput)
		return
	}

	context.JSON(http.StatusOK, actionOutput)
}

func (controller PublicController) DeleteAction(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	actionId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = controller.DBManager.DeleteAction(actionId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateAction(context *gin.Context) {

	actionInput := types.Action{}
	err := context.ShouldBindJSON(&actionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	_, err = controller.DBManager.CreateAction(actionInput)
	fmt.Println("returned from create action")
	if err != nil {
		log.Println(err)
		fmt.Println("Error creating action")
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func (controller PublicController) UpdateAction(context *gin.Context) {

	actionInput := types.Action{}
	err := context.ShouldBindJSON(&actionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	err = controller.DBManager.UpdateAction(actionInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
