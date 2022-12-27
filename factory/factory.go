package factory

import (
	// 	bedDelivery "github.com/KamarRS-App/features/bed/delivery"
	// 	bedRepo "github.com/KamarRS-App/features/bed/repository"
	// 	bedService "github.com/KamarRS-App/features/bed/service"

	// 	bedReservationDelivery "github.com/KamarRS-App/features/bedreservation/delivery"
	// 	bedReservationRepo "github.com/KamarRS-App/features/bedreservation/repository"
	// 	bedReservationService "github.com/KamarRS-App/features/bedreservation/service"

	// 	checkupReservationDelivery "github.com/KamarRS-App/features/checkupreservation/delivery"
	// 	checkupReservationRepo "github.com/KamarRS-App/features/checkupreservation/repository"
	// 	checkupReservationService "github.com/KamarRS-App/features/checkupreservation/service"

	// 	dailyPracticeDelivery "github.com/KamarRS-App/features/dailypractice/delivery"
	// 	dailyPracticeRepo "github.com/KamarRS-App/features/dailypractice/repository"
	// 	dailyPracticeService "github.com/KamarRS-App/features/dailypractice/service"

	doctorDelivery "github.com/KamarRS-App/features/doctor/delivery"
	doctorRepo "github.com/KamarRS-App/features/doctor/repository"
	doctorService "github.com/KamarRS-App/features/doctor/service"

	hospitalDelivery "github.com/KamarRS-App/features/hospital/delivery"
	hospitalRepo "github.com/KamarRS-App/features/hospital/repository"
	hospitalService "github.com/KamarRS-App/features/hospital/service"

	// 	hospitalStaffRepo "github.com/KamarRS-App/features/hospitalstaff/repository"
	// 	hospitalStaffService "github.com/KamarRS-App/features/hospitalstaff/service"
	// 	hospitalStaffDelivery "github.com/KamarRS-App/features/hospitalstaff/delivery"

	patientDelivery "github.com/KamarRS-App/features/patient/delivery"
	patientRepo "github.com/KamarRS-App/features/patient/repository"
	patientService "github.com/KamarRS-App/features/patient/service"

	// 	policlinicDelivery "github.com/KamarRS-App/features/policlinic/delivery"
	// 	policlinicRepo "github.com/KamarRS-App/features/policlinic/repository"
	// 	policlinicService "github.com/KamarRS-App/features/policlinic/service"

	userDelivery "github.com/KamarRS-App/features/user/delivery"
	userRepo "github.com/KamarRS-App/features/user/repository"
	userService "github.com/KamarRS-App/features/user/service"

	authDelivery "github.com/KamarRS-App/features/auth/delivery"
	authRepo "github.com/KamarRS-App/features/auth/repository"
	authService "github.com/KamarRS-App/features/auth/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	// bedRepoFactory := bedRepo.New(db)
	// bedServiceFactory := bedService.New(bedRepoFactory)
	// bedDelivery.New(bedServiceFactory, e)

	// bedReservationRepoFactory := bedReservationRepo.New(db)
	// bedReservationServiceFactory := bedReservationService.New(bedReservationRepoFactory)
	// bedReservationDelivery.New(bedReservationServiceFactory, e)

	// checkupReservationRepoFactory := checkupReservationRepo.New(db)
	// checkupReservationServiceFactory := checkupReservationService.New(checkupReservationRepoFactory)
	// checkupReservationDelivery.New(checkupReservationServiceFactory, e)

	// dailyPracticeRepoFactory := dailyPracticeRepo.New(db)
	// dailyPracticeServiceFactory := dailyPracticeService.New(dailyPracticeRepoFactory)
	// dailyPracticeDelivery.New(dailyPracticeServiceFactory, e)

	doctorRepoFactory := doctorRepo.New(db)
	doctorServiceFactory := doctorService.New(doctorRepoFactory)
	doctorDelivery.New(doctorServiceFactory, e)

	hospitalRepoFactory := hospitalRepo.New(db)
	hospitalServiceFactory := hospitalService.New(hospitalRepoFactory)
	hospitalDelivery.New(hospitalServiceFactory, e)

	// hospitalStaffRepoFactory := hospitalStaffRepo.New(db)
	// hospitalStaffServiceFactory := hospitalStaffService.New(hospitalStaffRepoFactory)
	// hospitalStaffDelivery.New(hospitalStaffServiceFactory, e)

	patientRepoFactory := patientRepo.New(db)
	patientServiceFactory := patientService.New(patientRepoFactory)
	patientDelivery.New(patientServiceFactory, e)

	// policlinicRepoFactory := policlinicRepo.New(db)
	// policlinicServiceFactory := policlinicService.New(policlinicRepoFactory)
	// policlinicDelivery.New(policlinicServiceFactory, e)

	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

}
