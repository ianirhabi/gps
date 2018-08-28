package inputiuran

import (
	"fmt"
	"strings"

	orm "github.com/alfatih/beego/orm"
	"gps.com/gps/model"
)

func PostIuran(r RequestIuran) (d interface{}, e error) {
	var iuran []*model.Iuran_anggota
	o := orm.NewOrm()
	var data Respons
	if bulan, ero := Bulan(r.Statusbulan); ero == nil {
		s := strings.Split(r.Tanggalbayar, "-")
		tahun, bln, tanggal := s[0], s[1], s[2]
		if buln, ero := Bulan(bln); ero == nil {
			fmt.Println(tahun, buln, tanggal)
			tanggalbayar := tahun + "-" + buln + "-" + tanggal

			if f, d := o.Raw("SELECT * FROM iuran_anggota where user_id =? and status_bulan = ? and status_tahun=? and status_minggu =?", r.UserId, bulan, tahun, r.Statusminggu).QueryRows(&iuran); d == nil {
				// jika data tidak ada maka akan di masuk ke data base
				if f == 0 {
					iuran := model.Iuran_anggota{UserId: r.UserId,
						StatusBulan:  bulan,
						StatusMinggu: r.Statusminggu,
						StatusTahun:  r.Statustahun,
						StatusBayar:  r.Statusbayar,
						NamaAnggota:  r.Namaanggota,
						Tanggalbayar: tanggalbayar}
					if df, err := o.Insert(&iuran); err == nil {
						fmt.Println("debug berhasil = ", df, err)
						data.Data = iuran
						data.Status = "berhasil"
					} else {
						data.Status = "gagal"
						fmt.Println("debug gagal = ", df, err)
					}

				} else {
					data.Status = "sudahinputbaru"
					fmt.Println("debug gagal = ", f, d)
				}
			}
		}
	}
	return data, e
}

func IuranAnggotinfo(bln string, tahun string, userid string, minggu string) (d interface{}, e error) {

	o := orm.NewOrm()
	var data Respons
	var iuran []*model.Iuran_anggota
	if bulan, err := Bulan(bln); err == nil {
		if f, d := o.Raw("SELECT * FROM iuran_anggota where user_id =? and status_bulan = ? and status_tahun=? and status_minggu =?", userid, bulan, tahun, minggu).QueryRows(&iuran); d == nil {
			if f > 0 {
				fmt.Println("debug masuk disini ", f, d)
				data.Status = "berhasil"
				data.Data = iuran
			} else if f == 0 {
				fmt.Println("debug masuk = 0 ", f, d, bulan, userid, tahun, minggu)
				data.Tdk = f
				data.Pembayaran = "belum selesai membayar"
			}
		} else {
			fmt.Println("debug gagal = ", f, d)
			data.Status = "gagal coy"
		}
	}
	return data, e
}
func Bulan(bulan string) (b string, e error) {

	if bulan == "Jan" {
		bulan = "01"
	} else if bulan == "Feb" {
		bulan = "02"
	} else if bulan == "Mar" {
		bulan = "03"
	} else if bulan == "Apr" {
		bulan = "04"
	} else if bulan == "Mei" {
		bulan = "05"
	} else if bulan == "Jun" {
		bulan = "06"
	} else if bulan == "Jul" {
		bulan = "07"
	} else if bulan == "Agust" {
		bulan = "08"
	} else if bulan == "Sept" {
		bulan = "09"
	} else if bulan == "Okt" {
		bulan = "10"
	} else if bulan == "Nov" {
		bulan = "11"
	} else if bulan == "Des" {
		bulan = "12"
	} else {
		bulan = bulan
	}
	return bulan, e
}

func IuranAnggoDelete(bln string, tahun string, userid string, minggu string) (d interface{}, e error) {
	o := orm.NewOrm()
	var data Respons
	var iuran []*model.Iuran_anggota
	if bulan, err := Bulan(bln); err == nil {
		if f, d := o.Raw("delete FROM iuran_anggota where user_id =? and status_bulan = ? and status_tahun=? and status_minggu =?", userid, bulan, tahun, minggu).QueryRows(&iuran); d == nil {
			//jika terdapat data
			if f == 0 {
				data.Status = "berhasil"
				data.Data = "data behasil dihapus"
			} else {
				fmt.Println("debug gagal == ", f, d)
			}
		}
	}
	return data, e
}
