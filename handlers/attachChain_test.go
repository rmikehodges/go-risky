package handlers_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"go-risky/database"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/assert"
// 	"github.com/google/uuid"
// )

// //Functions to test handlers/attackChain.go

// var businessId = "568d1c21-c83f-4f4f-815b-9ee6490fe8f7"
// var threatId = "d27d7d83-0145-4caf-aab3-b250c23f07b5"
// var attackChainId = "a32ec646-48e8-400b-b2aa-626c3085f1bf"

// func TestGetAttackChains(t *testing.T) {
// 	router := gin.Default()
// 	router.GET("/attackChains", getAttackChains)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/attackChains?businessId="+businessId, nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)
// }

// func TestGetAttackChain(t *testing.T) {
// 	router := gin.Default()
// 	router.GET("/attackChain", getAttackChain)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/attackChain?id="+attackChainId, nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)
// }

// func TestCreateAttackChain(t *testing.T) {
// 	router := gin.Default()
// 	router.POST("/attackChain", createAttackChain)
// 	w := httptest.NewRecorder()
// 	//Create a new AttackChainInput struct with prepared data
// 	attackChainInput := AttackChainInput{
// 		Name:        "Test",
// 		Description: "test",
// 		BusinessID:  uuid.MustParse(businessId),
// 		ThreatID:    uuid.MustParse(threatId),
// 	}

// 	attackChainInputJson, err := json.Marshal(attackChainInput)

// 	if err != nil {
// 		t.Errorf("Error while marshalling attackChainInputJson: %v\n", err)
// 	}

// 	req, _ := http.NewRequest("POST", "http://localhost:8081/attackChain", bytes.NewBuffer(attackChainInputJson))
// 	req.Header.Set("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)
// }

// func TestUpdateAttackChain(t *testing.T) {
// 	router := gin.Default()
// 	router.PATCH("/attackChain", updateAttackChain)
// 	w := httptest.NewRecorder()

// 	attackChainInput := AttackChainInput{
// 		ID:          uuid.MustParse(attackChainId),
// 		Name:        "Test",
// 		Description: "test",
// 		BusinessID:  uuid.MustParse(businessId),
// 	}

// 	attackChainInputJson, err := json.Marshal(attackChainInput)

// 	if err != nil {
// 		t.Errorf("Error while marshalling attackChainInputJson: %v\n", err)
// 	}

// 	req, _ := http.NewRequest("PATCH", "/attackChain", bytes.NewBuffer(attackChainInputJson))
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)
// }

// func TestDeleteAttackChain(t *testing.T) {
// 	router := gin.Default()
// 	router.DELETE("/attackChain", deleteAttackChain)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/attackChain", nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)
// }

// func TestInputToModel(t *testing.T) {
// 	attackChainInput := AttackChainInput{
// 		ID:   uuid.New(),
// 		Name: "Test",
// 	}
// 	_, err := inputToModel(attackChainInput)
// 	assert.Equal(t, nil, err)
// }

// func TestModelToOutput(t *testing.T) {
// 	attackChainModel := database.AttackChainModel{
// 		ID:   uuid.New(),
// 		Name: "Test",
// 	}
// 	_, err := modelToOutput(attackChainModel)
// 	assert.Equal(t, nil, err)
// }
