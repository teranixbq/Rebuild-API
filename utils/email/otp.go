package email

import (
	"log"
	"os"
	"strings"
)

func SendOTPEmail(emailAddress string, otp string) {
	go func() {
		// Buka file template email.
		filePath := "utils/email/templates/otp.html"
		file, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("gagal membaca template email: %v", err)
			return
		}
		emailTemplate := string(file)

		emailContent := strings.Replace(emailTemplate, "{{.Otp}}", otp, -1)

		_, errEmail := SendEmailSMTPForOTP([]string{emailAddress}, emailContent, otp)
		if errEmail != nil {
			log.Printf("gagal mengirim otp: %v", errEmail)
		}
	}()
}
