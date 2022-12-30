package service

import (
	bedreservation "github.com/KamarRS-App/KamarRS-App/features/bedReservation"
)

type bedReservationService struct {
	bedReservationRepository bedreservation.RepositoryInterface
}

func New(repo bedreservation.RepositoryInterface) bedreservation.ServiceInterface {
	return &bedReservationService{
		bedReservationRepository: repo,
	}
}

// GetRegistrations implements bedreservation.ServiceInterface
func (s *bedReservationService) GetRegistrations(page int, limit int, hospitalId int) (data []bedreservation.BedReservationCore, totalpage int, err error) {
	offset := (page - 1) * limit
	data, totalpage, err = s.bedReservationRepository.GetRegistrations(limit, offset, hospitalId)
	if err != nil {
		return nil, 0, err
	}
	return
}

// Create implements bedreservation.ServiceInterface
func (s *bedReservationService) Create(input bedreservation.BedReservationCore, userId uint) (data bedreservation.BedReservationCore, err error) {
	data, err = s.bedReservationRepository.Create(input, userId)
	if err != nil {
		return bedreservation.BedReservationCore{}, err
	}
	return data, nil
}

// GetPayment implements bedreservation.ServiceInterface
func (s *bedReservationService) GetPayment(kodeDaftar string) (data bedreservation.BedReservationCore, err error) {
	data, err = s.bedReservationRepository.GetPayment(kodeDaftar)
	if err != nil {
		return bedreservation.BedReservationCore{}, err
	}
	return data, nil
}

// CreatePayment implements bedreservation.ServiceInterface
func (s *bedReservationService) CreatePayment(input bedreservation.BedReservationCore) (data bedreservation.BedReservationCore, err error) {
	data, err = s.bedReservationRepository.CreatePayment(input)
	if err != nil {
		return bedreservation.BedReservationCore{}, err
	}
	return data, nil
}

// PaymentNotif implements bedreservation.ServiceInterface
func (s *bedReservationService) PaymentNotif(callback bedreservation.BedReservationCore) (err error) {
	err = s.bedReservationRepository.PaymentNotif(callback)
	if err != nil {
		return err
	}
	return nil
}
