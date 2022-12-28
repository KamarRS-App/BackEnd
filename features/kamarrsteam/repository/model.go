package repository

import (
	"github.com/KamarRS-App/KamarRS-App/features/kamarrsteam"

	"gorm.io/gorm"
)

type KamarRsTeam struct {
	gorm.Model
	Email     string
	KataSandi string
	Peran     string
}

func FromKamarRsTeamCoretoModel(dataCore kamarrsteam.KamarRsTeamCore) KamarRsTeam {
	kamarrsteamGorm := KamarRsTeam{
		Email:     dataCore.Email,
		KataSandi: dataCore.KataSandi,
		Peran:     dataCore.Peran,
	}
	return kamarrsteamGorm
}

func (dataModel *KamarRsTeam) ToKamarRsTeamCore() kamarrsteam.KamarRsTeamCore {
	return kamarrsteam.KamarRsTeamCore{
		ID:        dataModel.ID,
		Email:     dataModel.Email,
		KataSandi: dataModel.KataSandi,
		Peran:     dataModel.Peran,
	}
}
