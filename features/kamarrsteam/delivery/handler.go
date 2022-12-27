package delivery

import (
	"net/http"

	"github.com/KamarRS-App/KamarRS-App/features/kamarrsteam"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"
	"github.com/labstack/echo/v4"
)

type KamarRsTeamDelivery struct {
	kamarrsteamService kamarrsteam.ServiceInterface
}

func New(service kamarrsteam.ServiceInterface, e *echo.Echo) {
	handler := &KamarRsTeamDelivery{
		kamarrsteamService: service,
	}

	e.POST("/kamarrsteam", handler.CreateTeam)
}

func (d *KamarRsTeamDelivery) CreateTeam(c echo.Context) error {
	input := KamarRsTeamRequest{}
	input.Peran = "super admin"
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	dataCore := requestToCore(input)
	err := d.kamarrsteamService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Create New Homestay"))
}
