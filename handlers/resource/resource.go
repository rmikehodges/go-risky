package resource

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

func inputToModel(resourceInput ResourceInput) (resourceModel database.ResourceModel, err error) {
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

func modelToOutput(resourceModel database.ResourceModel) (resourceOutput ResourceOutput, err error) {
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

func modelsToOutputs(resourceModels []database.ResourceModel) (resourceOutputs []ResourceOutput, err error) {
	for _, resourceModel := range resourceModels {
		resourceOutput, err := modelToOutput(resourceModel)
		if err != nil {
			return resourceOutputs, err
		}
		resourceOutputs = append(resourceOutputs, resourceOutput)
	}
	return
}

func getResources(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

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

	resourcesModels, err := db.GetResources(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	resourceOutput, err := modelsToOutputs(resourcesModels)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, resourceOutput)
}

func getResource(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

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

	resourceModel, err := db.GetResource(resourceId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	resourceOutput, err := modelToOutput(resourceModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, resourceOutput)
}

func deleteResource(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

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

	err = db.DeleteResource(resourceId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createResource(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	resourceInput := ResourceInput{}
	err := context.ShouldBindJSON(&resourceInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	resourceModel, err := inputToModel(resourceInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = db.CreateResource(resourceModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateResource(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	resourceInput := ResourceInput{}
	err := context.ShouldBindJSON(&resourceInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	resourceModel, err := inputToModel(resourceInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = db.UpdateResource(resourceModel)
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
