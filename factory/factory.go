package factory

import (

	// 	bedDelivery "github.com/KamarRS-App/KamarRS-App/features/bed/delivery"
	// 	bedRepo "github.com/KamarRS-App/KamarRS-App/features/bed/repository"
	// 	bedService "github.com/KamarRS-App/KamarRS-App/features/bed/service"

	// 	bedReservationDelivery "github.com/KamarRS-App/KamarRS-App/features/bedreservation/delivery"
	// 	bedReservationRepo "github.com/KamarRS-App/KamarRS-App/features/bedreservation/repository"
	// 	bedReservationService "github.com/KamarRS-App/KamarRS-App/features/bedreservation/service"

	dailyPracticeDelivery "github.com/KamarRS-App/KamarRS-App/features/dailyPractice/delivery"
	dailyPracticeRepo "github.com/KamarRS-App/KamarRS-App/features/dailyPractice/repository"
	dailyPracticeService "github.com/KamarRS-App/KamarRS-App/features/dailyPractice/service"

	doctorDelivery "github.com/KamarRS-App/KamarRS-App/features/doctor/delivery"
	doctorRepo "github.com/KamarRS-App/KamarRS-App/features/doctor/repository"
	doctorService "github.com/KamarRS-App/KamarRS-App/features/doctor/service"

	hospitalDelivery "github.com/KamarRS-App/KamarRS-App/features/hospital/delivery"
	hospitalRepo "github.com/KamarRS-App/KamarRS-App/features/hospital/repository"
	hospitalService "github.com/KamarRS-App/KamarRS-App/features/hospital/service"

	// 	hospitalStaffRepo "github.com/KamarRS-App/KamarRS-App/features/hospitalstaff/repository"
	// 	hospitalStaffService "github.com/KamarRS-App/KamarRS-App/features/hospitalstaff/service"
	// 	hospitalStaffDelivery "github.com/KamarRS-App/KamarRS-App/features/hospitalstaff/delivery"

	patientDelivery "github.com/KamarRS-App/KamarRS-App/features/patient/delivery"
	patientRepo "github.com/KamarRS-App/KamarRS-App/features/patient/repository"
	patientService "github.com/KamarRS-App/KamarRS-App/features/patient/service"

	"github.com/KamarRS-App/KamarRS-App/features/hospitalstaff/delivery"
	"github.com/KamarRS-App/KamarRS-App/features/hospitalstaff/repository"
	"github.com/KamarRS-App/KamarRS-App/features/hospitalstaff/service"

	policlinicDelivery "github.com/KamarRS-App/KamarRS-App/features/policlinic/delivery"
	policlinicRepo "github.com/KamarRS-App/KamarRS-App/features/policlinic/repository"
	policlinicService "github.com/KamarRS-App/KamarRS-App/features/policlinic/service"

	userDelivery "github.com/KamarRS-App/KamarRS-App/features/user/delivery"
	userRepo "github.com/KamarRS-App/KamarRS-App/features/user/repository"
	userService "github.com/KamarRS-App/KamarRS-App/features/user/service"

	authDelivery "github.com/KamarRS-App/KamarRS-App/features/auth/delivery"
	authRepo "github.com/KamarRS-App/KamarRS-App/features/auth/repository"
	authService "github.com/KamarRS-App/KamarRS-App/features/auth/service"

	kamarRsTeamDelivery "github.com/KamarRS-App/KamarRS-App/features/kamarrsteam/delivery"
	kamarRsTeamRepo "github.com/KamarRS-App/KamarRS-App/features/kamarrsteam/repository"
	kamarRsTeamService "github.com/KamarRS-App/KamarRS-App/features/kamarrsteam/service"

	checkupReservationDelivery "github.com/KamarRS-App/KamarRS-App/features/checkupReservation/delivery"
	checkupReservationRepo "github.com/KamarRS-App/KamarRS-App/features/checkupReservation/repository"
	checkupReservationService "github.com/KamarRS-App/KamarRS-App/features/checkupReservation/service"

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

	checkupReservationRepoFactory := checkupReservationRepo.New(db)
	checkupReservationServiceFactory := checkupReservationService.New(checkupReservationRepoFactory)
	checkupReservationDelivery.New(checkupReservationServiceFactory, e)

	dailyPracticeRepoFactory := dailyPracticeRepo.New(db)
	dailyPracticeServiceFactory := dailyPracticeService.New(dailyPracticeRepoFactory)
	dailyPracticeDelivery.New(dailyPracticeServiceFactory, e)

	doctorRepoFactory := doctorRepo.New(db)
	doctorServiceFactory := doctorService.New(doctorRepoFactory)
	doctorDelivery.New(doctorServiceFactory, e)

	hospitalRepoFactory := hospitalRepo.New(db)
	hospitalServiceFactory := hospitalService.New(hospitalRepoFactory)
	hospitalDelivery.New(hospitalServiceFactory, e)

	hospitalStaffRepoFactory := repository.New(db)
	hospitalStaffServiceFactory := service.New(hospitalStaffRepoFactory)
	delivery.New(hospitalStaffServiceFactory, e)

	patientRepoFactory := patientRepo.New(db)
	patientServiceFactory := patientService.New(patientRepoFactory)
	patientDelivery.New(patientServiceFactory, e)

	policlinicRepoFactory := policlinicRepo.New(db)
	policlinicServiceFactory := policlinicService.New(policlinicRepoFactory)
	policlinicDelivery.New(policlinicServiceFactory, e)

	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

	kamarRsTeamRepoFactory := kamarRsTeamRepo.New(db)
	kamarRsTeamServiceFactory := kamarRsTeamService.New(kamarRsTeamRepoFactory)
	kamarRsTeamDelivery.New(kamarRsTeamServiceFactory, e)

}
