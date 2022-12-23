package mysql

import (
	"fmt"
	"log"

	"kamarRS/config"

	bedRepo "kamarRS/features/bed/repository"
	bedReservationRepo "kamarRS/features/bedReservation/repository"
	checkupReservationRepo "kamarRS/features/checkupReservation/repository"
	dailyPracticeRepo "kamarRS/features/dailyPractice/repository"
	doctorRepo "kamarRS/features/doctor/repository"
	hospitalRepo "kamarRS/features/hospital/repository"
	hospitalStaffRepo "kamarRS/features/hospitalStaff/repository"
	patientRepo "kamarRS/features/patient/repository"
	policlinicRepo "kamarRS/features/policlinic/repository"
	userRepo "kamarRS/features/user/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func DBMigration(db *gorm.DB) {
	db.AutoMigrate(&userRepo.User{})
	db.AutoMigrate(&patientRepo.Patient{})
	db.AutoMigrate(&hospitalRepo.Hospital{})
	db.AutoMigrate(&hospitalStaffRepo.HospitalStaff{})
	db.AutoMigrate(&bedRepo.Bed{})
	db.AutoMigrate(&bedReservationRepo.BedReservation{})
	db.AutoMigrate(&doctorRepo.Doctor{})
	db.AutoMigrate(&policlinicRepo.Policlinic{})
	db.AutoMigrate(&dailyPracticeRepo.Practice{})
	db.AutoMigrate(&checkupReservationRepo.CheckupReservation{})

}
