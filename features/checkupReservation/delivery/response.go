package delivery

import (
	checkupreservation "kamarRS/features/checkupReservation"
	"time"
)

type CheckupReservationResponse struct {
	ID        uint             `json:"id"`
	CreatedAt time.Time        `json:"createdat"`
	Patient   PatientResponse  `json:"patient"`
	Practice  PracticeResponse `json:"practice"`
}

type PatientResponse struct {
	ID                    uint
	NoKk                  string
	Nik                   string
	NamaPasien            string
	JenisKelamin          string
	TanggalLahir          string
	Usia                  int
	NamaWali              string
	EmailWali             string
	NoTelponWali          string
	AlamatKtp             string
	ProvinsiKtp           string
	KabupatenKotaKtp      string
	AlamatDomisili        string
	ProvinsiDomisili      string
	KabupatenKotaDomisili string
	NoBpjs                string
	KelasBpjs             string
	FotoKtp               string
	FotoBpjs              string
}

type PracticeResponse struct {
	ID             uint
	TanggalPraktik string
	KuotaHarian    int
	Status         string
	PoliclinicID   uint
}

// -----------------Checkup Reserve--------------------
func fromCore(dataCore checkupreservation.CheckupReservationCore) CheckupReservationResponse {
	return CheckupReservationResponse{
		ID:        dataCore.ID,
		CreatedAt: dataCore.CreatedAt,
		Patient: PatientResponse{
			ID:                    dataCore.Patient.ID,
			NoKk:                  dataCore.Patient.NoKk,
			Nik:                   dataCore.Patient.Nik,
			NamaPasien:            dataCore.Patient.NamaPasien,
			JenisKelamin:          dataCore.Patient.JenisKelamin,
			TanggalLahir:          dataCore.Patient.TanggalLahir,
			Usia:                  dataCore.Patient.Usia,
			NamaWali:              dataCore.Patient.NamaWali,
			EmailWali:             dataCore.Patient.EmailWali,
			NoTelponWali:          dataCore.Patient.NoTelponWali,
			AlamatKtp:             dataCore.Patient.AlamatKtp,
			ProvinsiKtp:           dataCore.Patient.ProvinsiKtp,
			KabupatenKotaKtp:      dataCore.Patient.KabupatenKotaKtp,
			AlamatDomisili:        dataCore.Patient.AlamatDomisili,
			ProvinsiDomisili:      dataCore.Patient.ProvinsiDomisili,
			KabupatenKotaDomisili: dataCore.Patient.KabupatenKotaDomisili,
			NoBpjs:                dataCore.Patient.NoBpjs,
			KelasBpjs:             dataCore.Patient.KelasBpjs,
			FotoKtp:               dataCore.Patient.FotoKtp,
			FotoBpjs:              dataCore.Patient.FotoBpjs,
		},
		Practice: PracticeResponse{
			ID:             dataCore.Practice.ID,
			TanggalPraktik: dataCore.Practice.TanggalPraktik,
			KuotaHarian:    dataCore.Practice.KuotaHarian,
			Status:         dataCore.Practice.Status,
			PoliclinicID:   dataCore.Practice.PoliclinicID,
		},
	}
}

// data dari core ke response
func fromCoreList(dataCore []checkupreservation.CheckupReservationCore) []CheckupReservationResponse {
	var dataResponse []CheckupReservationResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

//-----------------------------------------------------
