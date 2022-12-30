package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// var key string

func JWTMiddleware() echo.MiddlewareFunc {

	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey:    []byte(os.Getenv("SECRET_JWT")),
	})

}

// func CreateToken(userId int, role string) (string, error) {

// 	claims := jwt.MapClaims{}
// 	claims["authorized"] = true
// 	claims["userId"] = userId
// 	claims["role"] = role
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))

// }

// func ExtractToken(c echo.Context) (int, string) {

// 	user := c.Get("user").(*jwt.Token)

// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		userId := claims["userId"].(float64)
// 		role := claims["role"].(string)

// 		return int(userId), role
// 	}

// 	return 0, ""
// }

// // func ExtractTokenUserRole(e echo.Context) string {
// // 	user := e.Get("user").(*jwt.Token)
// // 	if user.Valid {
// // 		claims := user.Claims.(jwt.MapClaims)
// // 		role := claims["role"].(string)
// // 		return role
// // 	}
// // 	return ""
// // }

// func ExtractTokenUserName(e echo.Context) string {
// 	user := e.Get("user").(*jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		name := claims["name"].(string)
// 		return name
// 	}
// 	return ""
// }

func ExtractTokenUserName(e echo.Context) string {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return name
	}
	return ""
}

func CreateTokenTeam(teamId int, peran string, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = teamId
	claims["peran"] = peran
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func ExtractTokenTeamRole(e echo.Context) string {
	team := e.Get("user").(*jwt.Token)
	if team.Valid {
		claims := team.Claims.(jwt.MapClaims)
		peran := claims["peran"].(string)
		return peran
	}
	return ""
}
func ExtractTokenTeamId(e echo.Context) int {
	team := e.Get("user").(*jwt.Token)
	if team.Valid {
		claims := team.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}
