package hospital

type HospitalCore struct {
	ID                uint
	KodeRs            string `validate:"required"`
	Nama              string `validate:"required"`
	Foto              string `validate:"required"`
	Alamat            string `validate:"required"`
	Provinsi          string `validate:"required"`
	KabupatenKota     string `validate:"required"`
	Kecamatan         string `validate:"required"`
	KodePos           string `validate:"required"`
	NoTelpon          string `validate:"required"`
	Email             string `valiidate:"required,email,unique"`
	KelasRs           string `validate:"required"`
	PemilikPengelola  string `validate:"required"`
	JumlahTempatTidur int    `validate:"required"`
	StatusPenggunaan  string `validate:"required"`
	BiayaRegistrasi   int    `validate:"required"`
}

type ServiceInterface interface {
	Create(input HospitalCore) (err error)
	GetAll(provinsi, kabKota, nama string, page, limit int) (data []HospitalCore, totalPage int, err error)
	GetById(id int) (data HospitalCore, err error)
	Update(input HospitalCore, id int) (err error)
	Delete(id int) (err error)
}

type RepositoryInterface interface {
	Create(input HospitalCore) (err error)
	GetAll(provinsi, kabKota, nama string, limit, offset int) (data []HospitalCore, totalPage int, err error)
	GetById(id int) (data HospitalCore, err error)
	Update(input HospitalCore, id int) (err error)
	Delete(id int) (row int, err error)
}
