package repository

import (
	"github.com/KamarRS-App/features/doctor"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Nama        string
	Spesialis   string
	Email       string
	No_Telpon   string
	Foto        string
	Policlinics []Policlinic
}

type Policlinic struct {
	gorm.Model
	Nama_Poli   string
	Jam_Praktik string
	HospitalID  uint
	DoctorID    uint
}

func FromCoreToModel(dataCore doctor.DoctorCore) Doctor {
	doctorGorm := Doctor{
		Nama:      dataCore.Nama,
		Spesialis: dataCore.Spesialis,
		Email:     dataCore.Email,
		No_Telpon: dataCore.NoTelpon,
		Foto:      dataCore.Foto,
	}
	return doctorGorm //insert doctor from core
}

//---------------------Doctor----------------------------------

func (dataModel *Doctor) toCore() doctor.DoctorCore {
	return doctor.DoctorCore{
		ID:        dataModel.ID,
		Nama:      dataModel.Nama,
		Spesialis: dataModel.Spesialis,
		Email:     dataModel.Email,
		NoTelpon:  dataModel.No_Telpon,
		Foto:      dataModel.Foto,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Doctor) []doctor.DoctorCore {
	var dataCore []doctor.DoctorCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

//-----------------------------------------------------------------
