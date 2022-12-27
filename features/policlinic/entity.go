package policlinic

import (
	"time"
)

type CorePoliclinic struct {
	ID         uint
	NamaPoli   string `validate:"required"`
	JamPraktik string `validate:"required"`
	HospitalID uint   `validate:"required"`
	DoctorID   uint
	Doctor     CoreDoctor
	CreatedAt  time.Time
	UpdatedAt  time.Time
	// Practices  []CorePractice
}

// type CorePractice struct {
// 	ID             uint
// 	TanggalPraktik string
// 	KuotaHarian    int
// 	Status         string
// 	PoliclinicID   uint
// 	CreatedAt      time.Time
// 	UpdatedAt      time.Time
// }

type CoreDoctor struct {
	ID        uint
	Nama      string
	Spesialis string
	Email     string
	NoTelpon  string
	Foto      string
}

type ServiceInterface interface {
	Create(input CorePoliclinic) (err error)
	GetAll() (data []CorePoliclinic, err error)
	GetById(id int) (data CorePoliclinic, err error)
	Update(input CorePoliclinic, id int) (err error)
	Delete(id int) (err error)
}

type RepositoryInterface interface {
	Create(input CorePoliclinic) (row int, err error)
	GetAll() (data []CorePoliclinic, err error)
	GetById(id int) (data CorePoliclinic, err error)
	Update(input CorePoliclinic, id int) (err error)
	Delete(id int) (row int, err error)
}
