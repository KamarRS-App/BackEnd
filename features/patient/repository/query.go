package repository

import (
	"errors"
	"fmt"

	"github.com/KamarRS-App/KamarRS-App/features/patient"
	"github.com/KamarRS-App/KamarRS-App/features/user/repository"

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

	// var patient []Patient

	// tx1 := repo.db.Find(&patient)
	// if tx1.Error != nil {
	// 	return tx1.Error
	// }

	// for _, v := range patient {
	// 	if input.Nik == v.Nik || input.NoBpjs == v.NoBpjs {
	// 		return errors.New("eror input data")
	// 	}

	// }
	var user repository.User
	tx2 := repo.db.First(&user, input.UserID)
	if tx2.Error != nil {
		return tx2.Error
	}

	if user.Nokk == "" {
		fmt.Println("anda harus melengkapai data diri anda terlebih dahulu untuk mendaftarkan pasien")
		return errors.New("anda harus melengkapai data diri anda terlebih dahulu untuk mendaftarkan pasien")
	}

	if input.NoKk != user.Nokk {
		fmt.Println("hanya bisa mendaftarkan keluarga")
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

	tx1 := repo.db.Unscoped().Delete(&patiens, id)
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
func (repo *patientRepository) GetByUserId(limit, offset, id int) (data []patient.CorePatient, totalpage int, err error) {
	var patient []Patient
	var count int64
	rx := repo.db.Model(&patient).Where("user_id", id).Count(&count)
	if rx.Error != nil {
		return nil, 0, rx.Error
	}
	fmt.Println("count", count)
	if rx.RowsAffected == 0 {
		return nil, 0, errors.New("error query count")
	}
	// var totalpage int
	if int(count)%limit == 0 {
		totalpage = int(count) / limit
	} else {
		totalpage = (int(count) / limit) + 1
	}
	var patients []Patient

	tx := repo.db.Where("user_id=?", id).Limit(limit).Offset(offset).Find(&patients)

	if tx.Error != nil {

		return nil, 0, tx.Error
	}
	gorms := ListModelTOCore(patients)
	return gorms, totalpage, nil
}

// Update implements patient.RepositoryInterface
func (repo *patientRepository) Update(id, userId int, input patient.CorePatient) error {
	var patient []Patient

	tx3 := repo.db.Find(&patient)
	if tx3.Error != nil {
		return tx3.Error
	}

	// for _, v := range patient {
	// 	if input.Nik == v.Nik || input.NoBpjs == v.NoBpjs {
	// 		return errors.New("eror input data")
	// 	}

	// }
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
	if userId != int(users.ID) {
		fmt.Println("Buka pasien yang di daftarkan")
		return errors.New("buka pasien yang di daftarkan")
	}
	if input.NoKk != "" {
		if input.NoKk != users.Nokk {
			fmt.Println("no kk pasien harus sama dengan no kk user")
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
