package email

import (
	"log"
	"os"
	"recything/utils/constanta"
)

func SendVerificationEmail(emailAddress string, token string) {
	go func() {
		verificationLink := constanta.VERIFICATION_URL + token
		fileContent, err := os.ReadFile("utils/email/templates/account_registration.html")
		if err != nil {
			log.Printf("gagal membaca template email: %v", err)
			return
		}

		_, errEmail := SendEmailSMTP([]string{emailAddress}, string(fileContent), verificationLink)
		if errEmail != nil {
			log.Printf("gagal mengirim email verifikasi: %v", errEmail)
		}
	}()
}
