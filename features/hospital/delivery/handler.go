package delivery

import (
	"net/http"
	"strconv"

	"github.com/KamarRS-App/KamarRS-App/features/hospital"
	"github.com/KamarRS-App/KamarRS-App/middlewares"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"
	"github.com/labstack/echo/v4"
)

type HospitalDelivery struct {
	hospitalService hospital.ServiceInterface
}

func New(service hospital.ServiceInterface, e *echo.Echo) {
	handler := &HospitalDelivery{
		hospitalService: service,
	}
	e.POST("/hospitals", handler.Create, middlewares.JWTMiddleware())
	e.GET("/hospitals", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/hospitals/:id", handler.GetById, middlewares.JWTMiddleware())
	e.PUT("/hospitals/:id", handler.UpdateData, middlewares.JWTMiddleware())
	e.DELETE("/hospitals/:id", handler.Delete, middlewares.JWTMiddleware())
}

// Post
func (delivery *HospitalDelivery) Create(c echo.Context) error {
	hospitalInput := HospitalRequest{}
	errBind := c.Bind(&hospitalInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	fotoHospital, _ := helper.UploadFotoHospital(c, "foto")

	hospitalInput.Foto = fotoHospital

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := hospitalInput.ToCore()
	err := delivery.hospitalService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

// Get All Villa (Homepage)
func (delivery *HospitalDelivery) GetAll(c echo.Context) error {
	results, err := delivery.hospitalService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := FromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all hospitals", dataResponse))
}

// Get by ID
func (delivery *HospitalDelivery) GetById(c echo.Context) error {
	id, errBind := strconv.Atoi(c.Param("id"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	IdHospital, err := delivery.hospitalService.GetById(id)
	dataResponse := FromCore(IdHospital)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get hospitals", dataResponse))
}

// Update
func (delivery *HospitalDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	hospitalInput := HospitalRequest{}
	errBind := c.Bind(&hospitalInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := hospitalInput.ToCore()
	errUpt := delivery.hospitalService.Update(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data hospitals"))
}

// delete villa
func (delivery *HospitalDelivery) Delete(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	errDel := delivery.hospitalService.Delete(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete hospitals"+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}
