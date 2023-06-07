package asset

//Write unit tests for the functions declared in handlers/asset/asset.go

import (
	"bytes"
	"go-risky/riskyrouter"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

// Functions to test handlers/asset/asset.go with dummy data
var businessId = "568d1c21-c83f-4f4f-815b-9ee6490fe8f7"
var assetId = "465804b9-e5aa-49e1-b844-61ba3d928b84"

func TestGetAssets(t *testing.T) {
	router := riskyrouter.InitializeRouter()
	router.GET("/assets", getAssets)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/assets?businessId="+businessId, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetAsset(t *testing.T) {
	router := gin.Default()
	router.GET("/asset?id="+assetId, getAsset)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/asset", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateAsset(t *testing.T) {
	var createdAsset AssetOutput
	router := gin.Default()
	router.POST("/asset", createAsset)
	w := httptest.NewRecorder()
	assetInput, err := json.Marshal(AssetInput{Name: "test", BusinessID: uuid.MustParse(businessId)})
	assert.Equal(t, err, nil)
	req, _ := http.NewRequest("POST", "/asset", bytes.NewBuffer(assetInput))
	router.ServeHTTP(w, req)
	err = json.Unmarshal(w.Body.Bytes(), &createdAsset)
	assert.Equal(t, err, nil)

	assert.Equal(t, createdAsset.BusinessID.String(), businessId)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateAsset(t *testing.T) {
	router := gin.Default()
	router.PATCH("/asset", updateAsset)
	w := httptest.NewRecorder()
	assetInput, err := json.Marshal(AssetInput{ID: uuid.MustParse(assetId), Name: "test", BusinessID: uuid.MustParse(businessId)})
	assert.Equal(t, err, nil)
	req, _ := http.NewRequest("POST", "/asset", bytes.NewBuffer(assetInput))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteAsset(t *testing.T) {
	var createdAsset AssetOutput
	router := gin.Default()
	router.DELETE("/asset", deleteAsset)
	w := httptest.NewRecorder()

	assetInput, err := json.Marshal(AssetInput{Name: "test", BusinessID: uuid.MustParse(businessId)})
	assert.Equal(t, err, nil)
	req, _ := http.NewRequest("POST", "/asset", bytes.NewBuffer(assetInput))
	router.ServeHTTP(w, req)
	err = json.Unmarshal(w.Body.Bytes(), &createdAsset)
	assert.Equal(t, err, nil)

	req, _ = http.NewRequest("DELETE", "/asset?id="+createdAsset.ID.String(), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
