package lihatsms

import (
	"fmt"

	orm "github.com/alfatih/beego/orm"
	"gps.com/gps/model"
)

func Kirimsms(a Requestsms, username string) (res interface{}, e error) {
	o := orm.NewOrm()
	var data Respons

	sms := model.Smsuser{Userid: username,
		Number:  a.Number,
		Message: a.Message}
	if df, e := o.Insert(&sms); e == nil {
		data.Status = "berhasil"
		data.Data = &sms
	} else {
		fmt.Println("debug kirim sms gagal === ", df, e)
		data.Status = "gagal"
	}
	return data, e
}
