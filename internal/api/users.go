package api

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"go-landing-page/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userRepository repositories.UserRepository
}

func NewUserHandler(userRepository repositories.UserRepository) UserHandler {
	return UserHandler{userRepository: userRepository}
}

type HashPasswordReq struct {
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (h *UserHandler) HashPassword(c echo.Context) error {
	body := new(HashPasswordReq)
	if err := c.Bind(body); err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to hash password")
	}

	return c.JSON(http.StatusOK, map[string]string{"password_hash": string(passwordHash)})
}

func (h *UserHandler) Login(c echo.Context) error {
	body := new(LoginReq)
	if err := c.Bind(body); err != nil {
		return err
	}

	user, err := h.userRepository.GetByEmail(body.Email)
	if err != nil {
		log.Println("[Login] Erro ao buscar usuario por email:", body.Email, err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Email ou senha incorretos")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))
	if err != nil {
		log.Println("[Login] Erro ao comparar hash de senha:", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Email ou senha incorretos")
	}

	token, err := h.GenerateToken(user.ID, user.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "falha ao gerar token")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *UserHandler) GenerateToken(userId int64, email string) (string, error) {
	claims := CustomClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.FormatInt(userId, 10),
			Audience:  jwt.ClaimStrings{"go-landing-page-api"},
			Issuer:    "go-landing-page",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
