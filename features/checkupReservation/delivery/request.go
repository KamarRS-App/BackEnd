package delivery

import checkupreservation "github.com/KamarRS-App/KamarRS-App/features/checkupReservation"

type CheckupReservationRequest struct {
	PatientID  uint   `json:"patient_id" form:"patient_id"`
	PracticeID uint   `json:"practice_id" form:"practice_id"`
	NamaDokter string `json:"nama_dokter" form:"nama_dokter"`
	NoAntrian  string `json:"no_antrian" form:"no_antrian"`
}

func (req *CheckupReservationRequest) reqToCore() checkupreservation.CheckupReservationCore {
	return checkupreservation.CheckupReservationCore{
		PatientID:  req.PatientID,
		PracticeID: req.PracticeID,
		NamaDokter: req.NamaDokter,
		NoAntrian:  req.NoAntrian,
	}
}
