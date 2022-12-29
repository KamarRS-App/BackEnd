package repository

import (
	"github.com/KamarRS-App/KamarRS-App/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama      string
	Email     string
	Nokk      string
	Nik       string
	KataSandi string
	NoTelpon  string
	Patients  []Patient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Patient struct {
	gorm.Model
	NoKk                  string
	Nik                   string
	NamaPasien            string
	JenisKelamin          string
	TanggalLahir          string
	Usia                  int
	NamaWali              string
	EmailWali             string
	NoTelponWali          string
	AlamatKtp             string
	ProvinsiKtp           string
	KabupatenKotaKtp      string
	AlamatDomisili        string
	ProvinsiDomisili      string
	KabupatenKotaDomisili string
	NoBpjs                string
	KelasBpjs             string
	FotoKtp               string
	FotoBpjs              string
	UserID                uint
}

func FromUserCoreToModel(dataCore user.CoreUser) User {
	userGorm := User{
		Nama:      dataCore.Nama,
		Email:     dataCore.Email,
		Nokk:      dataCore.Nokk,
		Nik:       dataCore.Nik,
		KataSandi: dataCore.KataSandi,
		NoTelpon:  dataCore.NoTelpon,
	}
	return userGorm //insert user
}
func (dataModel *User) ModelsToCore() user.CoreUser {
	return user.CoreUser{
		ID:        dataModel.ID,
		Nama:      dataModel.Nama,
		Email:     dataModel.Email,
		Nokk:      dataModel.Nokk,
		Nik:       dataModel.Nik,
		KataSandi: dataModel.KataSandi,
		NoTelpon:  dataModel.NoTelpon,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
		// Patients:  LoadpatientModeltoCore(dataModel.Patients),
	}
}
func ListModelTOCore(dataModel []User) []user.CoreUser {
	var dataCore []user.CoreUser
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}

// jika diperlukan untuk preload data daily pracctice
// func LoadpatientModeltoCore(model []Patient) []user.CorePatient {
// 	var core []user.CorePatient
// 	for , v := range model {
// 		core = append(core, v.ModelsToCore())
// 	}
// 	return core

// }

// func (dataModel *Patient) ModelsToCore() user.CorePatient {
// 	return user.CorePatient{
// 		ID:                    dataModel.ID,
// 		NoKk:                  dataModel.NoKk,
// 		Nik:                   dataModel.Nik,
// 		NamaPasien:            dataModel.NamaPasien,
// 		JenisKelamin:          dataModel.JenisKelamin,
// 		TanggalLahir:          dataModel.TanggalLahir,
// 		Usia:                  dataModel.Usia,
// 		NamaWali:              dataModel.NamaWali,
// 		EmailWali:             dataModel.EmailWali,
// 		NoTelponWali:          dataModel.NoTelponWali,
// 		AlamatKtp:             dataModel.AlamatKtp,
// 		ProvinsiKtp:           dataModel.ProvinsiKtp,
// 		KabupatenKotaKtp:      dataModel.KabupatenKotaKtp,
// 		AlamatDomisili:        dataModel.AlamatDomisili,
// 		ProvinsiDomisili:      dataModel.ProvinsiDomisili,
// 		KabupatenKotaDomisili: dataModel.KabupatenKotaDomisili,
// 		NoBpjs:                dataModel.NoBpjs,
// 		KelasBpjs:             dataModel.KelasBpjs,
// 		FotoKtp:               dataModel.FotoKtp,
// 		FotoBpjs:              dataModel.FotoBpjs,
// 		UserID:                dataModel.UserID,
// 		CreatedAt:             dataModel.CreatedAt,
// 		UpdatedAt:             dataModel.UpdatedAt,
// 	}
// }
