package delivery

type AuthRequest struct {
	Email     string `json:"email" form:"email"`
	KataSandi string `json:"kata_sandi" form:"kata_sandi"`
}

type AuthRequestTeam struct {
	Email     string `json:"email" form:"email"`
	KataSandi string `json:"kata_sandi" form:"kata_sandi"`
}
