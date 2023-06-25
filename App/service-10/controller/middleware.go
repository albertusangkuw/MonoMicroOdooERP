package controller

import (
	"errors"
	"net/http"
	"service-10/dto"
	"service-10/utils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTHandler(next echo.HandlerFunc) echo.HandlerFunc {
	var jwtSecret = []byte(utils.EnvVar("JWT_SECRET_KEY"))
	return func(c echo.Context) error {
		cookie, err := c.Cookie(utils.EnvVar("JWT_COOKIE_NAME"))
		if err != nil {
			return c.String(http.StatusUnauthorized, "No JWT in Cookie")

		}
		token, err := jwt.ParseWithClaims(cookie.Value, &dto.JWTData{},
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unable to parse JWT")
				}
				return jwtSecret, nil
			})

		claims, ok := token.Claims.(*dto.JWTData)
		if !ok || !token.Valid {
			return c.String(http.StatusUnauthorized, "Invalid JWT")
		}
		c.Set("jwt_data", claims)
		return next(c)
	}
}

var LoggerHandler = middleware.LoggerWithConfig(middleware.LoggerConfig{
	Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
		`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
	CustomTimeFormat: "2006-01-02 15:04:05.00000",
})
