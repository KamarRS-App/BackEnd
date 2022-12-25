package repository

import (
	"kamarRS/features/doctor"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Nama        string
	Bidang      string
	Email       string
	No_Telpon   string
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
		Bidang:    dataCore.Bidang,
		Email:     dataCore.Email,
		No_Telpon: dataCore.NoTelpon,
	}
	return doctorGorm //insert doctor from core
}

//---------------------Doctor----------------------------------

func (dataModel *Doctor) toCore() doctor.DoctorCore {
	return doctor.DoctorCore{
		ID:       dataModel.ID,
		Nama:     dataModel.Nama,
		Bidang:   dataModel.Bidang,
		Email:    dataModel.Email,
		NoTelpon: dataModel.No_Telpon,
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
