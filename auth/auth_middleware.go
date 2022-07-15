package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func AuthVerification(next echo.HandlerFunc)(echo.HandlerFunc){

	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")

		if err != nil {
			if err == http.ErrNoCookie {
				panic(echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials"))
			}

			panic(echo.NewHTTPError(http.StatusUnprocessableEntity, "Please relogin"))
		}

		claim := VerifyCookie(cookie)

		//will only renew within 1 minutes before token expires
		if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()) > 60*time.Second {
			return next(c)
		}

		access, err := GenerateAccessToken(claim.Username)

		if (err != nil) {
			panic(echo.NewHTTPError(http.StatusUnprocessableEntity, "Please relogin"))
		}
	
		compromised := SetCookie(access)
	
		c.SetCookie(compromised)
		
		return next(c)

	}

}

