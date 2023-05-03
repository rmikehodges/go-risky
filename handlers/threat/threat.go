package threat

import (
	"go-risky/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ThreatInput struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	BusinessID  uuid.UUID `json:"businessId"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ThreatOutput struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	BusinessID  uuid.UUID `json:"businessId"`
	CreatedAt   time.Time `json:"createdAt"`
}

func inputToModel(threatInput ThreatInput) (threatModel database.ThreatModel, err error) {
	//This is where you do input validation sanitization
	threatModel.ID = threatInput.ID
	threatModel.Name = threatInput.Name
	threatModel.BusinessID = threatInput.BusinessID
	threatModel.CreatedAt = threatInput.CreatedAt

	return

}

func modelToOutput(threatModel database.ThreatModel) (threatOutput ThreatOutput, err error) {
	//This is where you do input validation sanitization
	threatOutput.ID = threatModel.ID
	threatOutput.Name = threatModel.Name
	threatOutput.BusinessID = threatModel.BusinessID
	threatOutput.CreatedAt = threatModel.CreatedAt

	return
}

func modelsToOutput(threatModels []database.ThreatModel) (threatOutput []ThreatOutput, err error) {
	//This is where you do input validation sanitization
	for _, model := range threatModels {
		output, err := modelToOutput(model)
		if err != nil {
			return []ThreatOutput{}, err
		}
		threatOutput = append(threatOutput, output)
	}

	return
}

func getThreats(context *gin.Context) {
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

	threatmodel, err := database.GetThreats(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, threatmodel)
		return
	}

	threatOutput, err := modelsToOutput(threatmodel)

	if err != nil {
		context.JSON(http.StatusNotFound, threatOutput)
		return
	}

	context.JSON(http.StatusOK, threatOutput)
}

func getThreat(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	threatId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	threatOutput, err := database.GetThreat(threatId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, threatOutput)
		return
	}

	context.JSON(http.StatusOK, threatOutput)
}

func deleteThreat(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	threatId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	err = database.DeleteThreat(threatId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func createThreat(context *gin.Context) {
	threatInput := ThreatInput{}
	err := context.ShouldBindJSON(&threatInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	threatModel, err := inputToModel(threatInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	err = database.CreateThreat(threatModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func updateThreat(context *gin.Context) {
	threatInput := ThreatInput{}
	err := context.ShouldBindJSON(&threatInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	threatModel, err := inputToModel(threatInput)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	err = database.UpdateThreat(threatModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func ThreatRoutes(router *gin.Engine) {
	router.GET("/threats", getThreats)
	router.GET("/threat/:id", getThreat)
	router.DELETE("/threat/:id", deleteThreat)
	router.PATCH("/threat/:id", updateThreat)
	router.POST("/threats", createThreat)
}
