package kamarrsteam

type KamarRsTeamCore struct {
	ID       uint
	Email    string `valiidate:"required,email,unique"`
	Password string `validate:"required"`
	Peran    string `validate:"required"`
}

type ServiceInterface interface {
	Create(input KamarRsTeamCore) error
}

type RepositoryInterface interface {
	Create(input KamarRsTeamCore) error
	FindTeam(email string) (result KamarRsTeamCore, err error)
}
