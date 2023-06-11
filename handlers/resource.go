package handlers

import (
	"go-risky/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type ResourceInput struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  zeronull.Text `json:"description"`
	Cost         float32       `json:"cost" db:"cost"`
	Unit         string        `json:"unit" db:"unit"`
	Total        float32       `json:"total"`
	ResourceType string        `json:"resourceType" db:"resource_type"`
	BusinessID   uuid.UUID     `json:"businessId" db:"business_id"`
}

type ResourceOutput struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  zeronull.Text `json:"description"`
	Cost         float32       `json:"cost" db:"cost"`
	Unit         string        `json:"unit" db:"unit"`
	Total        float32       `json:"total"`
	ResourceType string        `json:"resourceType" db:"resource_type"`
	BusinessID   uuid.UUID     `json:"businessId" db:"business_id"`
	CreatedAt    time.Time     `json:"createdAt" db:"created_at"`
}

type ResourceOutputs []ResourceOutput

func (resourceInput ResourceInput) inputToModel() (resourceModel database.ResourceModel, err error) {
	resourceModel.ID = resourceInput.ID
	resourceModel.Name = resourceInput.Name
	resourceModel.Description = resourceInput.Description
	resourceModel.Cost = resourceInput.Cost
	resourceModel.Unit = resourceInput.Unit
	resourceModel.Total = resourceInput.Total
	resourceModel.ResourceType = resourceInput.ResourceType
	resourceModel.BusinessID = resourceInput.BusinessID
	return
}

func (resourceOutput *ResourceOutput) modelToOutput(resourceModel database.ResourceModel) (err error) {
	resourceOutput.ID = resourceModel.ID
	resourceOutput.Name = resourceModel.Name
	resourceOutput.Description = resourceModel.Description
	resourceOutput.Cost = resourceModel.Cost
	resourceOutput.Unit = resourceModel.Unit
	resourceOutput.Total = resourceModel.Total
	resourceOutput.ResourceType = resourceModel.ResourceType
	resourceOutput.BusinessID = resourceModel.BusinessID
	resourceOutput.CreatedAt = resourceModel.CreatedAt
	return
}

func resourceModelsToOutputs(resourceModels []database.ResourceModel) (resourceOutputs ResourceOutputs, err error) {
	for _, resourceModel := range resourceModels {
		resourceOutput := ResourceOutput{}
		err := resourceOutput.modelToOutput(resourceModel)
		if err != nil {
			return nil, err
		}
		resourceOutputs = append(resourceOutputs, resourceOutput)
	}
	return
}

func (controller PublicController) GetResources(context *gin.Context) {

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

	resourcesModels, err := controller.DBManager.GetResources(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	resourceOutputs, err := resourceModelsToOutputs(resourcesModels)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, resourceOutputs)
}

func (controller PublicController) GetResource(context *gin.Context) {

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

	resourceModel, err := controller.DBManager.GetResource(resourceId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	var resourceOutput ResourceOutput

	err = resourceOutput.modelToOutput(resourceModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, resourceOutput)
}

func (controller PublicController) DeleteResource(context *gin.Context) {

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

	err = controller.DBManager.DeleteResource(resourceId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateResource(context *gin.Context) {

	resourceInput := ResourceInput{}
	err := context.ShouldBindJSON(&resourceInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	resourceModel, err := resourceInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	resourceId, err := controller.DBManager.CreateResource(resourceModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, resourceId)
}

func (controller PublicController) UpdateResource(context *gin.Context) {

	resourceInput := ResourceInput{}
	err := context.ShouldBindJSON(&resourceInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	resourceModel, err := resourceInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = controller.DBManager.UpdateResource(resourceModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
