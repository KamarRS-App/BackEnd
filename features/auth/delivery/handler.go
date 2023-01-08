package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/KamarRS-App/KamarRS-App/features/auth"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"
	"github.com/labstack/echo/v4"
)

type AuthDelivery struct {
	authServices auth.ServiceInterface
}

func New(service auth.ServiceInterface, e *echo.Echo) {
	handler := &AuthDelivery{
		authServices: service,
	}

	e.POST("/login/users", handler.login)
	e.POST("/login/kamarrsteams", handler.loginTeam)
	e.POST("/login/staffs", handler.loginStaff)
	e.GET("/auth/google/login", oauthGoogleLogin)
	e.GET("/auth/google/callback", handler.oauthGoogleCallback)

}

func (delivery *AuthDelivery) login(c echo.Context) error {
	authInput := AuthRequest{}
	errBind := c.Bind(&authInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	token, dataUser, err := delivery.authServices.Login(authInput.Email, authInput.KataSandi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed login"))
	}

	z := []byte(authInput.KataSandi)
	errPass := bcrypt.CompareHashAndPassword([]byte(dataUser.KataSandi), z)
	if errPass != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Incorrect Password "+errPass.Error()))
	}
	// x := middlewares.ExtractTokenUserId(c)
	data := map[string]interface{}{
		"user_id": dataUser.ID,
		"token":   token,
		"name":    dataUser.Nama,
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success login", data))

}
func (d *AuthDelivery) loginTeam(c echo.Context) error {
	authInput := AuthRequestTeam{}
	errBind := c.Bind(&authInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	token, data, err := d.authServices.LoginTeam(authInput.Email, authInput.KataSandi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed login"))
	}
	passCheck := helper.CheckPasswordHash(authInput.KataSandi, data.KataSandi)
	if !passCheck {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Incorrect Password"))
	}

	res := map[string]interface{}{
		"team_id": data.ID,
		"token":   token,
		"peran":   data.Peran,
		"email":   data.Email,
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success login", res))
}

func (delivery *AuthDelivery) loginStaff(c echo.Context) error {
	authInput := AuthRequest{}

	errBind := c.Bind(&authInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	token, datastaff, err := delivery.authServices.LoginStaff(authInput.Email, authInput.KataSandi)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed login"))
	}

	z := []byte(authInput.KataSandi)
	errPass := bcrypt.CompareHashAndPassword([]byte(datastaff.KataSandi), z)
	if errPass != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Incorrect Password "+errPass.Error()))
	}

	data := map[string]interface{}{
		"staff_id": datastaff.ID,
		"token":    token,
		"name":     datastaff.Nama,
		// "peran":    x,
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success login", data))

}

func oauthGoogleLogin(c echo.Context) error {
	// var w http.ResponseWriter
	// var r *http.Request
	// Create oauthState cookie
	oauthState := helper.GenerateStateOauthCookie(c)

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/
	u := helper.AuthConfig().AuthCodeURL(oauthState)
	c.Redirect(http.StatusTemporaryRedirect, u)
	return c.JSON(http.StatusOK, "succes")
}

func (delivery *AuthDelivery) oauthGoogleCallback(c echo.Context) error {
	// Read oauthState from Cookie
	oauthState, _ := c.Cookie("oauthstate")

	if c.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return errors.New("eorror callback")
	}

	data, err := helper.GetUserDataFromGoogle(c.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return err
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	// fmt.Fprintf(c, "UserInfo: %s\n", data)
	var google auth.Oauth
	errUnmarshal := json.Unmarshal(data, &google)
	if errUnmarshal != nil {
		log.Fatal("error unmarshal")
	}

	token, dataUser, err := delivery.authServices.LoginOauth(google)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed login"))
	}

	// data1 := map[string]interface{}{
	// 	"user_id": dataUser.ID,
	// 	"token":   token,
	// 	"name":    dataUser.Nama,
	// }
	// fmt.Sprintf("token=%s&nama=%s&userid=%d", token, dataUser.Nama, dataUser.ID)
	return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("https://rawatinap.netlify.app/login/auth/google?token=%s&nama=%s&userid=%d", token, dataUser.Nama, dataUser.ID))

}
