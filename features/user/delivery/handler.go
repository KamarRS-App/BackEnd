package delivery

import (
	"net/http"

	"github.com/KamarRS-App/features/user"
	"github.com/KamarRS-App/features/user/service"
	"github.com/KamarRS-App/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserDeliv struct {
	UserService user.ServiceInterface
}

func New(Service user.ServiceInterface, e *echo.Echo) {
	handler := &UserDeliv{
		UserService: Service,
	}

	e.POST("/users", handler.Create)
	// e.GET("/users", handler.GetAll, middlewares.JWTMiddleware())
	// e.PUT("/users/:id", handler.Update, middlewares.JWTMiddleware())
	// e.DELETE("/users/:id", handler.DeleteById, middlewares.JWTMiddleware())
	// e.GET("/users/:id", handler.GetById, middlewares.JWTMiddleware())

}
func (delivery *UserDeliv) Create(c echo.Context) error {

	// roletoken := middlewares.ExtractTokenUserRole(c)
	// log.Println("Role Token", roletoken)
	// if roletoken != "Admin" {
	// 	return c.JSON(http.StatusUnauthorized, helper.PesanGagalHelper("tidak bisa diakses khusus admin!!!"))
	// }

	Inputuser := RequestUser{} //penangkapan data user reques dari entities user
	errbind := c.Bind(&Inputuser)

	generatePass := service.Bcript(Inputuser.KataSandi)
	Inputuser.KataSandi = generatePass

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errbind.Error()))
	}
	dataCore := Inputuser.reqToCore() //data mapping yang diminta create
	errResultCore := delivery.UserService.Create(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("berhasil create user"))
}
