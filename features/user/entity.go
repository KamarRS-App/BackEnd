package user

import (
	"time"
)

type CoreUser struct {
	ID        uint
	Nama      string
	Email     string `validate:"required"`
	Nokk      string
	Nik       string
	KataSandi string
	NoTelpon  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface { //sebagai contract yang dibuat di layer service

	Create(input CoreUser) (err error) // menambahkah data user berdasarkan data usercore
	Update(id int, input CoreUser) error
}

type RepositoryInterface interface { // berkaitan database

	Create(input CoreUser) (err error)
	Update(id int, input CoreUser) error
}
