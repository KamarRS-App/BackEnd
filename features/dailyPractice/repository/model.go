package repository

import (
	dailypractice "kamarRS/features/dailyPractice"

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

func FromCoreToModel(dataCore dailypractice.DailyPracticeCore) Practice {
	practiceGorm := Practice{
		Tanggal_Praktik: dataCore.TanggalPraktik,
		Kuota_Harian:    dataCore.KuotaHarian,
		Status:          dataCore.Status,
		PoliclinicID:    dataCore.PoliclinicID,
	}
	return practiceGorm //insert practice from core
}

//---------------------Daily Practice----------------------------------

func (dataModel *Practice) toCore() dailypractice.DailyPracticeCore {
	return dailypractice.DailyPracticeCore{
		ID:             dataModel.ID,
		TanggalPraktik: dataModel.Tanggal_Praktik,
		KuotaHarian:    dataModel.Kuota_Harian,
		Status:         dataModel.Status,
		PoliclinicID:   dataModel.PoliclinicID,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Practice) []dailypractice.DailyPracticeCore {
	var dataCore []dailypractice.DailyPracticeCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

//----------------------------------------------------------------------------
