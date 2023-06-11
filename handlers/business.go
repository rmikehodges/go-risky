package handlers

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

type BusinessOutputs []BusinessOutput

func (businessInput BusinessInput) inputToModel() (businessModel database.BusinessModel, err error) {
	//This is where you do input validation sanitization
	businessModel.ID = businessInput.ID
	businessModel.Name = businessInput.Name
	businessModel.Revenue = businessInput.Revenue
	businessModel.CreatedAt = businessInput.CreatedAt

	return

}

func (businessOutput *BusinessOutput) modelToOutput(businessModel database.BusinessModel) (err error) {
	//This is where you do input validation sanitization
	businessOutput.ID = businessModel.ID
	businessOutput.Name = businessModel.Name
	businessOutput.Revenue = businessModel.Revenue
	businessOutput.CreatedAt = businessModel.CreatedAt

	return
}

func businessModelsToOutput(businessModels []database.BusinessModel) (businessOutputs BusinessOutputs, err error) {
	//This is where you do input validation sanitization
	for _, model := range businessModels {
		businessOutput := BusinessOutput{}
		err := businessOutput.modelToOutput(model)
		if err != nil {
			return nil, err
		}
		businessOutputs = append(businessOutputs, businessOutput)
	}

	return
}

func (controller PublicController) GetBusinesses(context *gin.Context) {

	businessmodel, err := controller.DBManager.GetBusinesses()

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, businessmodel)
		return
	}

	businessOutputs, err := businessModelsToOutput(businessmodel)

	if err != nil {
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

	businessInput := BusinessInput{}
	err := context.ShouldBindJSON(&businessInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}
	businessModel, err := businessInput.inputToModel()
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	businessOutput, err := controller.DBManager.CreateBusiness(businessModel)
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

	businessInput := BusinessInput{}
	err := context.ShouldBindJSON(&businessInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	businessModel, err := businessInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	err = controller.DBManager.UpdateBusiness(businessModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, businessModel.ID.String()+" updated")
}
