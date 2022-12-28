package delivery

import (
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
	errPass := bcrypt.CompareHashAndPassword([]byte(dataUser.Kata_Sandi), z)
	if errPass != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Incorrect Password "+errPass.Error()))
	}

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
		"email":   data.Email,
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success login", res))
}
