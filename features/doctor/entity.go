package doctor

type DoctorCore struct {
	ID           uint
	Nama         string
	Spesialis    string
	Email        string
	NoTelpon     string
	Foto         string
	PoliclinicID uint
}

type ServiceInterface interface {
	Create(input DoctorCore) (err error)
	GetAll() (data []DoctorCore, err error)
	GetById(id int) (data DoctorCore, err error)
	Update(input DoctorCore, id int) (err error)
	Delete(id int) (err error)
}

type RepositoryInterface interface {
	Create(input DoctorCore) (row int, err error)
	GetAll() (data []DoctorCore, err error)
	GetById(id int) (data DoctorCore, err error)
	Update(input DoctorCore, id int) (err error)
	Delete(id int) (row int, err error)
}
