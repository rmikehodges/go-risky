package handlers

import (
	"fmt"
	"go-risky/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller PublicController) GetUsers(context *gin.Context) {
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

	userOutputs, err := controller.DBManager.GetUsers(organizationId.String())

	if err != nil {
		log.Printf("Get Users Error %s", err)
		context.JSON(http.StatusNotFound, userOutputs)
		return
	}

	context.JSON(http.StatusOK, userOutputs)
}

func (controller PublicController) GetUsersByGroup(context *gin.Context) {
	id, ok := context.GetQuery("groupId")
	if !ok {
		log.Println("Parameter groupId not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	groupId, err := uuid.Parse(id)
	if err != nil {
		log.Println("groupId is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	userOutputs, err := controller.DBManager.GetUsers(groupId.String())

	if err != nil {
		log.Printf("Get Users Error %s", err)
		context.JSON(http.StatusNotFound, userOutputs)
		return
	}

	context.JSON(http.StatusOK, userOutputs)
}

func (controller PublicController) GetUser(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	userId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	userOutput, err := controller.DBManager.GetUser(userId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, userOutput)
		return
	}

	context.JSON(http.StatusOK, userOutput)
}

func (controller PublicController) DeleteUser(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	userId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = controller.DBManager.DeleteUser(userId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateUser(context *gin.Context) {

	userInput := types.User{}
	err := context.ShouldBindJSON(&userInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	_, err = controller.DBManager.CreateUser(userInput)
	fmt.Println("returned from create user")
	if err != nil {
		log.Println(err)
		fmt.Println("Error creating user")
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func (controller PublicController) UpdateUser(context *gin.Context) {

	userInput := types.User{}
	err := context.ShouldBindJSON(&userInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	err = controller.DBManager.UpdateUser(userInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
