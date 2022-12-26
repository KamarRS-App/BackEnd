package delivery

import (
	"net/http"
	"strconv"

	"github.com/KamarRS-App/features/user"
	"github.com/KamarRS-App/features/user/service"
	"github.com/KamarRS-App/middlewares"
	"github.com/KamarRS-App/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserDeliv struct {
	UserService user.ServiceInterface
}

func New(Service user.ServiceInterface, e *echo.Echo) {
	handler := &UserDeliv{
		UserService: Service,
	}

	e.POST("/users", handler.Create)
	e.GET("/users/:id", handler.GetById) //untuk sementara pake param karena login belum bisa
	e.PUT("/users", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", handler.DeleteById)
	// e.GET("/users/:id", handler.GetById, middlewares.JWTMiddleware())

}
func (delivery *UserDeliv) Create(c echo.Context) error {

	// roletoken := middlewares.ExtractTokenUserRole(c)
	// log.Println("Role Token", roletoken)
	// if roletoken != "Admin" {
	// 	return c.JSON(http.StatusUnauthorized, helper.PesanGagalHelper("tidak bisa diakses khusus admin!!!"))
	// }

	Inputuser := RequestUser{} //penangkapan data user reques dari entities user
	errbind := c.Bind(&Inputuser)

	generatePass := service.Bcript(Inputuser.KataSandi)
	Inputuser.KataSandi = generatePass

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errbind.Error()))
	}
	dataCore := Inputuser.reqToCore() //data mapping yang diminta create
	errResultCore := delivery.UserService.Create(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Akun user berhasil dibuat"))
}

func (delivery *UserDeliv) Update(c echo.Context) error {

	///////////bisa digunakan ketika Login sudah selesai/////////////////////////////////////////////////////

	userIdtoken := middlewares.ExtractTokenUserId(c)
	// log.Println("user_id_token", userIdtoken)
	userInput := RequestUser{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := userInput.reqToCore()

	err := delivery.UserService.Update(userIdtoken, dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Gagal merubah data user"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Perubahan Data Berhasil"))
}

func (delivery *UserDeliv) GetById(c echo.Context) error {

	///////////bisa digunakan ketika Login sudah selesai/////////////////////////////////////////////////////

	// userIdtoken := middlewares.ExtractTokenUserId(c)

	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//sementara pake param dlu

	id, _ := strconv.Atoi(c.Param("id"))

	result, err := delivery.UserService.GetById(id) //memanggil fungsi service yang ada di folder service//jika return nya 2 maka variable harus ada 2

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = UserCoreToUserRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca  user", ResponData))
}

func (delivery *UserDeliv) DeleteById(c echo.Context) error {
	///////////bisa digunakan ketika Login sudah selesai/////////////////////////////////////////////////////

	// userIdtoken := middlewares.ExtractTokenUserId(c)

	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	//sementara pake param dlu

	id, _ := strconv.Atoi(c.Param("id"))

	///////////////////////////////////////////////////////////////////////////////////////////////////////////
	err := delivery.UserService.DeleteById(id) //memanggil fungsi service yang ada di folder service
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr Hapus data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil menghapus user"))
}
