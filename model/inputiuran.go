package model

import (
	orm "github.com/alfatih/irhabi/orm"
)

func init() {
	orm.RegisterModel(new(Iuran_anggota))
}

type Iuran_anggota struct {
	Id           int64  `orm:"column(id);auto"json:"id"`
	UserId       string `orm:"column(user_id);size(100)"json:"userid"`
	StatusBulan  string `orm:"column(status_bulan);size(100)"json:"status_bulan"`
	StatusMinggu string `orm:"column(status_minggu);size(100)"json:"status_minggu"`
	StatusTahun  string `orm:"column(status_tahun);size(100)"json:"status_tahun"`
	StatusBayar  string `orm:"column(status_bayar);size(100)"json:"status_bayar"`
	NamaAnggota  string `orm:"column(nama_anggota);size(100)"json:"nama_anggota"`
	Tanggalbayar string `orm:"column(tanggal_bayar);size(100)"json:"tanggal_bayar"`
}
