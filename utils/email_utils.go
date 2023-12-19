package utils

import (
	"log"
	"fmt"
	"os"
	"gopkg.in/gomail.v2"
	"github.com/joho/godotenv"
)

func init(){

	// Load environmental variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func SendVerificationEmail(toEmail, verificationCode string) error {
	email := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Email Verification Code")
	m.SetBody("text/html", fmt.Sprintf("Your verification code is: %s", verificationCode))

	
	d := gomail.NewDialer("smtp.gmail.com", 587, email, password)
	
	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending email: ", err)
		return err
	} else {
		log.Println("Email sent successfully")
	}

	return nil
}


