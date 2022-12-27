package repository

import (
	hospitalstaff "github.com/KamarRS-App/KamarRS-App/features/hospitalStaff"

	"gorm.io/gorm"
)

type HospitalStaff struct {
	gorm.Model
	Nama       string
	Email      string
	Kata_Sandi string
	Peran      string
	HospitalID uint
}

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
	Jumlah_Tempat_Tidur string
	Status_Penggunaan   string
	Biaya_Pendaftaran   string
	HospitalStaffs      []HospitalStaff
}

func FromStaffCore(dataCore hospitalstaff.HospitalStaffCore) HospitalStaff { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	staffGorm := HospitalStaff{
		Nama:       dataCore.Nama,
		Email:      dataCore.Email,
		Kata_Sandi: dataCore.KataSandi,
		Peran:      dataCore.Peran,
		HospitalID: dataCore.HospitalID,
	}
	return staffGorm //insert user
}
func (dataModel *HospitalStaff) ModelsToCore() hospitalstaff.HospitalStaffCore { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return hospitalstaff.HospitalStaffCore{
		ID:         dataModel.ID,
		Nama:       dataModel.Nama,
		Email:      dataModel.Email,
		KataSandi:  dataModel.Kata_Sandi,
		Peran:      dataModel.Peran,
		HospitalID: dataModel.HospitalID,
	}
}

func ListModelTOCore(dataModel []HospitalStaff) []hospitalstaff.HospitalStaffCore { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []hospitalstaff.HospitalStaffCore
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
