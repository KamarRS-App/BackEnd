test-user:
	go test ./features/user/service... -coverprofile=cover.out && go tool cover -html=cover.out

test-staff:
	go test ./features/hospitalstaff/service... -coverprofile=cover.out && go tool cover -html=cover.out

test-login:
	go test ./features/auth/service... -coverprofile=cover.out && go tool cover -html=cover.out

test-bed:
	go test ./features/bed/service... -coverprofile=cover.out && go tool cover -html=cover.out

test-bedres:
	go test ./features/bedReservation/service... -coverprofile=cover.out && go tool cover -html=cover.out
test-practice:
	go test ./features/dailyPractice/service... -coverprofile=cover.out && go tool cover -html=cover.out


test-team:
	go test ./features/kamarrsteam/service... -coverprofile=cover.out && go tool cover -html=cover.out

test-reservation:
	go test ./features/checkupReservation/service... -coverprofile=cover.out && go tool cover -html=cover.out


test-hospital:
	go test ./features/hospital/service... -coverprofile=cover.out && go tool cover -html=cover.out

test-doctor:
	go test ./features/doctor/service... -coverprofile=cover.out && go tool cover -html=cover.out
test-policlinic:
	go test ./features/policlinic/service... -coverprofile=cover.out && go tool cover -html=cover.out