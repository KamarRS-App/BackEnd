package delivery

import "github.com/KamarRS-App/KamarRS-App/features/kamarrsteam"

type KamarRsTeamRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Peran    string `json:"peran" form:"peran"`
}

func requestToCore(KamarRsTeamInput KamarRsTeamRequest) kamarrsteam.KamarRsTeamCore {
	kamarrsteamCoreData := kamarrsteam.KamarRsTeamCore{
		Email:    KamarRsTeamInput.Email,
		Password: KamarRsTeamInput.Password,
		Peran:    KamarRsTeamInput.Peran,
	}
	return kamarrsteamCoreData
}
