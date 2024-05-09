package auth

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/models"
)

var codeMap = make(map[string]string)

func GetCodeSendMail(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = SendEmail(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
    w.WriteHeader(200)
}
func GenerateOTP() string {
	const otpLength = 6
	otpChars := "0123456789"
	otp := make([]byte, otpLength)

	rand.Read(otp)

	for i := range otp {
		otp[i] = otpChars[int(otp[i])%len(otpChars)]
	}

	return string(otp)
}
func SendEmail(email string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	from := "lucphan1602@gmail.com"
	password := "ynex thsd sfwp gref"

	otp := GenerateOTP()
	codeMap[email] = otp

	subject := "Mã Xác Thực"
	body := fmt.Sprintf("Mã xác thực của bạn là: %s", otp)
	message := fmt.Sprintf("Subject: %s\n\n%s", subject, body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(message))

	if err != nil {
		return err
	}

	return nil
}
