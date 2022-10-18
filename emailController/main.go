package emailcontroller

import (
	"fmt"

	"net/smtp"
)

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
} // Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

var fromMail string = "saiprasad@levantate.tech"
var password string = "password"

// Email data to email

func Send(mail []string) (string, bool) {
	to := mail
	smtpServer := smtpServer{host: "smtp.hostinger.com", port: "465"}
	message := []byte("This is a really unimaginative message, I know.")
	auth := smtp.PlainAuth("", fromMail, password, smtpServer.host)
	err := smtp.SendMail(smtpServer.Address(), auth, fromMail, to, message)
	if err != nil {
		errormsg := fmt.Sprintf("smtp error: %s", err)
		fmt.Printf(err.Error())
		return errormsg, true
	}
	return fmt.Sprintf("Email sent to %s!", to), false // return a confirmation message to the user, or utilize it for further processing
}

func ForgotPassword(mail []string) (string, bool) {
	to := mail
	smtpServer := smtpServer{host: "smtp.hostinger.com", port: "465"}
	message := []byte("This is a really unimaginative message, I know.")
	auth := smtp.PlainAuth("", fromMail, password, smtpServer.host)
	err := smtp.SendMail(smtpServer.Address(), auth, fromMail, to, message)
	if err != nil {
		errormsg := fmt.Sprintf("smtp error: %s", err)
		fmt.Printf(err.Error())
		return errormsg, true
	}
	return fmt.Sprintf("Email sent to %s!", to), false
}

func WelcomeMail(mail []string) (string, bool) {
	to := mail
	smtpServer := smtpServer{host: "smtp.hostinger.com", port: "465"}
	message := []byte("This is a really unimaginative message, I know.")
	auth := smtp.PlainAuth("", fromMail, password, smtpServer.host)
	err := smtp.SendMail(smtpServer.Address(), auth, fromMail, to, message)
	if err != nil {
		errormsg := fmt.Sprintf("smtp error: %s", err)
		fmt.Printf(err.Error())
		return errormsg, true
	}
	return fmt.Sprintf("Email sent to %s!", to), false
}

func OrderMail(mail []string) (string, bool) {
	to := mail
	smtpServer := smtpServer{host: "smtp.hostinger.com", port: "465"}
	message := []byte("This is a really unimaginative message, I know.")
	auth := smtp.PlainAuth("", fromMail, password, smtpServer.host)
	err := smtp.SendMail(smtpServer.Address(), auth, fromMail, to, message)
	if err != nil {
		errormsg := fmt.Sprintf("smtp error: %s", err)
		fmt.Printf(err.Error())
		return errormsg, true
	}
	return fmt.Sprintf("Email sent to %s!", to), false
}
