package handlers

import (
	"go-risky/types"
	"go-risky/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var auth *types.Authenticator

func (controller PublicController) Callback(context *gin.Context) {
	authCode, err := context.Cookie("authCode")
	if err != nil {
		context.JSON(http.StatusUnauthorized, "Failed to get auth code.")
		return
	}

	token, err := auth.Exchange(context.Request.Context(), authCode)
	if err != nil {
		context.JSON(http.StatusUnauthorized, "Failed to exchange an authorization code for a token.")
		return
	}

	idToken, err := auth.VerifyIDToken(context.Request.Context(), token)
	if err != nil {
		context.JSON(http.StatusInternalServerError, "Failed to verify ID Token.")
		return
	}

	context.SetCookie("authCode", "", -1, "/", "localhost", false, true)

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userOutput, err := controller.DBManager.GetUserByEmail(profile["email"].(string))
	if err != nil {
		context.JSON(http.StatusInternalServerError, "Failed to get user by email.")
		return
	}

	if userOutput.ID == uuid.Nil.String() {
		log.Println("User not found")
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	jwtToken, err := util.GenerateJWT(userOutput)
	if err != nil {
		log.Printf("Error generating JWT: %s", err)
		context.IndentedJSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	context.Header("Authorization", "Bearer "+jwtToken)
	context.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
}
