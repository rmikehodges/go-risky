package capability

import (
	"go-risky/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

//Create types CapabilityInput and CapabilityOutput that match the database model CapabilityModel

type CapabilityInput struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	CreatedAt   time.Time     `json:"createdAt"`
}

type CapabilityOutput struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId"`
	CreatedAt   time.Time     `json:"createdAt"`
}

//Create functions modelToOutput, inputToModel, and modelsToOutput that convert between the database model and the input/output types

func modelToOutput(capabilityModel database.CapabilityModel) (capabilityOutput CapabilityOutput, err error) {
	//This is where you do input validation sanitization
	capabilityOutput.ID = capabilityModel.ID
	capabilityOutput.Name = capabilityModel.Name
	capabilityOutput.Description = capabilityModel.Description
	capabilityOutput.BusinessID = capabilityModel.BusinessID
	capabilityOutput.CreatedAt = capabilityModel.CreatedAt
	return
}

func inputToModel(capabilityInput CapabilityInput) (capabilityModel database.CapabilityModel, err error) {
	//This is where you do input validation sanitization
	capabilityModel.ID = capabilityInput.ID
	capabilityModel.Name = capabilityInput.Name
	capabilityModel.Description = capabilityInput.Description
	capabilityModel.BusinessID = capabilityInput.BusinessID
	capabilityModel.CreatedAt = capabilityInput.CreatedAt
	return
}

func modelsToOutput(capabilityModels []database.CapabilityModel) (capabilityOutputs []CapabilityOutput, err error) {
	for _, capabilityModel := range capabilityModels {
		capabilityOutput, err := modelToOutput(capabilityModel)
		if err != nil {
			return capabilityOutputs, err
		}
		capabilityOutputs = append(capabilityOutputs, capabilityOutput)
	}
	return
}

func getCapabilities(context *gin.Context) {
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

	capabilityModel, err := db.GetCapabilities(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, capabilityModel)
		return
	}

	capabilityOutput, err := modelsToOutput(capabilityModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, capabilityOutput)
		return
	}

	context.JSON(http.StatusOK, capabilityOutput)
}

func getCapability(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	capabilityId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	capabilityOutput, err := db.GetCapability(capabilityId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, capabilityOutput)
		return
	}

	context.JSON(http.StatusOK, capabilityOutput)
}

func deleteCapability(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	capabilityId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = db.DeleteCapability(capabilityId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createCapability(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	capabilityInput := CapabilityInput{}
	err := context.ShouldBindJSON(&capabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	capabilityModel, err := inputToModel(capabilityInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = db.CreateCapability(capabilityModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateCapability(context *gin.Context) {
	db := context.MustGet("DBManager").(*database.DBManager)

	capabilityInput := CapabilityInput{}
	err := context.ShouldBindJSON(&capabilityInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	capabilityModel, err := inputToModel(capabilityInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
		return
	}
	err = db.UpdateCapability(capabilityModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func CapabilityRoutes(router *gin.Engine) {
	router.GET("/capabilities", getCapabilities)
	router.GET("/capability", getCapability)
	router.DELETE("/capability", deleteCapability)
	router.PATCH("/capability", updateCapability)
	router.POST("/capabilities", createCapability)
}
