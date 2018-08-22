package user

type Register struct {
	Username string `json:"username"`
	Nama     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"tokenfirebase"`
	Imei     string `json:"imei"`
	lat      string `json:"lat"`
	lont     string `json:"long"`
	status   string `json:"status"`
	Time     string `json:"time"`
	Tanggal  string `json:"tanggal"`
}

type Respons struct {
	Status string      `json: "Status"`
	Data   interface{} `json:"data"`
}
type RequestLogin struct {
	Username string `json: "username"`
	Password string `json: "password"`
}
