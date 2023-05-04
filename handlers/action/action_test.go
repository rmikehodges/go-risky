package action

import (
	"bytes"
	"encoding/json"
	"go-risky/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/google/uuid"
)

//Functions to test handlers/action.go

var businessId = "568d1c21-c83f-4f4f-815b-9ee6490fe8f7"
var capabilityId = "d27d7d83-0145-4caf-aab3-b250c23f07b5"
var actionId = "a32ec646-48e8-400b-b2aa-626c3085f1bf"

func TestGetActions(t *testing.T) {
	router := gin.Default()
	router.GET("/actions", getActions)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/actions?businessId="+businessId, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetAction(t *testing.T) {
	router := gin.Default()
	router.GET("/action", getAction)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/action?id="+actionId, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateAction(t *testing.T) {
	router := gin.Default()
	router.POST("/action", createAction)
	w := httptest.NewRecorder()
	//Create a new ActionInput struct with prepared data
	capabilityId := uuid.MustParse(capabilityId)
	actionInput := ActionInput{
		Name:         "Test",
		Description:  "test",
		BusinessID:   uuid.MustParse(businessId),
		CapabilityID: &capabilityId,
		Complexity:   "LOW",
	}

	actionInputJson, err := json.Marshal(actionInput)

	if err != nil {
		t.Errorf("Error while marshalling actionInputJson: %v\n", err)
	}

	req, _ := http.NewRequest("POST", "http://localhost:8081/action", bytes.NewBuffer(actionInputJson))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUpdateAction(t *testing.T) {
	router := gin.Default()
	router.PATCH("/action", updateAction)
	w := httptest.NewRecorder()

	capabilityId := uuid.MustParse(capabilityId)
	actionInput := ActionInput{
		ID:           uuid.MustParse(actionId),
		Name:         "Test",
		Description:  "test",
		BusinessID:   uuid.MustParse(businessId),
		CapabilityID: &capabilityId,
		Complexity:   "LOW",
	}

	actionInputJson, err := json.Marshal(actionInput)

	if err != nil {
		t.Errorf("Error while marshalling actionInputJson: %v\n", err)
	}

	req, _ := http.NewRequest("PATCH", "/action", bytes.NewBuffer(actionInputJson))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteAction(t *testing.T) {
	router := gin.Default()
	router.DELETE("/action", deleteAction)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/action", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestInputToModel(t *testing.T) {
	actionInput := ActionInput{
		ID:   uuid.New(),
		Name: "Test",
	}
	_, err := inputToModel(actionInput)
	assert.Equal(t, nil, err)
}

func TestModelToOutput(t *testing.T) {
	actionModel := database.ActionModel{
		ID:   uuid.New(),
		Name: "Test",
	}
	_, err := modelToOutput(actionModel)
	assert.Equal(t, nil, err)
}
