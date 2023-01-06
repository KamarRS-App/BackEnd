test:
	go test ./features/auth/service... -coverprofile=cover.out 
	go test ./features/bed/service... -coverprofile=cover.out 
	go test ./features/bedReservation/service... -coverprofile=cover.out 
	go test ./features/checkupReservation/service... -coverprofile=cover.out 
	go test ./features/dailyPractice/service... -coverprofile=cover.out 
	go test ./features/doctor/service... -coverprofile=cover.out 
	go test ./features/hospital/service... -coverprofile=cover.out 
	go test ./features/hospitalstaff/service... -coverprofile=cover.out 
	go test ./features/kamarrsteam/service... -coverprofile=cover.out 
	go test ./features/patient/service... -coverprofile=cover.out 
	go test ./features/policlinic/service... -coverprofile=cover.out 
	go test ./features/user/service... -coverprofile=cover.out 
