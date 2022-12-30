package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KamarRS-App/KamarRS-App/features/patient"
	middlewares "github.com/KamarRS-App/KamarRS-App/middlewares"
	"github.com/KamarRS-App/KamarRS-App/utils/helper"
	"github.com/labstack/echo/v4"
)

type PatientDeliv struct {
	PatientService patient.ServiceInterface
}

func New(Service patient.ServiceInterface, e *echo.Echo) {
	handler := &PatientDeliv{
		PatientService: Service,
	}

	e.GET("/patients", handler.GetAllPatient, middlewares.JWTMiddleware())
	e.POST("/patients", handler.Create, middlewares.JWTMiddleware())
	e.GET("/patients/:id", handler.GetByPatientId, middlewares.JWTMiddleware())
	e.PUT("/patients/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/patients/:id", handler.DeleteById, middlewares.JWTMiddleware())
	e.GET("/users/:id/patients", handler.GetByUserId, middlewares.JWTMiddleware())

}
func (delivery *PatientDeliv) Create(c echo.Context) error {

	userId := middlewares.ExtractTokenTeamId(c)
	role := middlewares.ExtractTokenTeamRole(c)
	if role != "" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("fitur hanya untuk user"))
	}

	InputPatient := RequestPatient{}
	errbind := c.Bind(&InputPatient)

	fotoKtp, _ := helper.UploadFotoKTP(c, "foto_ktp")

	fotoBpjs, _ := helper.UploadFotoBPJS(c, "foto_bpjs")

	InputPatient.UserID = uint(userId)
	InputPatient.FotoKtp = fotoKtp
	InputPatient.FotoBpjs = fotoBpjs

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr: "+errbind.Error()))
	}
	dataCore := InputPatient.reqToCore() //data mapping yang diminta create
	errResultCore := delivery.PatientService.Create(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr:"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Berhasil mendaftarkan Pasien"))
}

func (delivery *PatientDeliv) GetByPatientId(c echo.Context) error {

	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return errConv
	}

	result, err := delivery.PatientService.GetByPatientId(id) //memanggil fungsi service yang ada di folder service//jika return nya 2 maka variable harus ada 2

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = PatientCoreToPatientRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil membaca data pasien", ResponData))
}

func (delivery *PatientDeliv) GetByUserId(c echo.Context) error {
	page := c.QueryParam("page") // input page
	pagination, _ := strconv.Atoi(page)
	limit := 10 // set default limit buat pagination
	userId, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return errConv
	}

	data, totalpage, err := delivery.PatientService.GetByUserId(pagination, limit, userId) //memanggil fungsi service yang ada di folder service//jika return nya 2 maka variable harus ada 2

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = ListpatientCoreTopatientRespon(data)
	return c.JSON(http.StatusOK, helper.SuccessWithDataPaginationResponse("berhasil menampilkan patient yang di daftarkan oleh user", ResponData, totalpage))
}

func (delivery *PatientDeliv) GetAllPatient(c echo.Context) error {

	role := middlewares.ExtractTokenTeamRole(c)
	fmt.Println(role)
	if role != "admin" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Maaf anda tidak punya akses ke data ini"))
	}

	result, err := delivery.PatientService.GetAllPatient() //memanggil fungsi service yang ada di folder service//jika return nya 2 maka variable harus ada 2

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = ListpatientCoreTopatientRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil menampilkan semua pasien", ResponData))
}

func (delivery *PatientDeliv) Update(c echo.Context) error {

	userId := middlewares.ExtractTokenTeamId(c)

	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return errConv
	}

	patientInput := RequestPatient{}
	errBind := c.Bind(&patientInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	fotoKtp, _ := helper.UploadFotoKTP(c, "foto_ktp")

	fotoBpjs, _ := helper.UploadFotoBPJS(c, "foto_bpjs")

	if fotoKtp != "" {
		patientInput.FotoKtp = fotoKtp

	}
	if fotoBpjs != "" {
		patientInput.FotoBpjs = fotoBpjs

	}

	dataCore := patientInput.reqToCore()

	err := delivery.PatientService.Update(id, userId, dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Gagal merubah data user"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Perubahan Data Berhasil"))
}

func (delivery *PatientDeliv) DeleteById(c echo.Context) error {

	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return errConv
	}

	err := delivery.PatientService.DeleteById(id) //memanggil fungsi service yang ada di folder service
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr Hapus data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil menghapus Pasien"))
}
