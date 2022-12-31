package repository

import (
	"github.com/KamarRS-App/KamarRS-App/features/doctor"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Nama         string
	Spesialis    string
	Email        string
	NoTelpon     string
	Foto         string
	PoliclinicID uint
}

type Policlinic struct {
	gorm.Model
	NamaPoli   string
	JamPraktik string
	HospitalID uint
	Doctors    []Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func FromCore(dataCore doctor.DoctorCore) Doctor {
	doctorGorm := Doctor{
		Nama:      dataCore.Nama,
		Spesialis: dataCore.Spesialis,
		Email:     dataCore.Email,
		NoTelpon:  dataCore.NoTelpon,
		Foto:      dataCore.Foto,
		PoliclinicID: dataCore.PoliclinicID,
	}
	return doctorGorm //insert doctor from core
}

//---------------------Doctor----------------------------------

func (dataModel *Doctor) ToCore() doctor.DoctorCore {
	return doctor.DoctorCore{
		ID:           dataModel.ID,
		Nama:         dataModel.Nama,
		Spesialis:    dataModel.Spesialis,
		Email:        dataModel.Email,
		NoTelpon:     dataModel.NoTelpon,
		Foto:         dataModel.Foto,
		PoliclinicID: dataModel.PoliclinicID,
	}
}

// mengubah slice struct model gorm ke slice struct core
func ToCoreList(dataModel []Doctor) []doctor.DoctorCore {
	var dataCore []doctor.DoctorCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.ToCore())
	}
	return dataCore
}

//-----------------------------------------------------------------
