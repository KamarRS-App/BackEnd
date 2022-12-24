package repository

import (
	"kamarRS/features/patient"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	No_Kk                   string
	Nik                     string
	Nama_Pasien             string
	Jenis_Kelamin           string
	Tanggal_Lahir           string
	Usia                    int
	Nama_Wali               string
	Email_Wali              string
	No_Telpon_Wali          string
	Alamat_Ktp              string
	Provinsi_Ktp            string
	Kabupaten_Kota_Ktp      string
	Alamat_Domisili         string
	Provinsi_Domisili       string
	Kabupaten_Kota_Domisili string
	No_Bpjs                 string
	Kelas_Bpjs              string
	Foto_Ktp                string
	Foto_Bpjs               string
	UserID                  uint
	BedReservation          BedReservation
	CheckupReservation      CheckupReservation
}

type User struct {
	gorm.Model
	Nama       string
	Email      string
	No_kk      string
	Nik        string
	Kata_Sandi string
	No_Telpon  string
	Patients   []Patient
}

type BedReservation struct {
	gorm.Model
	Hospital_Id       uint
	Status_Pasien     string
	Biaya_Registrasi  int
	Order_Id          string
	Link_Pembayaran   string
	Status_Pembayaran string
	PatientID         uint
}

type CheckupReservation struct {
	gorm.Model
	PatientID  uint
	PracticeID uint
}

func FromPatientCore(dataCore patient.CorePatient) Patient { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	patientGorm := Patient{
		No_Kk:                   dataCore.NoKk,
		Nik:                     dataCore.Nik,
		Nama_Pasien:             dataCore.NamaPasien,
		Jenis_Kelamin:           dataCore.JenisKelamin,
		Tanggal_Lahir:           dataCore.TanggalLahir,
		Usia:                    dataCore.Usia,
		Nama_Wali:               dataCore.NamaWali,
		Email_Wali:              dataCore.EmailWali,
		No_Telpon_Wali:          dataCore.NoTelponWali,
		Alamat_Ktp:              dataCore.AlamatKtp,
		Provinsi_Ktp:            dataCore.ProvinsiKtp,
		Kabupaten_Kota_Ktp:      dataCore.KabupatenKotaKtp,
		Alamat_Domisili:         dataCore.AlamatDomisili,
		Provinsi_Domisili:       dataCore.ProvinsiDomisili,
		Kabupaten_Kota_Domisili: dataCore.KabupatenKotaDomisili,
		No_Bpjs:                 dataCore.NoBpjs,
		Kelas_Bpjs:              dataCore.KelasBpjs,
		Foto_Ktp:                dataCore.FotoKtp,
		Foto_Bpjs:               dataCore.FotoBpjs,
		UserID:                  dataCore.UserID,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return patientGorm //insert user
}
func (dataModel *Patient) ModelsToCore() patient.CorePatient { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return patient.CorePatient{
		ID:                    dataModel.ID,
		NoKk:                  dataModel.No_Kk,
		Nik:                   dataModel.Nik,
		NamaPasien:            dataModel.Nama_Pasien,
		JenisKelamin:          dataModel.Jenis_Kelamin,
		TanggalLahir:          dataModel.Tanggal_Lahir,
		Usia:                  dataModel.Usia,
		NamaWali:              dataModel.Nama_Wali,
		EmailWali:             dataModel.Email_Wali,
		NoTelponWali:          dataModel.No_Telpon_Wali,
		AlamatKtp:             dataModel.Alamat_Ktp,
		ProvinsiKtp:           dataModel.Provinsi_Ktp,
		KabupatenKotaKtp:      dataModel.Kabupaten_Kota_Ktp,
		AlamatDomisili:        dataModel.Alamat_Domisili,
		ProvinsiDomisili:      dataModel.Provinsi_Domisili,
		KabupatenKotaDomisili: dataModel.Kabupaten_Kota_Domisili,
		NoBpjs:                dataModel.No_Bpjs,
		KelasBpjs:             dataModel.Kelas_Bpjs,
		FotoKtp:               dataModel.Foto_Ktp,
		FotoBpjs:              dataModel.Foto_Bpjs,
		UserID:                dataModel.UserID,
		CreatedAt:             dataModel.CreatedAt,
		UpdatedAt:             dataModel.UpdatedAt,
	}
}
func ListModelTOCore(dataModel []Patient) []patient.CorePatient { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []patient.CorePatient
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
