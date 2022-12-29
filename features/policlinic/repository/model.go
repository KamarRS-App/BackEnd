package repository

import (
	"github.com/KamarRS-App/KamarRS-App/features/policlinic"

	"gorm.io/gorm"
)

type Policlinic struct {
	gorm.Model
	NamaPoli   string
	JamPraktik string
	HospitalID uint
	Practices  []Practice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Doctors    []Doctor   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Hospital struct {
	gorm.Model
	KodeRs            string
	Nama              string
	Foto              string
	Alamat            string
	Provinsi          string
	KabupatenKota     string
	Kecamatan         string
	KodePos           string
	NoTelpon          string
	Email             string
	KelasRs           string
	Pengelola         string
	JumlahTempatTidur string
	StatusPenggunaan  string
	BiayaPendaftaran  string
	Policlinics       []Policlinic `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Doctor struct {
	gorm.Model
	Nama         string
	Spesialis    string
	Email        string
	NoTelpon     string
	Foto         string
	PoliclinicID uint
}

type Practice struct {
	gorm.Model
	TanggalPraktik string
	KuotaHarian    int
	Status         string
	PoliclinicID   uint
}

func FromCore(dataCore policlinic.CorePoliclinic) Policlinic {
	poliGorm := Policlinic{

		NamaPoli:   dataCore.NamaPoli,
		JamPraktik: dataCore.JamPraktik,
		HospitalID: dataCore.HospitalID,
	}
	return poliGorm //insert user
}
func (dataModel *Policlinic) ToCore() policlinic.CorePoliclinic {
	return policlinic.CorePoliclinic{

		ID:         dataModel.ID,
		NamaPoli:   dataModel.NamaPoli,
		JamPraktik: dataModel.JamPraktik,
		HospitalID: dataModel.HospitalID,
		// Doctor: policlinic.CoreDoctor{
		// 	ID:        dataModel.ToCore().Doctor.ID,
		// 	Nama:      dataModel.ToCore().Doctor.Nama,
		// 	Spesialis: dataModel.ToCore().Doctor.Spesialis,
		// 	Email:     dataModel.ToCore().Doctor.Email,
		// 	NoTelpon:  dataModel.ToCore().Doctor.NoTelpon,
		// 	Foto:      dataModel.ToCore().Doctor.Foto,
		// },
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}
func ToCoreList(dataModel []Policlinic) []policlinic.CorePoliclinic {
	var dataCore []policlinic.CorePoliclinic
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}

//jika diperlukan untuk preload data daily pracctice
// func LoadpraccticeModeltoCore(model []Practice) []practice.Core {
// 	var core []practice.Core
// 	for , v := range model {
// 		core = append(core, v.ModeltoCore())
// 	}
// 	return core

// }
