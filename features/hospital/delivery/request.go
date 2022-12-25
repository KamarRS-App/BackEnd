package delivery

type HospitalRequest struct {
	KodeRs            string `json:"kode_rs" form:"kode_rs"`
	Nama              string `json:"nama" form:"nama"`
	Foto              string `json:"foto" form:"foto"`
	Alamat            string `json:"alamat" form:"alamat"`
	Provinsi          string `json:"provinsi" form:"provinsi"`
	KabupatenKota     string `json:"kabupaten_kota" form:"kabupaten_kota"`
	Kecamatan         string `json:"kecamatan" form:"kecamatan"`
	NoTelpon          string `json:"no_telpon" form:"no_telpon"`
	Email             string `json:"email" form:"email"`
	KelasRs           string `json:"kelas_rs" form:"kelas_rs"`
	PemilikPengelola  string `json:"pemilik_pengelola" form:"pemilik_pengelola"`
	JumlahTempatTidur int    `json:"jumlah_tempat_tidur" form:"jumlah_tempat_tidur"`
	StatusPenggunaan  string `json:"status_penggunaan" form:"status_penggunaan"`
	BiayaRegistrasi   int    `json:"biaya_registrasi" form:"biaya_registrasi"`
}
