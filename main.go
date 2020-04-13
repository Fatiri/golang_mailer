package main
import (

	"github.com/golang_mailer/helper"
)
func main()  {
    // sender := helper.NewSender("ilhamfatiriblog@gmail.com", "fatiriilham@01031995")
    user := map[int]string{}
    user[1] = "ilhamfatiriblog@gmail.com"
    user[2] =  "fatiriilham@01031995"
	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := []string{"ilham.fatiri@thelionparcel.com"}

	Subject := "Testing HTLML Email from golang"
    message := "Hello Guys"
	bodyMessage := helper.WriteEmail(Receiver,"text/html", Subject, message, user)

    helper.SendMail(Receiver, Subject, bodyMessage, user)
}