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
