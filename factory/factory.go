package factory

// import (
// 	bedDelivery "kamarRS/features/bed/delivery"
// 	bedRepo "kamarRS/features/bed/repository"
// 	bedService "kamarRS/features/bed/service"

// 	bedReservationDelivery "kamarRS/features/bedReservation/delivery"
// 	bedReservationRepo "kamarRS/features/bedReservation/repository"
// 	bedReservationService "kamarRS/features/bedReservation/service"

// 	checkupReservationDelivery "kamarRS/features/checkupReservation/delivery"
// 	checkupReservationRepo "kamarRS/features/checkupReservation/repository"
// 	checkupReservationService "kamarRS/features/checkupReservation/service"

// 	dailyPracticeDelivery "kamarRS/features/dailyPractice/delivery"
// 	dailyPracticeRepo "kamarRS/features/dailyPractice/repository"
// 	dailyPracticeService "kamarRS/features/dailyPractice/service"

// 	doctorDelivery "kamarRS/features/doctor/delivery"
// 	doctorRepo "kamarRS/features/doctor/repository"
// 	doctorService "kamarRS/features/doctor/service"

// 	hospitalDelivery "kamarRS/features/hospital/delivery"
// 	hospitalRepo "kamarRS/features/hospital/repository"
// 	hospitalService "kamarRS/features/hospital/service"

// 	hospitalStaffRepo "kamarRS/features/HospitalStaff/repository"
// 	hospitalStaffService "kamarRS/features/HospitalStaff/service"
// 	hospitalStaffDelivery "kamarRS/features/hospitalStaff/delivery"

// 	patientDelivery "kamarRS/features/patient/delivery"
// 	patientRepo "kamarRS/features/patient/repository"
// 	patientService "kamarRS/features/patient/service"

// 	policlinicDelivery "kamarRS/features/policlinic/delivery"
// 	policlinicRepo "kamarRS/features/policlinic/repository"
// 	policlinicService "kamarRS/features/policlinic/service"

// 	userDelivery "kamarRS/features/bed/delivery"
// 	userRepo "kamarRS/features/bed/repository"
// 	userService "kamarRS/features/bed/service"

// 	"github.com/labstack/echo/v4"
// 	"gorm.io/gorm"
// )

// func InitFactory(e *echo.Echo, db *gorm.DB) {

// 	bedRepoFactory := bedRepo.New(db)
// 	bedServiceFactory := bedService.New(bedRepoFactory)
// 	bedDelivery.New(bedServiceFactory, e)

// 	bedReservationRepoFactory := bedReservationRepo.New(db)
// 	bedReservationServiceFactory := bedReservationService.New(bedReservationRepoFactory)
// 	bedReservationDelivery.New(bedReservationServiceFactory, e)

// 	checkupReservationRepoFactory := checkupReservationRepo.New(db)
// 	checkupReservationServiceFactory := checkupReservationService.New(checkupReservationRepoFactory)
// 	checkupReservationDelivery.New(checkupReservationServiceFactory, e)

// 	dailyPracticeRepoFactory := dailyPracticeRepo.New(db)
// 	dailyPracticeServiceFactory := dailyPracticeService.New(dailyPracticeRepoFactory)
// 	dailyPracticeDelivery.New(dailyPracticeServiceFactory, e)

// 	doctorRepoFactory := doctorRepo.New(db)
// 	doctorServiceFactory := doctorService.New(doctorRepoFactory)
// 	doctorDelivery.New(doctorServiceFactory, e)

// 	hospitalRepoFactory := hospitalRepo.New(db)
// 	hospitalServiceFactory := hospitalService.New(hospitalRepoFactory)
// 	hospitalDelivery.New(hospitalServiceFactory, e)

// 	hospitalStaffRepoFactory := hospitalStaffRepo.New(db)
// 	hospitalStaffServiceFactory := hospitalStaffService.New(hospitalStaffRepoFactory)
// 	hospitalStaffDelivery.New(hospitalStaffServiceFactory, e)

// 	patientRepoFactory := patientRepo.New(db)
// 	patientServiceFactory := patientService.New(patientRepoFactory)
// 	patientDelivery.New(patientServiceFactory, e)

// 	policlinicRepoFactory := policlinicRepo.New(db)
// 	policlinicServiceFactory := policlinicService.New(policlinicRepoFactory)
// 	policlinicDelivery.New(policlinicServiceFactory, e)

// 	userRepoFactory := userRepo.New(db)
// 	userServiceFactory := userService.New(userRepoFactory)
// 	userDelivery.New(userServiceFactory, e)

// }
