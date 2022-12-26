package kamarrsteam

type KamarRsTeamCore struct {
	ID       uint
	Email    string `validate:"required"`
	Password string `validate:"required"`
	Peran    string `validate:"required"`
}
