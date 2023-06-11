package handlers

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

type ThreatOutputs []ThreatOutput

func (threatInput ThreatInput) inputToModel() (threatModel database.ThreatModel, err error) {
	//This is where you do input validation sanitization
	threatModel.ID = threatInput.ID
	threatModel.Name = threatInput.Name
	threatModel.BusinessID = threatInput.BusinessID
	threatModel.CreatedAt = threatInput.CreatedAt

	return

}

func (threatOutput *ThreatOutput) modelToOutput(threatModel database.ThreatModel) (err error) {
	//This is where you do input validation sanitization
	threatOutput.ID = threatModel.ID
	threatOutput.Name = threatModel.Name
	threatOutput.BusinessID = threatModel.BusinessID
	threatOutput.CreatedAt = threatModel.CreatedAt

	return
}

func threatModelsToOutput(threatModels []database.ThreatModel) (threatOutputs ThreatOutputs, err error) {
	//This is where you do input validation sanitization
	for _, model := range threatModels {
		threatOutput := ThreatOutput{}
		err := threatOutput.modelToOutput(model)
		if err != nil {
			return nil, err
		}
		threatOutputs = append(threatOutputs, threatOutput)
	}

	return
}

func (controller PublicController) GetThreats(context *gin.Context) {

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

	threatModel, err := controller.DBManager.GetThreats(businessId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, threatModel)
		return
	}

	threatOutputs, err := threatModelsToOutput(threatModel)

	if err != nil {
		context.JSON(http.StatusNotFound, threatOutputs)
		return
	}

	context.JSON(http.StatusOK, threatOutputs)
}

func (controller PublicController) GetThreat(context *gin.Context) {

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

	threatModel, err := controller.DBManager.GetThreat(threatId.String())

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, nil)
		return
	}

	var threatOutput ThreatOutput
	err = threatOutput.modelToOutput(threatModel)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, threatOutput)
		return
	}

	context.JSON(http.StatusOK, threatOutput)
}

func (controller PublicController) DeleteThreat(context *gin.Context) {

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

	err = controller.DBManager.DeleteThreat(threatId.String())

	if err != nil {
		log.Println("Received Error from Database")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	context.JSON(http.StatusOK, "Success")
}

func (controller PublicController) CreateThreat(context *gin.Context) {

	threatInput := ThreatInput{}
	err := context.ShouldBindJSON(&threatInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}
	threatModel, err := threatInput.inputToModel()
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	threatId, err := controller.DBManager.CreateThreat(threatModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, threatId)
}

func (controller PublicController) UpdateThreat(context *gin.Context) {

	threatInput := ThreatInput{}
	err := context.ShouldBindJSON(&threatInput)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	threatModel, err := threatInput.inputToModel()

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	err = controller.DBManager.UpdateThreat(threatModel)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
