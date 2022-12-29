package hospital

type HospitalCore struct {
	ID                uint
	KodeRs            string
	Nama              string
	Foto              string
	Alamat            string
	Provinsi          string
	KabupatenKota     string
	Kecamatan         string
	KodePos           string
	NoTelpon          string
	Email             string
	KelasRs           string
	PemilikPengelola  string
	JumlahTempatTidur int
	StatusPenggunaan  string
	BiayaRegistrasi   int
}

type ServiceInterface interface {
	Create(input HospitalCore) (err error)
	GetAll() (data []HospitalCore, err error)
	GetById(id int) (data HospitalCore, err error)
	Update(input HospitalCore, id int) (err error)
	Delete(id int) (err error)
}

type RepositoryInterface interface {
	Create(input HospitalCore) (row int, err error)
	GetAll() (data []HospitalCore, err error)
	GetById(id int) (data HospitalCore, err error)
	Update(input HospitalCore, id int) (err error)
	Delete(id int) (row int, err error)
}
