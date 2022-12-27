package repository

import (
	"errors"

	"github.com/KamarRS-App/features/patient"
	"github.com/KamarRS-App/features/user/repository"

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
	var user repository.User
	tx2 := repo.db.First(&user, input.UserID)
	if tx2.Error != nil {
		return tx2.Error
	}

	if user.No_kk == "" {
		return errors.New("anda harus melengkapai data diri anda terlebih dahulu untuk mendaftarkan pasien")
	}

	if input.NoKk != user.No_kk {
		return errors.New("hanya bisa mendaftarkan keluarga")
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
func (repo *patientRepository) DeleteById(id int) error {
	patiens := Patient{}

	tx1 := repo.db.Delete(&patiens, id)
	if tx1.Error != nil {
		return tx1.Error
	}

	if tx1.RowsAffected == 0 {
		return errors.New("id not found")

	}

	return nil
}

// GetByPatientId implements patient.RepositoryInterface
func (repo *patientRepository) GetByPatientId(id int) (data patient.CorePatient, err error) {
	var patients Patient

	tx := repo.db.First(&patients, id)

	if tx.Error != nil {

		return patient.CorePatient{}, tx.Error
	}

	gorms := patients.ModelsToCore()
	return gorms, nil
}

// GetByUserId implements patient.RepositoryInterface
func (repo *patientRepository) GetByUserId(userid int) (data []patient.CorePatient, err error) {
	var patients []Patient

	tx := repo.db.Where("user_id=?", userid).Find(&patients)

	if tx.Error != nil {

		return nil, tx.Error
	}
	gorms := ListModelTOCore(patients)
	return gorms, nil
}

// Update implements patient.RepositoryInterface
func (repo *patientRepository) Update(id int, input patient.CorePatient) error {
	var pasien Patient

	tx1 := repo.db.First(&pasien, id)

	if tx1.Error != nil {

		return tx1.Error
	}
	var users User

	tx2 := repo.db.First(&users, pasien.UserID)

	if tx2.Error != nil {

		return tx2.Error
	}

	if input.NoKk != "" {
		if input.NoKk != users.No_kk {
			return errors.New("no kk pasien harus sama dengan no kk user")
		}

	}

	patientGorm := FromPatientCore(input)

	tx := repo.db.Model(&patientGorm).Where("id = ?", id).Updates(&patientGorm)

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// GetAllPatient implements patient.RepositoryInterface
func (repo *patientRepository) GetAllPatient() (data []patient.CorePatient, err error) {
	var patients []Patient

	tx := repo.db.Find(&patients)

	if tx.Error != nil {

		return nil, tx.Error
	}
	gorms := ListModelTOCore(patients)
	return gorms, nil
}
