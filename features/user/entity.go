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

}

type RepositoryInterface interface { // berkaitan database

	Create(input CoreUser) (err error)
}
