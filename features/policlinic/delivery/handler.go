package delivery

import (
	"net/http"
	"strconv"

	"github.com/KamarRS-App/KamarRS-App/features/policlinic"
	"github.com/KamarRS-App/KamarRS-App/middlewares"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"
	"github.com/labstack/echo/v4"
)

type PoliclinicDelivery struct {
	policlinicService policlinic.ServiceInterface
}

func New(service policlinic.ServiceInterface, e *echo.Echo) {
	handler := &PoliclinicDelivery{
		policlinicService: service,
	}
	e.POST("/policlinics", handler.Create, middlewares.JWTMiddleware())
	e.GET("/policlinics", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/policlinics/:id", handler.GetById, middlewares.JWTMiddleware())
	e.PUT("/policlinics/:id", handler.UpdateData, middlewares.JWTMiddleware())
	e.DELETE("/policlinics/:id", handler.Delete, middlewares.JWTMiddleware())
}

// Post
func (delivery *PoliclinicDelivery) Create(c echo.Context) error {
	policlinicInput := RequestPoliclinic{}
	errBind := c.Bind(&policlinicInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := policlinicInput.ToCore()
	err := delivery.policlinicService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

// Get All Villa (Homepage)
func (delivery *PoliclinicDelivery) GetAll(c echo.Context) error {
	results, err := delivery.policlinicService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := FromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all policlinics", dataResponse))
}

// Get by ID
func (delivery *PoliclinicDelivery) GetById(c echo.Context) error {
	id, errBind := strconv.Atoi(c.Param("id"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	Idpoliclinic, err := delivery.policlinicService.GetById(id)
	dataResponse := FromCore(Idpoliclinic)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get policlinics", dataResponse))
}

// Update
func (delivery *PoliclinicDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	policlinicInput := RequestPoliclinic{}
	errBind := c.Bind(&policlinicInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := policlinicInput.ToCore()
	errUpt := delivery.policlinicService.Update(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data policlinics"))
}

// delete villa
func (delivery *PoliclinicDelivery) Delete(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	errDel := delivery.policlinicService.Delete(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete policlinics"+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}
