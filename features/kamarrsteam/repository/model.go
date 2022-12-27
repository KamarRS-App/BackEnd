package repository

import (
	"github.com/KamarRS-App/features/kamarrsteam"

	"gorm.io/gorm"
)

type KamarRsTeam struct {
	gorm.Model
	Email    string
	Password string
	Peran    string
}

func FromKamarRsTeamCoretoModel(dataCore kamarrsteam.KamarRsTeamCore) KamarRsTeam {
	kamarrsteamGorm := KamarRsTeam{
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Peran:    dataCore.Peran,
	}
	return kamarrsteamGorm
}

func (dataModel *KamarRsTeam) ToKamarRsTeamCore() kamarrsteam.KamarRsTeamCore {
	return kamarrsteam.KamarRsTeamCore{
		ID:       dataModel.ID,
		Email:    dataModel.Email,
		Password: dataModel.Password,
		Peran:    dataModel.Peran,
	}
}
