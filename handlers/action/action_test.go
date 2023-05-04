package action

import (
	"go-risky/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/google/uuid"
)

//Functions to test handlers/action.go

func TestGetActions(t *testing.T) {
	router := gin.Default()
	router.GET("/actions", getActions)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/actions", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetAction(t *testing.T) {
	router := gin.Default()
	router.GET("/action", getAction)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/action", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateAction(t *testing.T) {
	router := gin.Default()
	router.POST("/action", createAction)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/action", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUpdateAction(t *testing.T) {
	router := gin.Default()
	router.PATCH("/action", updateAction)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/action", nil)
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

func TestActionRoutes(t *testing.T) {
	router := gin.Default()
	ActionRoutes(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/actions", nil)
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
