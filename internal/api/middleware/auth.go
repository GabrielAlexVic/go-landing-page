package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"go-landing-page/internal/api"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token de autorização ausente")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Formato de token inválido. Use 'Bearer <token>'")
		}

		tokenString := parts[1]
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.ParseWithClaims(tokenString, &api.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token inválido ou expirado")
		}

		claims, ok := token.Claims.(*api.CustomClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Claims de token inválidos")
		}

		c.Set("user_email", claims.Email)
		c.Set("user_id", claims.Subject)

		return next(c)
	}
}
