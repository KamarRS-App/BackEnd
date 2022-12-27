package delivery

import (
	"net/http"
	"strconv"

	"github.com/KamarRS-App/KamarRS-App/features/doctor"
	middlewares "github.com/KamarRS-App/KamarRS-App/middlewares"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"
	"github.com/labstack/echo/v4"
)

type DoctorDelivery struct {
	doctorService doctor.ServiceInterface
}

func New(service doctor.ServiceInterface, e *echo.Echo) {
	handler := &DoctorDelivery{
		doctorService: service,
	}
	e.POST("/doctors", handler.Create, middlewares.JWTMiddleware())
	e.GET("/doctors", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/doctors/:id", handler.GetById, middlewares.JWTMiddleware())
	e.PUT("/doctors/:id", handler.UpdateData, middlewares.JWTMiddleware())
	e.DELETE("/doctors/:id", handler.Delete, middlewares.JWTMiddleware())
}

// Post
func (delivery *DoctorDelivery) Create(c echo.Context) error {
	doctorInput := DoctorRequest{}
	errBind := c.Bind(&doctorInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := doctorInput.ToCore()
	err := delivery.doctorService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

// Get All Villa (Homepage)
func (delivery *DoctorDelivery) GetAll(c echo.Context) error {
	results, err := delivery.doctorService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := FromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all doctors", dataResponse))
}

// Get by ID
func (delivery *DoctorDelivery) GetById(c echo.Context) error {
	id, errBind := strconv.Atoi(c.Param("id"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	Iddoctor, err := delivery.doctorService.GetById(id)
	dataResponse := FromCore(Iddoctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get doctors", dataResponse))
}

// Update
func (delivery *DoctorDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	doctorInput := DoctorRequest{}
	errBind := c.Bind(&doctorInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := doctorInput.ToCore()
	errUpt := delivery.doctorService.Update(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data doctors"))
}

// delete villa
func (delivery *DoctorDelivery) Delete(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	errDel := delivery.doctorService.Delete(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete doctors"+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}
