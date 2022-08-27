package request

// SendMail is request parameter for send mail
type SendMail struct {
	ToAddress   string `json:"to_address" binding:"required,email"`
	FromAddress string `json:"from_address" binding:"required,email"`
	Subject     string `json:"subject" binding:"required"`
	Body        string `json:"body" binding:"required"`
}
