package repository

import (
	"github.com/KamarRS-App/KamarRS-App/features/hospital"

	"gorm.io/gorm"
)

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
	JumlahTempatTidur int
	StatusPenggunaan  string
	BiayaPendaftaran  int
	HospitalStaffs    []HospitalStaff `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Beds              []Bed           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Policlinics       []Policlinic    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type HospitalStaff struct {
	gorm.Model
	Nama       string
	Email      string
	KataSandi  string
	Peran      string
	HospitalID uint
}

type Bed struct {
	gorm.Model
	NamaTempatTidur string
	Ruangan         string
	Kelas           string
	Status          string
	HospitalID      uint
}

type Policlinic struct {
	gorm.Model
	NamaPoli   string
	JamPraktik string
	HospitalID uint
}

func FromCore(dataCore hospital.HospitalCore) Hospital {
	hospitalGorm := Hospital{
		KodeRs:            dataCore.KodeRs,
		Nama:              dataCore.Nama,
		Foto:              dataCore.Foto,
		Alamat:            dataCore.Alamat,
		Provinsi:          dataCore.Provinsi,
		KabupatenKota:     dataCore.KabupatenKota,
		Kecamatan:         dataCore.Kecamatan,
		KodePos:           dataCore.KodePos,
		NoTelpon:          dataCore.NoTelpon,
		Email:             dataCore.Email,
		KelasRs:           dataCore.KelasRs,
		Pengelola:         dataCore.PemilikPengelola,
		JumlahTempatTidur: dataCore.JumlahTempatTidur,
		StatusPenggunaan:  dataCore.StatusPenggunaan,
		BiayaPendaftaran:  dataCore.BiayaRegistrasi,
	}
	return hospitalGorm //insert hospital from core
}

//---------------------Hospital----------------------------------

func (dataModel *Hospital) ToCore() hospital.HospitalCore {
	return hospital.HospitalCore{
		ID:                dataModel.ID,
		KodeRs:            dataModel.KodeRs,
		Nama:              dataModel.Nama,
		Foto:              dataModel.Foto,
		Alamat:            dataModel.Alamat,
		Provinsi:          dataModel.Provinsi,
		KabupatenKota:     dataModel.KabupatenKota,
		Kecamatan:         dataModel.Kecamatan,
		KodePos:           dataModel.KodePos,
		NoTelpon:          dataModel.NoTelpon,
		Email:             dataModel.Email,
		KelasRs:           dataModel.KelasRs,
		PemilikPengelola:  dataModel.Pengelola,
		JumlahTempatTidur: dataModel.JumlahTempatTidur,
		StatusPenggunaan:  dataModel.StatusPenggunaan,
		BiayaRegistrasi:   dataModel.BiayaPendaftaran,
	}
}

// mengubah slice struct model gorm ke slice struct core
func ToCoreList(dataModel []Hospital) []hospital.HospitalCore {
	var dataCore []hospital.HospitalCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.ToCore())
	}
	return dataCore
}

//----------------------------------------------------------------------------
