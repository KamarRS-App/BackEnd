package delivery

import (
	"net/http"

	hospitalstaff "github.com/KamarRS-App/features/hospitalStaff"
	"github.com/KamarRS-App/middlewares"
	"github.com/KamarRS-App/utils/helper"

	"github.com/labstack/echo/v4"
)

type StaffDeliv struct {
	StaffService hospitalstaff.ServiceInterface
}

func New(Service hospitalstaff.ServiceInterface, e *echo.Echo) {
	handler := &StaffDeliv{
		StaffService: Service,
	}

	e.POST("/staffs", handler.Create, middlewares.JWTMiddleware())
	e.GET("/staffs", handler.GetStaff, middlewares.JWTMiddleware()) //untuk sementara pake param karena login belum bisa
	e.PUT("/staffs", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/staffs", handler.DeleteById, middlewares.JWTMiddleware())

}
func (delivery *StaffDeliv) Create(c echo.Context) error {

	Inputstaff := HospitalStaffRequest{} //penangkapan data user reques dari entities user
	errbind := c.Bind(&Inputstaff)

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr: "+errbind.Error()))
	}
	dataCore := Inputstaff.reqToCore() //data mapping yang diminta create
	errResultCore := delivery.StaffService.Create(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr:"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Akun Staff berhasil dibuat"))
}

func (delivery *StaffDeliv) Update(c echo.Context) error {

	userIdtoken := middlewares.ExtractTokenUserId(c)
	// log.Println("user_id_token", userIdtoken)
	Inputstaff := HospitalStaffRequest{}
	errBind := c.Bind(&Inputstaff) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := Inputstaff.reqToCore()

	err := delivery.StaffService.Update(userIdtoken, dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Gagal merubah data user"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Perubahan Data Berhasil"))
}

func (delivery *StaffDeliv) GetStaff(c echo.Context) error {
	userIdtoken := middlewares.ExtractTokenUserId(c)

	result, err := delivery.StaffService.GetStaff(userIdtoken)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = StaffCoreToStaffRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca  Staff", ResponData))
}

func (delivery *StaffDeliv) DeleteById(c echo.Context) error {

	userIdtoken := middlewares.ExtractTokenUserId(c)

	err := delivery.StaffService.DeleteById(userIdtoken) //memanggil fungsi service yang ada di folder service
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr Hapus data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil menghapus Staff"))
}
