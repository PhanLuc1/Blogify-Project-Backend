package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/database"
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

	subject := "Your Instagram Code"

	tmpl, err := template.ParseFiles("src/auth/email_template.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	data := struct {
		Code string
	}{
		Code: otp,
	}

	err = tmpl.Execute(&body, data)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("Subject: %s\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s", subject, body.String())

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(message))

	if err != nil {
		return err
	}

	return nil
}

func AuthenticateCode(w http.ResponseWriter, r *http.Request) {
	var jsonData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email := jsonData["email"]
	code := jsonData["code"]

	emailStr := email.(string)
	codeStr := code.(string)
	if codeMap[emailStr] != codeStr {
		http.Error(w, "code is incorrect", http.StatusUnauthorized)
		return
	}
	delete(codeMap, emailStr)
	w.WriteHeader(200)
}
func SendPasswordEmail(email, newPassword string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	from := "lucphan1602@gmail.com"
	password := "ynex thsd sfwp gref"

	subject := "Reset Your Password"

	tmpl, err := template.ParseFiles("src/auth/reset_password_email_template.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	data := struct {
		Password string
	}{
		Password: newPassword,
	}

	err = tmpl.Execute(&body, data)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("Subject: %s\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s", subject, body.String())

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(message))

	if err != nil {
		return err
	}

	return nil
}
func GeneratePassword() string {
	const passwordLength = 12
	passwordChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/~"
	password := make([]byte, passwordLength)

	rand.Read(password)

	for i := range password {
		password[i] = passwordChars[int(password[i])%len(passwordChars)]
	}

	return string(password)
}
func GetResetPasswordCode(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var count int
	err = database.Client.QueryRow("SELECT COUNT(*) FROM user WHERE email = ?", user.Email).Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	newPassword := GeneratePassword()
	hashedPassword := HashPassword(newPassword)

	query := "UPDATE user SET password = ? WHERE email = ?"
	_, err = database.Client.Exec(query, hashedPassword, user.Email)
	if err != nil {
		http.Error(w, "Error updating password", http.StatusInternalServerError)
		return
	}

	err = SendPasswordEmail(user.Email, newPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.WriteHeader(200)
}
