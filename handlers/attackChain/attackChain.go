package attackChain

import (
	"log"
	"net/http"

	"go-risky/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getAttackChains(context *gin.Context) {
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

	attackChainOutput, err := database.GetAttackChains(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, attackChainOutput)
		return
	}

	context.JSON(http.StatusOK, attackChainOutput)
}

func getAttackChain(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	attackChainId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	attackChainOutput, err := database.GetAttackChain(attackChainId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, attackChainOutput)
		return
	}

	context.JSON(http.StatusOK, attackChainOutput)
}

func deleteAttackChain(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	attackChainId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = database.DeleteAttackChain(attackChainId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createAttackChain(context *gin.Context) {
	attackChainInput := database.AttackChain{}
	err := context.ShouldBindJSON(&attackChainInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.CreateAttackChain(attackChainInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateAttackChain(context *gin.Context) {
	attackChainInput := database.AttackChain{}
	err := context.ShouldBindJSON(&attackChainInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	err = database.UpdateAttackChain(attackChainInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func AttackChainRoutes(router *gin.Engine) {
	router.GET("/attackChains", getAttackChains)
	router.GET("/attackChain", getAttackChain)
	router.DELETE("/attackChain", deleteAttackChain)
	router.PATCH("/attackChain", updateAttackChain)
	router.POST("/attackChain", createAttackChain)
}
