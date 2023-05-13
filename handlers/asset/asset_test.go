package asset

//Write unit tests for the functions declared in handlers/asset/asset.go

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

// Functions to test handlers/asset/asset.go with dummy data
var businessId = "568d1c21-c83f-4f4f-815b-9ee6490fe8f7"

func TestGetAssets(t *testing.T) {
	router := gin.Default()
	router.GET("/assets", getAssets)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/assets?businessId="+businessId, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetAsset(t *testing.T) {
	router := gin.Default()
	router.GET("/asset", getAsset)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/asset", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateAsset(t *testing.T) {
	router := gin.Default()
	router.POST("/asset", createAsset)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/asset", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUpdateAsset(t *testing.T) {
	router := gin.Default()
	router.PATCH("/asset", updateAsset)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/asset", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteAsset(t *testing.T) {
	router := gin.Default()
	router.DELETE("/asset", deleteAsset)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/asset", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
