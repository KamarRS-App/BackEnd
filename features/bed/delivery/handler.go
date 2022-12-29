package delivery

import (
	"net/http"
	"strconv"

	"github.com/KamarRS-App/KamarRS-App/features/bed"
	"github.com/KamarRS-App/KamarRS-App/middlewares"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"
	"github.com/labstack/echo/v4"
)

type BedDelivery struct {
	bedService bed.ServiceInterface
}

func New(service bed.ServiceInterface, e *echo.Echo) {
	handler := &BedDelivery{
		bedService: service,
	}
	e.POST("/beds", handler.Create, middlewares.JWTMiddleware())
	e.GET("/hospitals/:id/beds", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/beds/:id", handler.GetById, middlewares.JWTMiddleware())
	e.PUT("/beds/:id", handler.UpdateData, middlewares.JWTMiddleware())
	e.DELETE("/beds/:id", handler.Delete, middlewares.JWTMiddleware())
}

// Post
func (delivery *BedDelivery) Create(c echo.Context) error {
	bedInput := BedRequest{}
	errBind := c.Bind(&bedInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := bedInput.ToCore()
	err := delivery.bedService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

// Get All by Hospital_ID
func (delivery *BedDelivery) GetAll(c echo.Context) error {
	page := c.QueryParam("page") // input page
	pagination, _ := strconv.Atoi(page)
	limit := 10 // set default limit buat pagination
	hospitalId, errBind := strconv.Atoi(c.Param("id"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data"+errBind.Error()))
	}

	results, totalpage, err := delivery.bedService.GetAll(pagination, limit, hospitalId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := FromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataPaginationResponse("success read all beds", dataResponse, totalpage))
}

// Get by ID
func (delivery *BedDelivery) GetById(c echo.Context) error {
	id, errBind := strconv.Atoi(c.Param("id"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	IdBed, err := delivery.bedService.GetById(id)
	dataResponse := FromCore(IdBed)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get beds", dataResponse))
}

// Update
func (delivery *BedDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	bedInput := BedRequest{}
	errBind := c.Bind(&bedInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := bedInput.ToCore()
	errUpt := delivery.bedService.Update(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data beds"))
}

// Delete
func (delivery *BedDelivery) Delete(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	errDel := delivery.bedService.Delete(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete beds"+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}
