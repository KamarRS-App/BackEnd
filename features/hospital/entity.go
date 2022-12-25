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
	NoTelpon          string
	Email             string
	KelasRs           string
	PemilikPengelola  string
	JumlahTempatTidur int
	StatusPenggunaan  string
	BiayaRegistrasi   int
}
