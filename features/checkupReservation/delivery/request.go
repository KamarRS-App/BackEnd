package delivery

import "github.com/KamarRS-App/features/checkupreservation"

type CheckupReservationRequest struct {
	PatientID  uint `json:"patient_id" form:"patient_id"`
	PracticeID uint `json:"practice_id" form:"practice_id"`
}

func (req *CheckupReservationRequest) reqToCore() checkupreservation.CheckupReservationCore {
	return checkupreservation.CheckupReservationCore{
		PatientID:  req.PatientID,
		PracticeID: req.PracticeID,
	}
}
