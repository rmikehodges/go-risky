package handlers

import (
	"go-risky/types"
	"go-risky/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//this file will contain the logic to authenticate a user and check if a user isAuthenticated

func (controller PublicController) Login(context *gin.Context) {

	userOutput := types.User{}
	authentication := types.Authentication{}
	err := context.ShouldBindJSON(&authentication)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	organization, err := controller.DBManager.GetOrganizationByUserEmail(authentication.Email)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	if organization.OAuthEnabled {
		log.Println("OAuth is enabled for this user")
		auth, err = types.NewAuthenticator(organization)
		if err != nil {
			log.Printf("Error creating authenticator: %s", err)
			context.IndentedJSON(http.StatusBadRequest, "Bad request")
		}
		randomToken, err := util.GenerateRandomToken()
		if err != nil {
			log.Printf("Error generating random token for oAuth2: %s", err)
			context.IndentedJSON(http.StatusBadRequest, "Bad request")
		}
		context.SetCookie("authCode", randomToken, 3600, "/", "localhost", false, true)

		context.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(randomToken))

	} else {
		userOutput, err = controller.DBManager.AuthenticateWithPassword(authentication.Email, authentication.Password)
		if err != nil {
			log.Println(err)
			context.IndentedJSON(http.StatusNotFound, "Not Found")
			return
		}
		if userOutput.ID == uuid.Nil.String() {
			log.Println("User not found")
			context.IndentedJSON(http.StatusNotFound, "Not Found")
			return
		}

		token, err := util.GenerateJWT(userOutput)
		if err != nil {
			log.Printf("Error generating JWT: %s", err)
			context.IndentedJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
		context.Header("Authorization", "Bearer "+token)
		context.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}
}

func (controller PublicController) ChangePassword(context *gin.Context) {

	authentication := types.Authentication{}
	err := context.ShouldBindJSON(&authentication)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(authentication.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Error Hashing Password")
	}

	err = controller.DBManager.ChangePassword(authentication.ID, string(bcryptPassword))
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}

func (controller PublicController) InitiatePasswordReset(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		log.Println("Parameter id not found")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	userId, err := uuid.Parse(id)
	if err != nil {
		log.Println("ID is not a uuid")
		context.JSON(http.StatusNotFound, "Not found")
		return
	}

	resetToken, err := util.GenerateRandomToken()

	if err != nil {
		log.Printf("Initiate Password Reset Token Error: %s", err)
		return
	}

	err = controller.DBManager.InitiatePasswordReset(userId.String(), resetToken)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	//TODO: Send email with reset token or display it to the server logs.

	context.IndentedJSON(http.StatusOK, resetToken)
}

func (controller PublicController) ResetPassword(context *gin.Context) {

	authentication := types.Authentication{}
	err := context.ShouldBindJSON(&authentication)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Bad request")
	}

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(authentication.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, "Error Hashing Password")
	}

	err = controller.DBManager.ResetPassword(authentication.ID, string(bcryptPassword), authentication.PasswordResetToken)
	if err != nil {
		log.Println(err)
		context.IndentedJSON(http.StatusNotFound, "Not Found")
		return
	}

	context.IndentedJSON(http.StatusOK, "Success")
}
