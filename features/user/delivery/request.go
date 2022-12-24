package delivery

import (
	"kamarRS/features/user"
)

type RequestUser struct {
	ID        uint
	Nama      string
	Email     string
	Nokk      string
	Nik       string
	KataSandi string
	NoTelpon  string
}

func (req *RequestUser) reqToCore() user.CoreUser {
	return user.CoreUser{
		ID:        req.ID,
		Nama:      req.Nama,
		Email:     req.Email,
		Nokk:      req.Nokk,
		Nik:       req.Nik,
		KataSandi: req.KataSandi,
		NoTelpon:  req.NoTelpon,
	}

}
