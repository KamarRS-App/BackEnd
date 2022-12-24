package repository

import (
	"kamarRS/features/policlinic"

	"gorm.io/gorm"
)

type Policlinic struct {
	gorm.Model
	Nama_Poli   string
	Jam_Praktik string
	HospitalID  uint
	DoctorID    uint
	Practices   []Practice
}

type Hospital struct {
	gorm.Model
	Kode_Rs             string
	Nama                string
	Foto                string
	Alamat              string
	Provinsi            string
	Kabupaten_Kota      string
	Kecamatan           string
	No_Telpon           string
	Email               string
	Kelas_Rs            string
	Pengelola           string
	Jumlah_Tempat_Tidur string
	Status_Penggunaan   string
	Biaya_Pendaftaran   string
	Policlinics         []Policlinic
}

type Doctor struct {
	gorm.Model
	Nama        string
	Bidang      string
	Email       string
	No_Telpon   string
	Policlinics []Policlinic
}

type Practice struct {
	gorm.Model
	Tanggal_Praktik string
	Kuota_Harian    int
	Status          string
	PoliclinicID    uint
}

func FromPoliclinicCoreToModel(dataCore policlinic.CorePoliclinic) Policlinic { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	poliGorm := Policlinic{

		Nama_Poli:   dataCore.NamaPoli,
		Jam_Praktik: dataCore.JamPraktik,
		HospitalID:  dataCore.HospitalID,
		DoctorID:    dataCore.DoctorID,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return poliGorm //insert user
}
func (dataModel *Policlinic) ModelsToCore() policlinic.CorePoliclinic { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return policlinic.CorePoliclinic{

		ID:         dataModel.ID,
		NamaPoli:   dataModel.Nama_Poli,
		JamPraktik: dataModel.Jam_Praktik,
		HospitalID: dataModel.HospitalID,
		DoctorID:   dataModel.DoctorID,
		CreatedAt:  dataModel.CreatedAt,
		UpdatedAt:  dataModel.UpdatedAt,
	}
}
func ListModelTOCore(dataModel []Policlinic) []policlinic.CorePoliclinic { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []policlinic.CorePoliclinic
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}

//jika diperlukan untuk preload data daily pracctice
// func LoadpraccticeModeltoCore(model []Practice) []practice.Core {
// 	var core []practice.Core
// 	for _, v := range model {
// 		core = append(core, v.ModeltoCore())
// 	}
// 	return core

// }
