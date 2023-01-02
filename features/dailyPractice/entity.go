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
	GetAll(pagination, limit, id int) (data []PracticeCore, totalpage int, err error)
	GetById(id int) (data PracticeCore, err error)
	Update(input PracticeCore, id int) (err error)
}

type RepositoryInterface interface {
	Create(input PracticeCore) (row int, err error)
	GetAll(limit, offset, id int) (data []PracticeCore, totalpage int, err error)
	GetById(id int) (data PracticeCore, err error)
	Update(input PracticeCore, id int) (err error)
}
