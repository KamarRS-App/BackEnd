package service

import (
	"errors"
	"strings"

	hospitalstaff "github.com/KamarRS-App/KamarRS-App/features/hospitalStaff"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type staffService struct {
	staffRepository hospitalstaff.RepositoryInterface //data repository dri entities
	validate        *validator.Validate
}

func New(repo hospitalstaff.RepositoryInterface) hospitalstaff.ServiceInterface { //dengan kembalian user.service
	return &staffService{
		staffRepository: repo,
		validate:        validator.New(),
	}
}

// Create implements hospitalstaff.ServiceInterface
func (service *staffService) Create(input hospitalstaff.HospitalStaffCore) (err error) {
	lower := strings.ToLower(input.Email)
	input.Email = lower
	input.Peran = "Admin"
	generatePass := Bcript(input.KataSandi)
	input.KataSandi = generatePass

	if validateERR := service.validate.Struct(input); validateERR != nil {
		return validateERR
	}

	errCreate := service.staffRepository.Create(input)
	if errCreate != nil {
		return errors.New(" Gagal membuat akun staff")
	}

	return nil
}

// DeleteById implements hospitalstaff.ServiceInterface
func (service *staffService) DeleteById(id int) error {
	err := service.staffRepository.DeleteById(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return err
}

// GetStaff implements hospitalstaff.ServiceInterface
func (service *staffService) GetStaff(id int) (data hospitalstaff.HospitalStaffCore, err error) {
	data, err = service.staffRepository.GetStaff(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}

// Update implements hospitalstaff.ServiceInterface
func (service *staffService) Update(id int, input hospitalstaff.HospitalStaffCore) error {
	errUpdate := service.staffRepository.Update(id, input)

	if errUpdate != nil {
		return errors.New("gagal mengupdate data , querry error")
	}

	return nil
}

func Bcript(y string) string {
	password := []byte(y)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)

}
