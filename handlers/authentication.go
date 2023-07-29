package handlers

import (
	"go-risky/types"
	"go-risky/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

// this file will contain the logic to authenticate a user and check if a user isAuthenticated

//This store needs to be migrated to a database of some kind
var store = sessions.NewCookieStore([]byte(""))

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

		err = createSession(userOutput.ID, userOutput.OrganizationID, userOutput.GroupID, context)
		if err != nil {
			log.Println(err)
			context.IndentedJSON(http.StatusNotFound, "Not Found")
			return
		}

		// token, err := util.GenerateJWT(userOutput)
		// if err != nil {
		// 	log.Printf("Error generating JWT: %s", err)

		// 	return
		// }
		// context.Header("Authorization", "Bearer "+token)
		context.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
	}
}

func (controller PublicController) Logout(context *gin.Context) {
	session, err := store.Get(context.Request, "session")
	if err != nil {
		log.Printf("Error getting session on logout: %s", err)
		return
	}
	session.Values["authenticated"] = false
	session.Values["organizationId"] = ""
	session.Values["groupId"] = ""
	session.Values["userId"] = ""
	session.Options.MaxAge = -1
	session.Save(context.Request, context.Writer)
	context.IndentedJSON(http.StatusOK, "Success")
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

func createSession(userId string, organizationId string, groupId string, context *gin.Context) (err error) {
	session, _ := store.Get(context.Request, "session")
	session.Values["authenticated"] = true
	session.Values["organizationId"] = organizationId
	session.Values["groupId"] = groupId
	session.Values["userId"] = userId

	err = session.Save(context.Request, context.Writer)
	if err != nil {
		return
	}
	return
}
