package delivery

import "github.com/KamarRS-App/features/patient"

type RequestPatient struct {
	NoKk                  string `json:"no_kk" form:"no_kk"`
	Nik                   string `json:"nik" form:"nik"`
	NamaPasien            string `json:"nama_pasien" form:"nama_pasien"`
	JenisKelamin          string `json:"jenis_kelamin" form:"jenis_kelamin"`
	TanggalLahir          string `json:"tanggal_lahir" form:"tanggal_lahir"`
	Usia                  int    `json:"usia" form:"usia"`
	NamaWali              string `json:"nama_wali" form:"nama_wali"`
	EmailWali             string `json:"email_wali" form:"email_wali"`
	NoTelponWali          string `json:"no_telpon_wali" form:"no_telpon_wali"`
	AlamatKtp             string `json:"alamat_ktp" form:"alamat_ktp"`
	ProvinsiKtp           string `json:"provinsi_ktp" form:"provinsi_ktp"`
	KabupatenKotaKtp      string `json:"kabupaten_kota_ktp" form:"kabupaten_kota_ktp"`
	AlamatDomisili        string `json:"alamat_domisili" form:"alamat_domisili"`
	ProvinsiDomisili      string `json:"provinsi_domisili" form:"provinsi_domisili"`
	KabupatenKotaDomisili string `json:"kabupaten_kota_domisili" form:"kabupaten_kota_domisili"`
	NoBpjs                string `json:"no_bpjs" form:"no_bpjs"`
	KelasBpjs             string `json:"kelas_bpjs" form:"kelas_bpjs"`
	FotoKtp               string `json:"foto_ktp" form:"foto_ktp"`
	FotoBpjs              string `json:"foto_bpjs" form:"foto_bpjs"`
	UserID                uint   `json:"user_id" form:"user_id"`
	// BedReservation          BedReservation
	// CheckupReservation      CheckupReservation
}

func (req *RequestPatient) reqToCore() patient.CorePatient {
	return patient.CorePatient{
		NoKk:                  req.NoKk,
		Nik:                   req.Nik,
		NamaPasien:            req.NamaPasien,
		JenisKelamin:          req.JenisKelamin,
		TanggalLahir:          req.TanggalLahir,
		Usia:                  req.Usia,
		NamaWali:              req.NamaWali,
		EmailWali:             req.EmailWali,
		NoTelponWali:          req.NoTelponWali,
		AlamatKtp:             req.AlamatKtp,
		ProvinsiKtp:           req.ProvinsiKtp,
		KabupatenKotaKtp:      req.KabupatenKotaKtp,
		AlamatDomisili:        req.AlamatDomisili,
		ProvinsiDomisili:      req.ProvinsiDomisili,
		KabupatenKotaDomisili: req.KabupatenKotaDomisili,
		NoBpjs:                req.NoBpjs,
		KelasBpjs:             req.KelasBpjs,
		FotoKtp:               req.FotoKtp,
		FotoBpjs:              req.FotoBpjs,
		UserID:                req.UserID,
	}

}
