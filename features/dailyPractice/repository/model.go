package repository

import (
	"github.com/KamarRS-App/features/dailypractice"

	"gorm.io/gorm"
)

type Practice struct {
	gorm.Model
	Tanggal_Praktik string
	Kuota_Harian    int
	Status          string
	PoliclinicID    uint
}

type Policlinic struct {
	gorm.Model
	Nama_Poli   string
	Jam_Praktik string
	HospitalID  uint
	DoctorID    uint
	Practices   []Practice
}

func FromCore(dataCore dailypractice.PracticeCore) Practice {
	practiceGorm := Practice{
		Tanggal_Praktik: dataCore.TanggalPraktik,
		Kuota_Harian:    dataCore.KuotaHarian,
		Status:          dataCore.Status,
		PoliclinicID:    dataCore.PoliclinicID,
	}
	return practiceGorm //insert practice from core
}

//---------------------Daily Practice----------------------------------

func (dataModel *Practice) ToCore() dailypractice.PracticeCore {
	return dailypractice.PracticeCore{
		ID:             dataModel.ID,
		TanggalPraktik: dataModel.Tanggal_Praktik,
		KuotaHarian:    dataModel.Kuota_Harian,
		Status:         dataModel.Status,
		PoliclinicID:   dataModel.PoliclinicID,
	}
}

// mengubah slice struct model gorm ke slice struct core
func ToCoreList(dataModel []Practice) []dailypractice.PracticeCore {
	var dataCore []dailypractice.PracticeCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.ToCore())
	}
	return dataCore
}

//----------------------------------------------------------------------------
