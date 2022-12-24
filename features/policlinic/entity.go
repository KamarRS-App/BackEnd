package policlinic

import (
	"time"
)

type CorePoliclinic struct {
	ID         uint
	NamaPoli   string `validate:"required"`
	JamPraktik string `validate:"required"`
	HospitalID uint   `validate:"required"`
	DoctorID   uint   `validate:"required"`
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
