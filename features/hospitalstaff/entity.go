package hospitalstaff

type HospitalStaffCore struct {
	ID         uint
	Nama       string
	Email      string
	KataSandi  string
	Peran      string
	HospitalID uint
	Hospital   HospitalCore
}

type ServiceInterface interface { //sebagai contract yang dibuat di layer service

	Create(input HospitalStaffCore) (err error) // menambahkah data user berdasarkan data usercore
	Update(id int, input HospitalStaffCore) error
	GetStaff(id int) (data HospitalStaffCore, err error)
	GetAllStaff(limit, page int) (data []HospitalStaffCore, totalPage int, err error)
	DeleteById(id int) error
}

type RepositoryInterface interface { // berkaitan database

	Create(input HospitalStaffCore) (err error)
	Update(id int, input HospitalStaffCore) error
	GetStaff(id int) (data HospitalStaffCore, err error)
	GetAllStaff(limit, offset int) (data []HospitalStaffCore, totalPage int, err error)
	DeleteById(id int) error
}

type HospitalCore struct {
	ID                uint
	KodeRs            string
	Nama              string
	Foto              string
	Alamat            string
	Provinsi          string
	KabupatenKota     string
	Kecamatan         string
	NoTelpon          string
	Email             string
	KelasRs           string
	PemilikPengelola  string
	JumlahTempatTidur int
	StatusPenggunaan  string
	BiayaRegistrasi   int
}
