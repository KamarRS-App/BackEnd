package bed

type BedCore struct {
	ID              uint
	NamaTempatTidur string
	Ruangan         string
	Kelas           string
	Status          string
	HospitalID      uint
	Hospital        HospitalCore
}

type HospitalCore struct {
	ID   uint
	Nama string
}

type ServiceInterface interface {
	Create(input BedCore) (err error)
	GetAll(limit, offset, id int) (data []BedCore, err error)
	GetById(id int) (data BedCore, err error)
	Update(input BedCore, id int) (err error)
	Delete(id int) (err error)
}

type RepositoryInterface interface {
	Create(input BedCore) (row int, err error)
	GetAll(limit, offset, id int) (data []BedCore, err error)
	GetById(id int) (data BedCore, err error)
	Update(input BedCore, id int) (err error)
	Delete(id int) (row int, err error)
}
