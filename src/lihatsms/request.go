package lihatsms

type Requestsms struct {
	Number  string `json:"number"`
	Message string `json:"message"`
}

type Respons struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
