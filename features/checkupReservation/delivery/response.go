package delivery

import (
	"time"

	checkupreservation "github.com/KamarRS-App/KamarRS-App/features/checkupReservation"
)

type CheckupReservationResponse struct {
	ID         uint             `json:"id"`
	PatientID  uint             `json:"patient_id"`
	NamaDokter string           `json:"nama_dokter"`
	NoAntrian  string           `json:"no_antrian"`
	CreatedAt  time.Time        `json:"created_at"`
	Practice   PracticeResponse `json:"practice"`
	Patient    PatientResponse  `json:"patient"`
}

type PatientResponse struct {
	ID           uint   `json:"id"`
	NoKk         string `json:"no_kk"`
	JenisKelamin string `json:"jenis_kelamin"`
	Usia         int    `json:"usia"`
	Nik          string `json:"nik"`
	NamaPasien   string `json:"nama_pasien"`
}

type PracticeResponse struct {
	ID             uint   `json:"id"`
	TanggalPraktik string `json:"tanggal_praktik"`
	PoliclinicID   uint   `json:"policlinic_id"`
}

// -----------------Checkup Reserve--------------------
func fromCore(dataCore checkupreservation.CheckupReservationCore) CheckupReservationResponse {
	return CheckupReservationResponse{
		ID:         dataCore.ID,
		PatientID:  dataCore.PatientID,
		CreatedAt:  dataCore.CreatedAt,
		NamaDokter: dataCore.NamaDokter,
		NoAntrian:  dataCore.NoAntrian,
		Patient: PatientResponse{
			ID:           dataCore.Patient.ID,
			NoKk:         dataCore.Patient.NoKk,
			Nik:          dataCore.Patient.Nik,
			NamaPasien:   dataCore.Patient.NamaPasien,
			JenisKelamin: dataCore.Patient.JenisKelamin,
			Usia:         dataCore.Patient.Usia,
		},
		Practice: PracticeResponse{
			ID:             dataCore.Practice.ID,
			TanggalPraktik: dataCore.Practice.TanggalPraktik,
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

type CheckupReservationResponse1 struct {
	ID         uint              `json:"id"`
	PatientID  uint              `json:"patient_id"`
	NamaDokter string            `json:"nama_dokter"`
	NoAntrian  string            `json:"no_antrian"`
	CreatedAt  time.Time         `json:"created_at"`
	Practice   PracticeResponse1 `json:"practice"`
	Patient    PatientResponse1  `json:"patient"`
}

type PatientResponse1 struct {
	ID                    uint   `json:"id"`
	NoKk                  string `json:"no_kk"`
	Nik                   string `json:"nik"`
	NamaPasien            string `json:"nama_pasien"`
	JenisKelamin          string `json:"jenis_kelamin"`
	TanggalLahir          string `json:"tanggal_lahir"`
	Usia                  int    `json:"usia"`
	NamaWali              string `json:"nama_wali"`
	EmailWali             string `json:"email_wali"`
	NoTelponWali          string `json:"no_telpon_wali"`
	AlamatKtp             string `json:"alamat_ktp"`
	ProvinsiKtp           string `json:"provinsi_ktp"`
	KabupatenKotaKtp      string `json:"kabupaten_kota_ktp"`
	AlamatDomisili        string `json:"alamat_domisili"`
	ProvinsiDomisili      string `json:"provinsi_domisili"`
	KabupatenKotaDomisili string `json:"kabupaten_kota_domisili"`
	NoBpjs                string `json:"no_bpjs"`
	KelasBpjs             string `json:"kelas_bpjs"`
	FotoKtp               string `json:"foto_ktp"`
	FotoBpjs              string `json:"foto_bpjs"`
}

type PracticeResponse1 struct {
	ID             uint   `json:"id"`
	TanggalPraktik string `json:"tanggal_praktik"`
	PoliclinicID   uint   `json:"policlinic_id"`
}

// -----------------Checkup Reserve--------------------
func fromCore1(dataCore checkupreservation.CheckupReservationCore) CheckupReservationResponse1 {
	return CheckupReservationResponse1{
		ID:         dataCore.ID,
		PatientID:  dataCore.PatientID,
		NamaDokter: dataCore.NamaDokter,
		NoAntrian:  dataCore.NoAntrian,
		CreatedAt:  dataCore.CreatedAt,
		Patient: PatientResponse1{
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
		Practice: PracticeResponse1{
			ID:             dataCore.Practice.ID,
			TanggalPraktik: dataCore.Practice.TanggalPraktik,
			PoliclinicID:   dataCore.Practice.PoliclinicID,
		},
	}
}

// data dari core ke response
// func fromCoreList1(dataCore []checkupreservation.CheckupReservationCore) []CheckupReservationResponse1 {
// 	var dataResponse []CheckupReservationResponse1
// 	for _, v := range dataCore {
// 		dataResponse = append(dataResponse, fromCore1(v))
// 	}
// 	return dataResponse
// }
