package auth

import (
	"errors"
	"net/http"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
	"net/smtp"
	"ecommerce/config"
)

var envData = config.EnvConfig()

type claims struct {
	Username string `json: "username`
	jwt.StandardClaims
}

//on login func
func CheckPsw(saved, hash string) bool {
	status := bcrypt.CompareHashAndPassword([]byte(hash), []byte(saved))
	return status == nil
} 

//on login func
func GenerateAccessToken(user string)(string, error) {

	claimData:= &claims {
		Username: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimData)

	tokenData, err := token.SignedString([]byte((*envData).SECRET_KEY))

	return tokenData, err
}

//on login func
func SetCookie(tokenData string)(*http.Cookie){
	cookie := new(http.Cookie)

	cookie.Name = "token";
	cookie.Value = tokenData;
	cookie.Expires = time.Now().Add(2 * time.Minute);
	cookie.Path = "/api/admin";

	return cookie;
}

//on login func
func VerifyCookie(cookie *http.Cookie) (*claims, error) {

	claimData := new(claims)

	_, err := jwt.ParseWithClaims(cookie.Value, claimData, func(token *jwt.Token) (interface{}, error) {
		return []byte((*envData).SECRET_KEY), nil
	})

	return claimData, err
}

//on register or renew token
func AESEncrypt(data string, secret string)(string){
	txt := []byte(data)
	phase := []byte(secret)

	as, err := aes.NewCipher(phase)

	if err != nil {
		panic("500")
	}

	gcm, err := cipher.NewGCM(as)

	if err != nil {
		panic("500")
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic("500")
	}

	encryptedMsg := gcm.Seal(nonce, nonce, txt, nil)

	return fmt.Sprintf("%x", encryptedMsg)
}

//For future use
func AESDecrypt(data string, secret string)(string) {
	txt := []byte(data)
	phase := []byte(secret)

	as, err := aes.NewCipher(phase)

	if err != nil {
		panic("500")
	}

	gcm, err := cipher.NewGCM(as)

	if err != nil {
		panic("500")
	}

	nonce := make([]byte, gcm.NonceSize())

	nonceSize := gcm.NonceSize()
	if len(txt) < nonceSize {
        panic("500")
    }

	nonce, text := txt[:nonceSize], txt[nonceSize:]

    words, err := gcm.Open(nil, nonce, text, nil)

    if err != nil {
        panic("Fail to decrypt")
    }

	return fmt.Sprintf("%x", words)
}

//Send gmail func
func SendSMTPMail(to []string, subject string, content string)(error){

	auth := smtp.PlainAuth("", (*envData).MAIL_FROM, (*envData).MAIL_PSW, (*envData).MAIL_HOST)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msg := []byte(subject + mime + "\n" + content)
	address := fmt.Sprintf("%s:%s", (*envData).MAIL_HOST, (*envData).MAIL_PORT)

	if err := smtp.SendMail(address, auth, (*envData).MAIL_FROM, to, msg); err != nil {
			return errors.New("Something wrong on email sent")
	}

	return nil
}





