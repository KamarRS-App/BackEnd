package delivery

type HospitalResponse struct {
	ID                uint   `json:"id"`
	KodeRs            string `json:"kode_rs"`
	Nama              string `json:"nama"`
	Foto              string `json:"foto"`
	Alamat            string `json:"alamat"`
	Provinsi          string `json:"provinsi"`
	KabupatenKota     string `json:"kabupaten_kota"`
	Kecamatan         string `json:"kecamatan"`
	NoTelpon          string `json:"no_telpon"`
	Email             string `json:"email"`
	KelasRs           string `json:"kelas_rs"`
	PemilikPengelola  string `json:"pemilik_pengelola"`
	JumlahTempatTidur int    `json:"jumlah_tempat_tidur"`
	StatusPenggunaan  string `json:"status_penggunaan"`
	BiayaRegistrasi   int    `json:"biaya_registrasi"`
}
