package liability

import (
	"go-risky/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getLiabilities(context *gin.Context) {
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

	liabilityOutput, err := database.GetLiabilities(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, liabilityOutput)
		return
	}

	context.JSON(http.StatusOK, liabilityOutput)
}

func getLiability(context *gin.Context) {
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

	liabilityOutput, err := database.GetLiability(liabilityId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, liabilityOutput)
		return
	}

	context.JSON(http.StatusOK, liabilityOutput)
}

func deleteLiability(context *gin.Context) {
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

	err = database.DeleteLiability(liabilityId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createLiability(context *gin.Context) {
	liabilityInput := database.Liability{}
	err := context.ShouldBindJSON(&liabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.CreateLiability(liabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateLiability(context *gin.Context) {
	liabilityInput := database.Liability{}
	err := context.ShouldBindJSON(&liabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.UpdateLiability(liabilityInput)
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
	router.DELETE("/liability", deleteLiability)
	router.PATCH("/liability", updateLiability)
	router.POST("/liabilities", createLiability)
}
