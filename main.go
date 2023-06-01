package email_service

import "github.com/ijufumi/email-service/cmd/app"

func main() {
	err := app.RunApp()
	if err != nil {
		panic(err)
	}
}
