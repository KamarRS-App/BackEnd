package hospitalstaff

type HospitalStaffCore struct {
	ID         uint
	Nama       string
	Email      string
	KataSandi  string
	Peran      string
	HospitalID uint
}

type ServiceInterface interface { //sebagai contract yang dibuat di layer service

	Create(input HospitalStaffCore) (err error) // menambahkah data user berdasarkan data usercore
	Update(id int, input HospitalStaffCore) error
	GetStaff(id int) (data HospitalStaffCore, err error)
	DeleteById(id int) error
}

type RepositoryInterface interface { // berkaitan database

	Create(input HospitalStaffCore) (err error)
	Update(id int, input HospitalStaffCore) error
	GetStaff(id int) (data HospitalStaffCore, err error)
	DeleteById(id int) error
}
