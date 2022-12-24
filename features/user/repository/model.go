package repository

import (
	"kamarRS/features/user"

	"gorm.io/gorm"
)

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
}

func FromUserCoreToModel(dataCore user.CoreUser) User { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	userGorm := User{
		Nama:       dataCore.Nama,
		Email:      dataCore.Email,
		No_kk:      dataCore.Nokk,
		Nik:        dataCore.Nik,
		Kata_Sandi: dataCore.KataSandi,
		No_Telpon:  dataCore.NoTelpon,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return userGorm //insert user
}
func (dataModel *User) ModelsToCore() user.CoreUser { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return user.CoreUser{

		Nama:      dataModel.Nama,
		Email:     dataModel.Email,
		Nokk:      dataModel.No_kk,
		Nik:       dataModel.Nik,
		KataSandi: dataModel.Kata_Sandi,
		NoTelpon:  dataModel.No_Telpon,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
		Patients:  LoadpatientModeltoCore(dataModel.Patients),
	}
}
func ListModelTOCore(dataModel []User) []user.CoreUser { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []user.CoreUser
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}

// jika diperlukan untuk preload data daily pracctice
func LoadpatientModeltoCore(model []Patient) []user.CorePatient {
	var core []user.CorePatient
	for _, v := range model {
		core = append(core, v.ModelsToCore())
	}
	return core

}

func (dataModel *Patient) ModelsToCore() user.CorePatient { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return user.CorePatient{
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
