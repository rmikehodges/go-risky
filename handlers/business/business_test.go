package business

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

//Functions to test handlers/business.go

var businessId = "568d1c21-c83f-4f4f-815b-9ee6490fe8f7"

func TestGetBusinesses(t *testing.T) {
	router := gin.Default()
	router.GET("/businesses", getBusinesses)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/businesses?businessId="+businessId, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetBusiness(t *testing.T) {

	router := gin.Default()
	router.GET("/business", getBusiness)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/business?id="+businessId, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateBusiness(t *testing.T) {
	router := gin.Default()
	router.POST("/business", createBusiness)
	w := httptest.NewRecorder()
	//Create a new BusinessInput struct with prepared data
	businessInput := BusinessInput{
		Name:    "Test",
		Revenue: 0.0,
	}

	businessInputJson, err := json.Marshal(businessInput)

	if err != nil {
		t.Errorf("Error while marshalling businessInputJson: %v\n", err)
	}

	req, _ := http.NewRequest("POST", "http://localhost:8081/business", bytes.NewBuffer(businessInputJson))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUpdateBusiness(t *testing.T) {
	router := gin.Default()
	router.PATCH("/business", updateBusiness)
	w := httptest.NewRecorder()

	businessInput := BusinessInput{
		ID:      uuid.MustParse(businessId),
		Name:    "Test",
		Revenue: 0.0,
	}

	businessInputJson, err := json.Marshal(businessInput)

	if err != nil {
		t.Errorf("Error while marshalling businessInputJson: %v\n", err)
	}

	req, _ := http.NewRequest("PATCH", "/business", bytes.NewBuffer(businessInputJson))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteBusiness(t *testing.T) {
	router := gin.Default()
	router.DELETE("/business", deleteBusiness)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/business", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestInputToModel(t *testing.T) {
	businessInput := BusinessInput{
		ID:      uuid.New(),
		Name:    "Test",
		Revenue: 0.0,
	}
	_, err := inputToModel(businessInput)
	assert.Equal(t, nil, err)
}

func TestModelToOutput(t *testing.T) {
	businessModel := database.BusinessModel{
		ID:      uuid.New(),
		Name:    "Test",
		Revenue: 0.0,
	}
	_, err := modelToOutput(businessModel)
	assert.Equal(t, nil, err)
}
