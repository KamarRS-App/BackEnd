package repository

import (
	"kamarRS/features/bedreservation"

	"gorm.io/gorm"
)

type BedReservation struct {
	gorm.Model
	Hospital_Id       uint
	Status_Pasien     string
	Biaya_Registrasi  int
	Order_Id          string
	Payment_Method    string
	Link_Pembayaran   string
	Status_Pembayaran string
	PatientID         uint
	BedID             uint
	Patient           Patient
}

type Patient struct {
	gorm.Model
	No_Kk                   string
	Nik                     string
	Nama_Pasien             string
	Jenis_Kelamin           string
	Tanggal_Lahir           string
	Usia                    int
	Nama_Wali               string
	Email_Wali              string
	No_Telpon_Wali          string
	Alamat_Ktp              string
	Provinsi_Ktp            string
	Kabupaten_Kota_Ktp      string
	Alamat_Domisili         string
	Provinsi_Domisili       string
	Kabupaten_Kota_Domisili string
	No_Bpjs                 string
	Kelas_Bpjs              string
	Foto_Ktp                string
	Foto_Bpjs               string
	UserID                  uint
	BedReservationID        uint
	// BedReservation          BedReservation
}

type Bed struct {
	gorm.Model
	Nama_Tempat_Tidur string
	Ruangan           string
	Kelas             string
	Status            string
	HospitalID        uint
	BedReservation    BedReservation
}

func FromCoreToModel(dataCore bedreservation.BedReservationCore) BedReservation {
	bedresGorm := BedReservation{
		Status_Pasien:     dataCore.StatusPasien,
		Biaya_Registrasi:  dataCore.BiayaRegistrasi,
		Order_Id:          dataCore.OrderID,
		Payment_Method:    dataCore.PaymentMethod,
		Link_Pembayaran:   dataCore.LinkPembayaran,
		Status_Pembayaran: dataCore.StatusPembayaran,
		Hospital_Id:       dataCore.HospitalID,
		PatientID:         dataCore.PatientID,
		BedID:             dataCore.BedID,
	}
	return bedresGorm //insert bedreserve from core
}

//----------------------BedReserve Aja-------------------------------

func (dataModel *BedReservation) toCore() bedreservation.BedReservationCore {
	return bedreservation.BedReservationCore{
		ID:               dataModel.ID,
		StatusPasien:     dataModel.Status_Pasien,
		BiayaRegistrasi:  dataModel.Biaya_Registrasi,
		OrderID:          dataModel.Order_Id,
		PaymentMethod:    dataModel.Payment_Method,
		LinkPembayaran:   dataModel.Link_Pembayaran,
		StatusPembayaran: dataModel.Status_Pembayaran,
		HospitalID:       dataModel.Hospital_Id,
		BedID:            dataModel.BedID,
		Patient:          dataModel.Patient.toCoreP(),
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []BedReservation) []bedreservation.BedReservationCore {
	var dataCore []bedreservation.BedReservationCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

//----------------------Patient Aja-------------------------------

func (dataModel *Patient) toCoreP() bedreservation.PatientCore {
	return bedreservation.PatientCore{
		ID:                    dataModel.ID,
		NoKk:                  dataModel.No_Kk,
		Nik:                   dataModel.Nik,
		NamaPasien:            dataModel.Nama_Pasien,
		JenisKelamin:          dataModel.Jenis_Kelamin,
		TanggalLahir:          dataModel.Tanggal_Lahir,
		Usia:                  dataModel.Usia,
		NamaWali:              dataModel.Nama_Wali,
		EmailWali:             dataModel.Email_Wali,
		NoTelponWali:          dataModel.No_Telpon_Wali,
		AlamatKtp:             dataModel.Alamat_Ktp,
		ProvinsiKtp:           dataModel.Provinsi_Ktp,
		KabupatenKotaKtp:      dataModel.Kabupaten_Kota_Ktp,
		AlamatDomisili:        dataModel.Alamat_Domisili,
		ProvinsiDomisili:      dataModel.Provinsi_Domisili,
		KabupatenKotaDomisili: dataModel.Kabupaten_Kota_Domisili,
		NoBpjs:                dataModel.No_Bpjs,
		KelasBpjs:             dataModel.Kelas_Bpjs,
		FotoKtp:               dataModel.Foto_Ktp,
		FotoBpjs:              dataModel.Foto_Bpjs,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreListP(dataModel []Patient) []bedreservation.PatientCore {
	var dataCore []bedreservation.PatientCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreP())
	}
	return dataCore
}

//---------------------------------------------------------------------------------
