package controller

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"ecommerce/model"
	"ecommerce/auth"
	"ecommerce/db"
	"ecommerce/queue"
	"math/rand"
	"time"
	"errors"
	"github.com/dirkaholic/kyoo/job"
	"fmt"
)

type params struct {
	Username string `json: "username"`
	Email string `json: "email"`
	Password string `json: "password"`
}

type claims struct {
	Username string `json: "username`
	jwt.StandardClaims
}

//sub func to send verification email, used on register and login
func SendVerificationMail(info model.User, encrypted string)(error){
	godotenv.Load(".env")
	host := os.Getenv("HOST")

	verifyLink := host + "verify/" + encrypted

	subject := "Subject: " + "Sample Verification Email" + "!\n"

	content := fmt.Sprintf("<b>Click here to %s to verify</b>", verifyLink)

	//send to queue and process
	queue.Queue.Submit(&job.FuncExecutorJob{Func: func() error {
		return auth.SendSMTPMail([]string{info.Email}, subject, content)
	}})

	return nil
}

//login
func Login(c echo.Context)(err error) {
	godotenv.Load(".env")
	aes_secret := os.Getenv("AES_SECRET")

	current := time.Now()
	var user model.User
	var query params

	list := make([]string, 2)

	list = append(list, "dutchvanderlane", "tommy")

	if err:= c.Bind(&query); err != nil {
		return err
	}

	if ((query.Username == "" && query.Email == "") || query.Password == "") {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
	}

	//login using username or email
	if (query.Username != "") {
		db.DB.Where("username = ?", query.Username).Preload("Token").First(&user)
	} else {
		db.DB.Where("email = ?", query.Email).Preload("Token").First(&user)
	}

	if (!auth.CheckPsw(query.Password, user.Password)){
		panic(echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials"))
	}

	if user.Token.UserID > 1 {

		//renew token
		if (current.After(user.Token.Expires_at)){
			//delete expired token
			db.DB.Delete(&model.Token{}, user.Token.TokenID)

			//create new token
			encrypted := auth.AESEncrypt(user.Username, aes_secret)

			if encrypted == "500" {
				return echo.NewHTTPError(http.StatusInternalServerError, "Encryption Error") 
			}

			tok := model.Token{UserID: uint(user.UserID), AuthToken: encrypted, Generated_at: time.Now(), Expires_at: time.Now().Add(time.Minute * time.Duration(5))}
			db.DB.Create(&tok)

			//REsend verification mail
			SendVerificationMail(user, encrypted)

			return c.JSON(http.StatusOK, "Verification email has been sent. Please verify first before login")
		}

		return c.JSON(http.StatusOK, "Please verify your email first before login")
	}

	access, err := auth.GenerateAccessToken(user.Username)

	if (err != nil) {
		panic(echo.NewHTTPError(http.StatusUnprocessableEntity, "Please relogin"))
	}

	compromised := auth.SetCookie(access)

	c.SetCookie(compromised)

	return c.JSON(http.StatusOK, "Login Success")

}

//register
func Register(c echo.Context)(err error) {
	godotenv.Load(".env")
	aes_secret := os.Getenv("AES_SECRET")

	var users model.User
	var query params

	if err:= c.Bind(&query); err != nil {
		return err
	}

	//Prevent duplicate email
	emailFault := db.DB.Where("email = ?", query.Email).First(&users).Error 

	if !errors.Is(emailFault, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusBadRequest, "Please try another email")
	}

	//Prevent duplicate username
	usernameFault := db.DB.Where("username = ?", query.Username).First(&users).Error 

	if !errors.Is(usernameFault, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusBadRequest, "Please try another username")
	}

	rand.Seed(time.Now().UnixNano())

	userIdentity := rand.Intn(5000000000-10) + 10

	//Create User
	psw, _ := bcrypt.GenerateFromPassword([]byte(query.Password), 10)

	info := model.User{UserID: uint(userIdentity), Username: query.Username, Email: query.Email, Password: string(psw), Date: time.Now(), MemberShip: uint(2)}

	//Create Token
	encrypted := auth.AESEncrypt(info.Username, aes_secret)

	if encrypted == "500" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Encryption Error") 
	}

	tok := model.Token{UserID: uint(userIdentity), AuthToken: encrypted, Generated_at: time.Now(), Expires_at: time.Now().Add(time.Minute * time.Duration(5))}

	//create user
	db.DB.Create(&info)

	//create token
	db.DB.Create(&tok)

	//send verification mail
	SendVerificationMail(info, encrypted)

	return c.JSON(http.StatusOK, "Successfully Registered")
}

//verify through email link
func Verify(c echo.Context)(err error) {
	current := time.Now()

	var tok model.Token

	code := c.Param("id")

	//check token exist
	db.DB.Where("auth_token = ?", code).First(&tok)

	if tok.UserID == 0 {
		return echo.NewHTTPError(http.StatusInternalServerError, "No token found") 
	}

	//check token validity
	if (current.After(tok.Expires_at)){
		return echo.NewHTTPError(http.StatusUnauthorized, "Token Expired. Please relogin to retrieve token")
	}

	db.DB.Delete(&tok)

	return c.JSON(http.StatusOK, "Successful verify")
}