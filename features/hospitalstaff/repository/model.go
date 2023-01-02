package repository

import (
	"github.com/KamarRS-App/KamarRS-App/features/hospitalstaff"

	"gorm.io/gorm"
)

type HospitalStaff struct {
	gorm.Model
	Nama         string
	Email        string
	KataSandi    string
	Peran        string
	HospitalID   uint
	Hospital     Hospital
	HospitalName string
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
	HospitalStaffs    []HospitalStaff `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func FromStaffCore(dataCore hospitalstaff.HospitalStaffCore) HospitalStaff { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	staffGorm := HospitalStaff{
		Nama:         dataCore.Nama,
		Email:        dataCore.Email,
		KataSandi:    dataCore.KataSandi,
		Peran:        dataCore.Peran,
		HospitalID:   dataCore.HospitalID,
		HospitalName: dataCore.HospitalName,
	}
	return staffGorm //insert user
}
func (dataModel *HospitalStaff) ModelsToCore() hospitalstaff.HospitalStaffCore { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return hospitalstaff.HospitalStaffCore{
		ID:           dataModel.ID,
		Nama:         dataModel.Nama,
		Email:        dataModel.Email,
		KataSandi:    dataModel.KataSandi,
		Peran:        dataModel.Peran,
		HospitalID:   dataModel.HospitalID,
		HospitalName: dataModel.HospitalName,
		// Hospital: hospitalstaff.HospitalCore{
		// 	Nama: dataModel.Hospital.Nama,
		// },
	}
}

func ListModelTOCore(dataModel []HospitalStaff) []hospitalstaff.HospitalStaffCore { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []hospitalstaff.HospitalStaffCore
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}

func (dataModel *HospitalStaff) ModelsToCorePreload() hospitalstaff.HospitalStaffCore { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return hospitalstaff.HospitalStaffCore{
		ID:         dataModel.ID,
		Nama:       dataModel.Nama,
		Email:      dataModel.Email,
		KataSandi:  dataModel.KataSandi,
		Peran:      dataModel.Peran,
		HospitalID: dataModel.HospitalID,
		Hospital: hospitalstaff.HospitalCore{
			Nama: dataModel.Hospital.Nama,
		},
	}
}

func ListModelTOCorePreload(dataModel []HospitalStaff) []hospitalstaff.HospitalStaffCore { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []hospitalstaff.HospitalStaffCore
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
