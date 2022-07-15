package seed

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
	"time"
	"math/rand"
	"golang.org/x/crypto/bcrypt"
	 "ecommerce/model"
	 "github.com/joho/godotenv"
	 "os"
)

func SeedUserDB(db *gorm.DB) {
	godotenv.Load(".env")
	sample_email := os.Getenv("SAMPLE_EMAIL")

	var data [] model.User

	psw, _ := bcrypt.GenerateFromPassword([]byte("mictest1212"), 10)

	for i:=1; i<=2; i++ {

		rand.Seed(time.Now().UnixNano())

		x := 10 + i

		userIdentity := rand.Intn(5000000000-x) + x

		data = append(data, model.User {UserID: uint(userIdentity), Username: faker.LastName(), Email: sample_email, Password: string(psw), Date: time.Now(), MemberShip: uint(i)})
	}

	db.Create(&data)
}

