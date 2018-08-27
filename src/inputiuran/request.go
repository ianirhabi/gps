package inputiuran

type RequestIuran struct {
	UserId       int64  `json:"userid"`
	Statusbulan  string `json:"status_bulan"`
	Statusminggu string `json:"status_minggu"`
	Statustahun  string `json:"status_tahun"`
	Statusbayar  string `json:"status_bayar"`
	Namaanggota  string `json:"nama_anggota"`
	Tanggalbayar string `json:"tanggal_bayar"`
}

type Respons struct {
	Status     string      `json:"status"`
	Pembayaran string      `json:"status_bayar"`
	Data       interface{} `json:"data"`
	Tdk        int64       `json:"status_belum_bayar"`
}
