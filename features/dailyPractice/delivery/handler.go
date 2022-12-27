package delivery

import (
	"net/http"
	"strconv"

	"github.com/KamarRS-App/features/dailypractice"
	"github.com/KamarRS-App/middlewares"
	"github.com/KamarRS-App/utils/helper"
	"github.com/labstack/echo/v4"
)

type PracticeDelivery struct {
	practiceService dailypractice.ServiceInterface
}

func New(service dailypractice.ServiceInterface, e *echo.Echo) {
	handler := &PracticeDelivery{
		practiceService: service,
	}
	e.POST("/practices", handler.Create, middlewares.JWTMiddleware())
	e.GET("/practices", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/practices/:id", handler.GetById, middlewares.JWTMiddleware())
	e.PUT("/practices/:id", handler.UpdateData, middlewares.JWTMiddleware())
}

// Post
func (delivery *PracticeDelivery) Create(c echo.Context) error {
	practiceInput := PracticeRequest{}
	errBind := c.Bind(&practiceInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := practiceInput.ToCore()
	err := delivery.practiceService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

// Get All
func (delivery *PracticeDelivery) GetAll(c echo.Context) error {
	results, err := delivery.practiceService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := FromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all policlinics", dataResponse))
}

// Get by ID
func (delivery *PracticeDelivery) GetById(c echo.Context) error {
	id, errBind := strconv.Atoi(c.Param("id"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	Idpractice, err := delivery.practiceService.GetById(id)
	dataResponse := FromCore(Idpractice)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get practices", dataResponse))
}

// Update
func (delivery *PracticeDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	practiceInput := PracticeRequest{}
	errBind := c.Bind(&practiceInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := practiceInput.ToCore()
	errUpt := delivery.practiceService.Update(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data practices"))
}
