package delivery

import (
	"net/http"

	bedreservation "github.com/KamarRS-App/KamarRS-App/features/bedReservation"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"

	"github.com/labstack/echo/v4"
)

type BedReservationDelivery struct {
	BedReservationService bedreservation.ServiceInterface
}

func New(service bedreservation.ServiceInterface, e *echo.Echo) {
	handler := &BedReservationDelivery{
		BedReservationService: service,
	}

	e.POST("/registrations", handler.CreateRegistration)
	e.GET("/payments/:kodeDaftar", handler.GetPayment)
}

func (d *BedReservationDelivery) CreateRegistration(c echo.Context) error {
	input := BedReservationRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	dataCore := input.reqToCore()
	data, err := d.BedReservationService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}
	res := fromCore(data)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success create registrations", res))
}

func (d *BedReservationDelivery) GetPayment(c echo.Context) error {
	kodeDaftar := c.Param("kodeDaftar")
	data, err := d.BedReservationService.GetPayment(kodeDaftar)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}
	res := fromCore(data)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success create registrations", res))
}
