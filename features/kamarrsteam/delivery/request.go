package delivery

import "github.com/KamarRS-App/KamarRS-App/features/kamarrsteam"

type KamarRsTeamRequest struct {
	Email     string `json:"email" form:"email"`
	KataSandi string `json:"kata_sandi" form:"kata_sandi"`
	Peran     string `json:"peran" form:"peran"`
}

func requestToCore(KamarRsTeamInput KamarRsTeamRequest) kamarrsteam.KamarRsTeamCore {
	kamarrsteamCoreData := kamarrsteam.KamarRsTeamCore{
		Email:     KamarRsTeamInput.Email,
		KataSandi: KamarRsTeamInput.KataSandi,
		Peran:     KamarRsTeamInput.Peran,
	}
	return kamarrsteamCoreData
}
