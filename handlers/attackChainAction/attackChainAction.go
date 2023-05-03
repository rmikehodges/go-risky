package attackChainAction

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type attackChainAction struct {
	AttackChainID uuid.UUID `json:"attackChainid"`
	ActionId      uuid.UUID `json:"actionId"`
	BusinessId    uuid.UUID `json:"businessId"`
	Position      int       `json:"position"`
	CreatedAt     time.Time `json:"createdAt"`
}

var attackChainActions = []attackChainAction{}

func getAttackChainActions(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, attackChainActions)
}

func getAttackChainAction(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, attackChainActions)
}

func deleteAttackChainAction(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, attackChainActions)
}

func createAttackChainAction(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, attackChainActions)
}

func updateAttackChainAction(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, attackChainActions)
}

func AttackChainActionRoutes(router *gin.Engine) {
	router.GET("/attackChainActions", getAttackChainActions)
	router.GET("/attackChainAction/:id", getAttackChainAction)
	router.DELETE("/attackChainAction/:id", deleteAttackChainAction)
	router.PATCH("/attackChainAction/:id", updateAttackChainAction)
	router.POST("/attackChainActions", createAttackChainAction)
}
