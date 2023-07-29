package handlers

import (
	"fmt"
	"go-risky/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller PublicController) GetGroups(context *gin.Context) {
	id, ok := context.GetQuery("organizationId")
	if !ok {
		log.Println("Parameter organizationId not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	organizationId, err := uuid.Parse(id)
	if err != nil {
		log.Println("organizationId is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	groupOutputs, err := controller.DBManager.GetGroups(organizationId.String())

	if err != nil {
		log.Printf("Get Groups Error %s", err)
		context.JSON(http.StatusNotFound, groupOutputs)
		return
	}

	context.JSON(http.StatusOK, groupOutputs)
}


func (controller PublicController) GetGroup(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	groupId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	groupOutput, err := controller.DBManager.GetGroup(groupId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, groupOutput)
		return
	}

	context.JSON(http.StatusOK, groupOutput)
}

func (controller PublicController) DeleteGroup(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	groupId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = controller.DBManager.DeleteGroup(groupId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateGroup(context *gin.Context) {

	groupInput := types.Group{}
	err := context.ShouldBindJSON(&groupInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	_, err = controller.DBManager.CreateGroup(groupInput)
	fmt.Println("returned from create group")
	if err != nil {
		log.Println(err)
		fmt.Println("Error creating group")
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func (controller PublicController) UpdateGroup(context *gin.Context) {

	groupInput := types.Group{}
	err := context.ShouldBindJSON(&groupInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	err = controller.DBManager.UpdateGroup(groupInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
