package repository

import (
	"github.com/KamarRS-App/features/hospital"

	"gorm.io/gorm"
)

type Hospital struct {
	gorm.Model
	Kode_Rs             string
	Nama                string
	Foto                string
	Alamat              string
	Provinsi            string
	Kabupaten_Kota      string
	Kecamatan           string
	No_Telpon           string
	Email               string
	Kelas_Rs            string
	Pengelola           string
	Jumlah_Tempat_Tidur int
	Status_Penggunaan   string
	Biaya_Pendaftaran   int
	HospitalStaffs      []HospitalStaff
	Beds                []Bed
	Policlinics         []Policlinic
}

type HospitalStaff struct {
	gorm.Model
	Nama       string
	Email      string
	Kata_Sandi string
	Peran      string
	HospitalID uint
}

type Bed struct {
	gorm.Model
	Nama_Tempat_Tidur string
	Ruangan           string
	Kelas             string
	Status            string
	HospitalID        uint
}

type Policlinic struct {
	gorm.Model
	Nama_Poli   string
	Jam_Praktik string
	HospitalID  uint
}

func FromCore(dataCore hospital.HospitalCore) Hospital {
	hospitalGorm := Hospital{
		Kode_Rs:             dataCore.KodeRs,
		Nama:                dataCore.Nama,
		Foto:                dataCore.Foto,
		Alamat:              dataCore.Alamat,
		Provinsi:            dataCore.Provinsi,
		Kabupaten_Kota:      dataCore.KabupatenKota,
		Kecamatan:           dataCore.Kecamatan,
		No_Telpon:           dataCore.NoTelpon,
		Email:               dataCore.Email,
		Kelas_Rs:            dataCore.KelasRs,
		Pengelola:           dataCore.PemilikPengelola,
		Jumlah_Tempat_Tidur: dataCore.JumlahTempatTidur,
		Status_Penggunaan:   dataCore.StatusPenggunaan,
		Biaya_Pendaftaran:   dataCore.BiayaRegistrasi,
	}
	return hospitalGorm //insert hospital from core
}

//---------------------Hospital----------------------------------

func (dataModel *Hospital) toCore() hospital.HospitalCore {
	return hospital.HospitalCore{
		ID:                dataModel.ID,
		KodeRs:            dataModel.Kode_Rs,
		Nama:              dataModel.Nama,
		Foto:              dataModel.Foto,
		Alamat:            dataModel.Alamat,
		Provinsi:          dataModel.Provinsi,
		KabupatenKota:     dataModel.Kabupaten_Kota,
		Kecamatan:         dataModel.Kecamatan,
		NoTelpon:          dataModel.No_Telpon,
		Email:             dataModel.Email,
		KelasRs:           dataModel.Kelas_Rs,
		PemilikPengelola:  dataModel.Pengelola,
		JumlahTempatTidur: dataModel.Jumlah_Tempat_Tidur,
		StatusPenggunaan:  dataModel.Status_Penggunaan,
		BiayaRegistrasi:   dataModel.Biaya_Pendaftaran,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Hospital) []hospital.HospitalCore {
	var dataCore []hospital.HospitalCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

//----------------------------------------------------------------------------
