package repository

import (
	"github.com/KamarRS-App/KamarRS-App/features/policlinic"

	"gorm.io/gorm"
)

type Policlinic struct {
	gorm.Model
	Nama_Poli   string
	Jam_Praktik string
	HospitalID  uint
	DoctorID    uint
	Practices   []Practice
	Doctor      Doctor
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
	Spesialis   string
	Email       string
	No_Telpon   string
	Foto        string
	Policlinics []Policlinic
}

type Practice struct {
	gorm.Model
	Tanggal_Praktik string
	Kuota_Harian    int
	Status          string
	PoliclinicID    uint
}

func FromPoliclinicCoreToModel(dataCore policlinic.CorePoliclinic) Policlinic {
	poliGorm := Policlinic{

		Nama_Poli:   dataCore.NamaPoli,
		Jam_Praktik: dataCore.JamPraktik,
		HospitalID:  dataCore.HospitalID,
		DoctorID:    dataCore.DoctorID,
	}
	return poliGorm //insert user
}
func (dataModel *Policlinic) ModelsToCore() policlinic.CorePoliclinic {
	return policlinic.CorePoliclinic{

		ID:         dataModel.ID,
		NamaPoli:   dataModel.Nama_Poli,
		JamPraktik: dataModel.Jam_Praktik,
		HospitalID: dataModel.HospitalID,
		Doctor: policlinic.CoreDoctor{
			ID:        dataModel.Doctor.ID,
			Nama:      dataModel.Doctor.Nama,
			Spesialis: dataModel.Doctor.Spesialis,
			Email:     dataModel.Doctor.Email,
			NoTelpon:  dataModel.Doctor.No_Telpon,
			Foto:      dataModel.Doctor.Foto,
		},
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}
func ListModelTOCore(dataModel []Policlinic) []policlinic.CorePoliclinic {
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
