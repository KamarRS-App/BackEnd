package repository

import (
	"errors"

	"github.com/KamarRS-App/features/patient"

	"gorm.io/gorm"
)

type patientRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) patient.RepositoryInterface { // user.repository mengimplementasikan interface repository yang ada di entities
	return &patientRepository{
		db: db,
	}

}

// Create implements patient.RepositoryInterface
func (repo *patientRepository) Create(input patient.CorePatient) (err error) {
	var patient []Patient

	tx1 := repo.db.Find(&patient)
	if tx1.Error != nil {
		return tx1.Error
	}

	for _, v := range patient {
		if input.Nik == v.Nik || input.NoKk == v.No_Kk || input.NoBpjs == v.No_Bpjs {
			return errors.New("eror input data")
		}

	}

	patientGorm := FromPatientCore(input)

	tx := repo.db.Create(&patientGorm) // proses insert data

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// DeleteById implements patient.RepositoryInterface
func (*patientRepository) DeleteById(id int) error {
	panic("unimplemented")
}

// GetByPatientId implements patient.RepositoryInterface
func (*patientRepository) GetByPatientId(id int) (data patient.CorePatient, err error) {
	panic("unimplemented")
}

// GetByUserId implements patient.RepositoryInterface
func (*patientRepository) GetByUserId(userid int) (data patient.CorePatient, err error) {
	panic("unimplemented")
}

// Update implements patient.RepositoryInterface
func (*patientRepository) Update(id int, input patient.CorePatient) error {
	panic("unimplemented")
}
