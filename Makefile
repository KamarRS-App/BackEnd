test-user:
	go test ./features/user... -coverprofile=cover.out && go tool cover -html=cover.out

test-staff:
	go test ./features/hospitalstaff... -coverprofile=cover.out && go tool cover -html=cover.out

test-login:
	go test ./features/auth/service... -coverprofile=cover.out && go tool cover -html=cover.out

test-bed:
	go test ./features/bed... -coverprofile=cover.out && go tool cover -html=cover.out

test-bedres:
	go test ./features/bedReservation/service... -coverprofile=cover.out && go tool cover -html=cover.out
test-practice:
	go test ./features/dailyPractice/service... -coverprofile=cover.out && go tool cover -html=cover.out

