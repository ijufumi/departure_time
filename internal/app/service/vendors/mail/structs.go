package mail

// Contents contains some data for sending mail
type Contents struct {
	ToAddress   string
	FromAddress string
	Subject     string
	Body        string
}
