package bed

type BedCore struct {
	ID              uint
	NamaTempatTidur string
	Ruangan         string
	Kelas           string
	Status          string
	HospitalID      uint
	Hospital        HospitalCore
}

type HospitalCore struct {
	ID   uint
	Nama string
}
