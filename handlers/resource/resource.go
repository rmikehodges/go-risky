package resource

import (
	"go-risky/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getResources(context *gin.Context) {
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

	resourceOutput, err := database.GetResources(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, resourceOutput)
		return
	}

	context.JSON(http.StatusOK, resourceOutput)
}

func getResource(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	resourceId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	resourceOutput, err := database.GetResource(resourceId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, resourceOutput)
		return
	}

	context.JSON(http.StatusOK, resourceOutput)
}

func deleteResource(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	resourceId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = database.DeleteResource(resourceId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createResource(context *gin.Context) {
	resourceInput := database.ResourceModel{}
	err := context.ShouldBindJSON(&resourceInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.CreateResource(resourceInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateResource(context *gin.Context) {
	resourceInput := database.ResourceModel{}
	err := context.ShouldBindJSON(&resourceInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.UpdateResource(resourceInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func ResourceRoutes(router *gin.Engine) {
	router.GET("/resources", getResources)
	router.GET("/resource/:id", getResource)
	router.DELETE("/resource/:id", deleteResource)
	router.PATCH("/resource/:id", updateResource)
	router.POST("/resources", createResource)
}
