package auth

import (
	"errors"
	"net/http"
	"os"
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"github.com/joho/godotenv"
	"time"
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
	"net/smtp"
)

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
	godotenv.Load(".env")
	secret_key := os.Getenv("SECRET_KEY")

	claimData:= &claims {
		Username: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimData)

	tokenData, err := token.SignedString([]byte(secret_key))

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
func VerifyCookie(cookie *http.Cookie) *claims {
	godotenv.Load(".env")
	secret_key := os.Getenv("SECRET_KEY")

	claimData := new(claims)

	token, err := jwt.ParseWithClaims(cookie.Value, claimData, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret_key), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			panic(echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials"))
		}

		panic(echo.NewHTTPError(http.StatusUnprocessableEntity, "Please relogin"))
	}

	if !token.Valid {
		panic(echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials"))
	}

	return claimData
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
	godotenv.Load(".env")
	mail_host := os.Getenv("MAIL_HOST")
	mail_from := os.Getenv("MAIL_FROM")
	mail_psw := os.Getenv("MAIL_PSW")
	mail_port := os.Getenv("MAIL_PORT")

	auth := smtp.PlainAuth("", mail_from, mail_psw, mail_host)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msg := []byte(subject + mime + "\n" + content)
	address := fmt.Sprintf("%s:%s", mail_host, mail_port)

	if err := smtp.SendMail(address, auth, mail_from, to, msg); err != nil {
			return errors.New("Something wrong on email sent")
	}

	return nil
}





