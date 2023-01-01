package delivery

import (
	"net/http"
	"strconv"

	bedreservation "github.com/KamarRS-App/KamarRS-App/features/bedReservation"
	"github.com/KamarRS-App/KamarRS-App/middlewares"
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

	e.POST("/registrations", handler.CreateRegistration, middlewares.JWTMiddleware())
	e.GET("/payments/:kodeDaftar", handler.GetPayment, middlewares.JWTMiddleware())
	e.PUT("/payments/:kodeDaftar", handler.CreatePayment, middlewares.JWTMiddleware())
	e.POST("/midtrans", handler.UpdateMidtrans)
	e.GET("/hospitals/:hospitalId/registrations", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("registrations/:registration_id", handler.GetDetailRegistration, middlewares.JWTMiddleware())
}

func (d *BedReservationDelivery) CreateRegistration(c echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(c)
	if role != "" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Hanya bisa diakses user"))

	}
	userId := middlewares.ExtractTokenTeamId(c)

	input := BedReservationRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	dataCore := input.reqToCore()
	data, err := d.BedReservationService.Create(dataCore, uint(userId))
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

func (d *BedReservationDelivery) CreatePayment(c echo.Context) error {
	kodeDaftar := c.Param("kodeDaftar")
	input := BedReservationRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	input.KodeDaftar = kodeDaftar
	inputCore := input.reqToCore()
	data, err := d.BedReservationService.CreatePayment(inputCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}
	res := fromCore(data)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success create payment", res))
}

func (d *BedReservationDelivery) UpdateMidtrans(c echo.Context) error {
	var callback UpdateMidtrans
	errBind := c.Bind(&callback)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	callbackCore := callback.reqToCore()
	err := d.BedReservationService.PaymentNotif(callbackCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update payment from midtrans"))
}

func (d *BedReservationDelivery) GetAll(c echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(c)
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Hanya bisa diakses admin"))
	}
	pageQuery := c.QueryParam("page")
	page, _ := strconv.Atoi(pageQuery)
	limit := 10
	hospialId, errBind := strconv.Atoi(c.Param("hospitalId"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data"+errBind.Error()))
	}

	res, totalpage, err := d.BedReservationService.GetRegistrations(page, limit, hospialId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	data := fromCoreList(res)

	return c.JSON(http.StatusOK, helper.SuccessWithDataPaginationResponse("success read all bed reservations", data, totalpage))
}

func (d *BedReservationDelivery) GetDetailRegistration(c echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(c)
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Hanya bisa diakses admin"))
	}
	bedResId, _ := strconv.Atoi(c.Param("registration_id"))
	res, err := d.BedReservationService.GetById(uint(bedResId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	data := fromCore(res)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read beds registrations by ID", data))
}
