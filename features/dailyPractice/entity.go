package dailypractice

type PracticeCore struct {
	ID             uint
	TanggalPraktik string
	KuotaHarian    int
	Status         string
	PoliclinicID   uint
}

type ServiceInterface interface {
	Create(input PracticeCore) (err error)
	GetAll() (data []PracticeCore, err error)
	GetById(id int) (data PracticeCore, err error)
	Update(input PracticeCore, id int) (err error)
}

type RepositoryInterface interface {
	Create(input PracticeCore) (row int, err error)
	GetAll() (data []PracticeCore, err error)
	GetById(id int) (data PracticeCore, err error)
	Update(input PracticeCore, id int) (err error)
}
