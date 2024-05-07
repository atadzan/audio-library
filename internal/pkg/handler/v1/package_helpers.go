package v1

import (
	"time"

	"github.com/atadzan/audio-library/internal/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func generateToken(userId int) (string, error) {
	claims := &models.TokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resultToken, err := token.SignedString([]byte(models.JWTSigningKey))
	if err != nil {
		return "", err
	}
	return resultToken, nil
}

func getUserIdFromCtx(ctx *fiber.Ctx) (userId int) {
	rawValue := ctx.UserContext().Value(models.UserIdCtxKey)
	if rawValue != nil {
		userId, ok := rawValue.(int)
		if !ok {
			userId = 0
		}
		return userId
	}
	return userId
}
