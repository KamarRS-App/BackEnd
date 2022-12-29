package repository

import (
	dailypractice "github.com/KamarRS-App/KamarRS-App/features/dailyPractice"
	"gorm.io/gorm"
)

type Practice struct {
	gorm.Model
	TanggalPraktik string
	KuotaHarian    int
	Status         string
	PoliclinicID   uint
}

type Policlinic struct {
	gorm.Model
	NamaPoli   string
	JamPraktik string
	HospitalID uint
	DoctorID   uint
	Practices  []Practice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func FromCore(dataCore dailypractice.PracticeCore) Practice {
	practiceGorm := Practice{
		TanggalPraktik: dataCore.TanggalPraktik,
		KuotaHarian:    dataCore.KuotaHarian,
		Status:         dataCore.Status,
		PoliclinicID:   dataCore.PoliclinicID,
	}
	return practiceGorm //insert practice from core
}

//---------------------Daily Practice----------------------------------

func (dataModel *Practice) ToCore() dailypractice.PracticeCore {
	return dailypractice.PracticeCore{
		ID:             dataModel.ID,
		TanggalPraktik: dataModel.TanggalPraktik,
		KuotaHarian:    dataModel.KuotaHarian,
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
