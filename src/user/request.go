package user

type Register struct {
	Username string `json:"username"`
	Nama     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Imei     string `json:"imei"`
	Lat      string `json:"lat"`
	Lont     string `json:"lont"`
	status   string `json:"status"`
	Time     string `json:"time"`
	Tanggal  string `json:"tanggal"`
}

type Status struct {
	Aktif string `json:"aktif"`
}

type Respons struct {
	Status string      `json: "Status"`
	Data   interface{} `json:"data"`
}
type RequestLogin struct {
	Username string `json: "username"`
	Password string `json: "password"`
}
