package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"go-risky/database"
	"go-risky/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

//Functions to test handlers/action.go

var capabilityId = "d27d7d83-0145-4caf-aab3-b250c23f07b5"
var actionId = "a32ec646-48e8-400b-b2aa-626c3085f1bf"

var controller = &handlers.PublicController{}

func TestGetActions(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}

	controller.DBManager = dbManager
	router := gin.Default()
	router.GET("/actions", controller.GetActions)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/actions?businessId="+businessId, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetAction(t *testing.T) {
	router := gin.Default()
	router.GET("/action", controller.GetAction)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/action?id="+actionId, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateAction(t *testing.T) {
	router := gin.Default()
	router.POST("/action", controller.CreateAction)
	w := httptest.NewRecorder()
	//Create a new handlers.ActionInput struct with prepared data
	capabilityId := uuid.MustParse(capabilityId)
	actionInput := handlers.ActionInput{
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
	router.PATCH("/action", controller.UpdateAction)
	w := httptest.NewRecorder()

	capabilityId := uuid.MustParse(capabilityId)
	actionInput := handlers.ActionInput{
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
	router.DELETE("/action", controller.DeleteAction)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/action", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// func TestInputToModel(t *testing.T) {
// 	actionInput := handlers.ActionInput{
// 		ID:   uuid.New(),
// 		Name: "Test",
// 	}
// 	_, err := inputToModel(actionInput)
// 	assert.Equal(t, nil, err)
// }

// func TestModelToOutput(t *testing.T) {
// 	actionModel := database.ActionModel{
// 		ID:   uuid.New(),
// 		Name: "Test",
// 	}
// 	_, err := modelToOutput(actionModel)
// 	assert.Equal(t, nil, err)
// }
