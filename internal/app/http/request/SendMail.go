package request

type SendMail struct {
	ToAddress   string `json:"to_address"`
	FromAddress string `json:"from_address"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
}
