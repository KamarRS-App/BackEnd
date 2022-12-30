package service

import bedreservation "github.com/KamarRS-App/KamarRS-App/features/bedReservation"

type bedReservationService struct {
	bedReservationRepository bedreservation.RepositoryInterface
}

func New(repo bedreservation.RepositoryInterface) bedreservation.ServiceInterface {
	return &bedReservationService{
		bedReservationRepository: repo,
	}
}

// Create implements bedreservation.ServiceInterface
func (s *bedReservationService) Create(input bedreservation.BedReservationCore) (data bedreservation.BedReservationCore, err error) {
	data, err = s.bedReservationRepository.Create(input)
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
