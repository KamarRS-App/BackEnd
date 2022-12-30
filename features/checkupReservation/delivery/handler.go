package delivery

import (
	"net/http"

	checkupreservation "github.com/KamarRS-App/KamarRS-App/features/checkupReservation"
	"github.com/KamarRS-App/KamarRS-App/middlewares"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"
	"github.com/labstack/echo/v4"
)

type CheckupDelivery struct {
	checkupService checkupreservation.ServiceInterface
}

func New(Service checkupreservation.ServiceInterface, e *echo.Echo) {
	handler := &CheckupDelivery{
		checkupService: Service,
	}

	e.POST("/reservations", handler.Create, middlewares.JWTMiddleware())

	// e.GET("/users/:id", handler.GetById, middlewares.JWTMiddleware())

}
func (delivery *CheckupDelivery) Create(c echo.Context) error {

	role := middlewares.ExtractTokenTeamRole(c)
	if role != "" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Hanya bisa diakses user"))

	}
	userIdtoken := middlewares.ExtractTokenTeamId(c)

	inputCheckUp := CheckupReservationRequest{} //penangkapan data user reques dari entities user
	errbind := c.Bind(&inputCheckUp)

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr: "+errbind.Error()))
	}
	dataCore := inputCheckUp.reqToCore() //data mapping yang diminta create
	errResultCore := delivery.checkupService.Create(dataCore, userIdtoken)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr:"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("reservasi berhasil dibuat"))
}
