package business

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go-risky/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BusinessInput struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Revenue   float32   `json:"revenue"`
	CreatedAt time.Time `json:"createdAt"`
}

type BusinessOutput struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Revenue   float32   `json:"revenue"`
	CreatedAt time.Time `json:"createdAt"`
}

func inputToModel(businessInput BusinessInput) (businessModel database.BusinessModel, err error) {
	//This is where you do input validation sanitization
	businessModel.ID = businessInput.ID
	businessModel.Name = businessInput.Name
	businessModel.Revenue = businessInput.Revenue
	businessModel.CreatedAt = businessInput.CreatedAt

	return

}

func modelToOutput(businessModel database.BusinessModel) (businessOutput BusinessOutput, err error) {
	//This is where you do input validation sanitization
	businessOutput.ID = businessModel.ID
	businessOutput.Name = businessModel.Name
	businessOutput.Revenue = businessModel.Revenue
	businessOutput.CreatedAt = businessModel.CreatedAt

	return
}

func modelsToOutput(businessModels []database.BusinessModel) (businessOutput []BusinessOutput, err error) {
	//This is where you do input validation sanitization
	for _, model := range businessModels {
		output, err := modelToOutput(model)
		if err != nil {
			return []BusinessOutput{}, err
		}
		businessOutput = append(businessOutput, output)
	}

	return
}

func getBusinesses(context *gin.Context) {

	businessmodel, err := database.GetBusinesses()

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, businessmodel)
		return
	}

	businessOutput, err := modelsToOutput(businessmodel)

	if err != nil {
		context.JSON(http.StatusNotFound, businessOutput)
		return
	}

	context.JSON(http.StatusOK, businessOutput)
}

func getBusiness(context *gin.Context) {
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

	businessOutput, err := database.GetBusiness(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, businessOutput)
		return
	}

	context.JSON(http.StatusOK, businessOutput)
}

func deleteBusiness(context *gin.Context) {
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

	businessOutput, err := database.DeleteBusiness(businessId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, businessOutput)
}

func createBusiness(context *gin.Context) {
	businessInput := BusinessInput{}
	err := context.ShouldBindJSON(&businessInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}
	businessModel, err := inputToModel(businessInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	businessOutput, err := database.CreateBusiness(businessModel)
	fmt.Println("returned from create business")
	if err != nil {
		log.Println(err)
		fmt.Println("Error creating business")
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, businessOutput)
}

func updateBusiness(context *gin.Context) {
	businessInput := BusinessInput{}
	err := context.ShouldBindJSON(&businessInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	businessModel, err := inputToModel(businessInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	businessOutput, err := database.UpdateBusiness(businessModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, businessOutput)
}

func BusinessRoutes(router *gin.Engine) {
	router.GET("/businesses", getBusinesses)
	router.GET("/business", getBusiness)
	router.DELETE("/business", deleteBusiness)
	router.PATCH("/business", updateBusiness)
	router.POST("/business", createBusiness)
}
