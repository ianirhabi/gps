package inputiuran

type RequestIuran struct {
	UserId       string `json:"userid"`
	Statusbulan  string `json:"status_bulan"`
	Statusminggu string `json:"status_minggu"`
	Statustahun  string `json:"status_tahun"`
	Statusbayar  string `json:"status_bayar"`
	Namaanggota  string `json:"nama_anggota"`
	Tanggalbayar string `json:"tanggal_bayar"`
}

type Respons struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
