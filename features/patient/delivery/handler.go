package delivery

import (
	"net/http"
	"strconv"

	"github.com/KamarRS-App/features/patient"
	"github.com/KamarRS-App/middlewares"
	"github.com/KamarRS-App/utils/helper"
	"github.com/labstack/echo/v4"
)

type PatientDeliv struct {
	PatientService patient.ServiceInterface
}

func New(Service patient.ServiceInterface, e *echo.Echo) {
	handler := &PatientDeliv{
		PatientService: Service,
	}

	e.POST("/patients", handler.Create, middlewares.JWTMiddleware())
	e.GET("/patients/:id", handler.GetByPatientId, middlewares.JWTMiddleware())
	e.GET("/users/:id/patients", handler.GetByUserId, middlewares.JWTMiddleware())
	// e.GET("/users", handler.GetById, middlewares.JWTMiddleware()) //untuk sementara pake param karena login belum bisa
	// e.PUT("/users", handler.Update, middlewares.JWTMiddleware())
	// e.DELETE("/users", handler.DeleteById, middlewares.JWTMiddleware())
	// e.GET("/users/:id", handler.GetById, middlewares.JWTMiddleware())

}
func (delivery *PatientDeliv) Create(c echo.Context) error {

	userId := middlewares.ExtractTokenUserId(c)

	InputPatient := RequestPatient{}
	errbind := c.Bind(&InputPatient)

	fotoKtp, errKtp := helper.UploadFotoKTP(c, "foto_ktp")
	if errKtp != nil {
		return errKtp
	}
	fotoBpjs, errBpjs := helper.UploadFotoBPJS(c, "foto_bpjs")
	if errBpjs != nil {
		return errBpjs
	}
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
	userIdtoken, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return errConv
	}

	result, err := delivery.PatientService.GetByUserId(userIdtoken) //memanggil fungsi service yang ada di folder service//jika return nya 2 maka variable harus ada 2

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("erorr read data"))
	}
	var ResponData = ListpatientCoreTopatientRespon(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil menampilkan patient yang di daftarkan oleh user", ResponData))
}
