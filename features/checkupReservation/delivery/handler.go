package delivery

import (
	"net/http"
	"strconv"

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
	e.GET("/practices/:id/reservations", handler.GetByPracticesId, middlewares.JWTMiddleware())
	e.GET("/reservations/:id", handler.GetByreservationId, middlewares.JWTMiddleware())

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

func (delivery *CheckupDelivery) GetByPracticesId(c echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(c)
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Hanya bisa diakses admin"))

	}
	page := c.QueryParam("page") // input page
	pagination, _ := strconv.Atoi(page)
	limit := 10 // set default limit buat pagination
	practiceId, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return errConv
	}

	data, totalpage, err := delivery.checkupService.GetByPracticesId(pagination, limit, practiceId) //memanggil fungsi service yang ada di folder service//jika return nya 2 maka variable harus ada 2

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = fromCoreList(data)
	return c.JSON(http.StatusOK, helper.SuccessWithDataPaginationResponse("berhasil menampilkan reservasi berdasarkan id praktek", ResponData, totalpage))
}

func (delivery *CheckupDelivery) GetByreservationId(c echo.Context) error {

	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return errConv
	}

	result, err := delivery.checkupService.GetByreservationId(id) //memanggil fungsi service yang ada di folder service//jika return nya 2 maka variable harus ada 2

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = fromCore1(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca detail reservasi", ResponData))
}
