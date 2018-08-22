package user

import (
	"fmt"

	orm "github.com/alfatih/irhabi/orm"
	"gps.com/gps/model"
)

func PostRegis(r Register, imei string) (d interface{}, e error) {
	o := orm.NewOrm()
	var data Respons
	var validasi []*model.User

	if f, d := o.Raw("SELECT * FROM user where imei =?", imei).QueryRows(&validasi); d == nil {
		fmt.Println("imeinya ", f, d, imei)
		//Conversi Interface
		// var token = map[string]interface{}{
		// 	"nama": *validasi[0],
		// }

		//fmt.Println(validasi.Id)
		// s := strings.Split(*validasi[0], "{")
		// id := s[0]
		// fmt.Println("id nya ", id)
		if f == 0 {
			user := model.User{Username: r.Username,
				Password: r.Password,
				Imei:     imei,
				Token:    r.Token,
				Long:     r.lont,
				Lat:      r.lat,
				Status:   "pending",
				Time:     r.Time,
				Date:     r.Tanggal,
				Nama:     r.Nama}
			if df, e := o.Insert(&user); e == nil {
				data.Status = "berhasil insert"
				data.Data = &user
			} else {
				fmt.Println("debug kirim regis gagal === ", df, e)
				data.Status = "gagal"
			}
		} else {
			data.Status = "Anda sudah terdaftar"
			fmt.Println("data sudah ada")
		}
	} else {
		fmt.Println("kesalahan pada mysql")
	}
	return data, e
}

func Postlogin(r RequestLogin) (d interface{}, e error) {
	o := orm.NewOrm()
	var data Respons
	var uon []*model.User
	var u model.User
	if _, d := o.Raw("SELECT * from user").QueryRows(&uon); d == nil {
		for _, row := range uon {
			if row.Username == r.Username && row.Password == r.Password {
				if x := o.Raw("select * from user where username=?", r.Username).QueryRow(&u); x == nil {
					data.Status = "berhasil"
					data.Data = u
				}
			}
		}
	}
	return data, e
}

func Get() (d interface{}, e error) {
	var uon []*model.User
	o := orm.NewOrm()
	var data Respons

	if a, d := o.Raw("select * from user").QueryRows(&uon); d == nil {
		data.Data = uon
		data.Status = "berhasil"
	} else {
		data.Status = "gagal ambil data"
		fmt.Println("debug gagal ===== ", a, d)
	}
	return data, e
}
