package model

import orm "github.com/alfatih/irhabi/orm"

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id       int64  `orm:"column(id);auto"json:"id"`
	Username string `orm:"column(username);size(100)"json:"username"`
	Password string `orm:"column(password);size(100)"json:"password"`
	Imei     string `orm:"column(imei);size(100)"json:"imei"`
	Token    string `orm:"column(token);size(100)"json:"token"`
	Long     string `orm:"column(lont);size(100)"json:"lont"`
	Lat      string `orm:"column(lat);size(100)"json:"lat"`
	Status   string `orm:"column(status);size(100)"json:"status"`
	Time     string `orm:"column(time);size(100)"json:"time"`
	Date     string `orm:"column(tanggal);size(100)"json:"tanggal"`
	Nama     string `orm:"column(name);size(100)"json:"name"`
	Usergrup string `orm:"column(usergrup);size(100)"json:"usergrup"`
}
