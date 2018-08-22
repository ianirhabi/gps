package inputiuran

import (
	"fmt"

	"github.com/alfatih/irhabi/orm"
	"gps.com/gps/model"
)

func PostIuran(r RequestIuran) (d interface{}, e error) {
	o := orm.NewOrm()
	var data Respons
	iuran := model.Iuran_anggota{UserId: r.UserId,
		StatusBulan:  r.Statusbulan,
		StatusMinggu: r.Statusminggu,
		StatusTahun:  r.Statustahun,
		StatusBayar:  r.Statusbayar,
		NamaAnggota:  r.Namaanggota,
		Tanggalbayar: r.Tanggalbayar}
	if df, err := o.Insert(&iuran); err == nil {
		data.Data = iuran
		data.Status = "berhasil"
	} else {
		data.Status = "gagal"
		fmt.Println("debug gagal === ", df, err)
	}
	return data, e
}
