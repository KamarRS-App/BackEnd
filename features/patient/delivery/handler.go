package delivery

import (
	"net/http"

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
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Akun user berhasil dibuat"))
}
