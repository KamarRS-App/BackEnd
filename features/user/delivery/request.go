package delivery

import (
	"github.com/KamarRS-App/KamarRS-App/features/user"
)

type RequestUser struct {
	Nama      string `json:"nama" form:"nama"`
	Email     string `json:"email" form:"email"`
	Nokk      string `json:"no_kk" form:"no_kk"`
	Nik       string `json:"nik" form:"nik"`
	KataSandi string `json:"kata_sandi" form:"kata_sandi"`
	NoTelpon  string `json:"no_telpon" form:"no_telpon"`
}

func (req *RequestUser) reqToCore() user.CoreUser {
	return user.CoreUser{

		Nama:      req.Nama,
		Email:     req.Email,
		Nokk:      req.Nokk,
		Nik:       req.Nik,
		KataSandi: req.KataSandi,
		NoTelpon:  req.NoTelpon,
	}
}
