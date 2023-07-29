package handlers

import (
	"fmt"
	"go-risky/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (controller PublicController) GetOrganizations(context *gin.Context) {

	organizationOutputs, err := controller.DBManager.GetOrganizations()

	if err != nil {
		log.Printf("Get Organizations Error %s", err)
		context.JSON(http.StatusNotFound, organizationOutputs)
		return
	}

	context.JSON(http.StatusOK, organizationOutputs)
}

func (controller PublicController) GetOrganization(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	organizationId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	organizationOutput, err := controller.DBManager.GetOrganization(organizationId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, organizationOutput)
		return
	}

	context.JSON(http.StatusOK, organizationOutput)
}

func (controller PublicController) DeleteOrganization(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	organizationId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = controller.DBManager.DeleteOrganization(organizationId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateOrganization(context *gin.Context) {

	organizationInput := types.Organization{}
	err := context.ShouldBindJSON(&organizationInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	_, err = controller.DBManager.CreateOrganization(organizationInput)
	fmt.Println("returned from create organization")
	if err != nil {
		log.Println(err)
		fmt.Println("Error creating organization")
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func (controller PublicController) UpdateOrganization(context *gin.Context) {

	organizationInput := types.Organization{}
	err := context.ShouldBindJSON(&organizationInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	err = controller.DBManager.UpdateOrganization(organizationInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
