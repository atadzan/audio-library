package v1

import (
	"context"
	"strconv"
	"strings"

	"github.com/atadzan/audio-library/internal/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenRequiredParts = 2
)

func (h *Handler) parseToken(accessToken string) (int, error) {
	var claims models.TokenClaims
	_, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.JWTSigningKey), nil
	})

	if err != nil {
		return 0, err
	}

	return claims.UserId, nil
}

func (h *Handler) middlewareUserIdentify(c *fiber.Ctx) error {
	accessToken := c.Get("Authorization")
	if len(accessToken) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(newMsgResponse("empty bearer token"))
	}

	headerParts := strings.Split(accessToken, " ")
	if len(headerParts) != tokenRequiredParts {
		return c.Status(fiber.StatusUnauthorized).JSON(newMsgResponse("invalid token"))
	}
	userId, err := h.parseToken(headerParts[1])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse("invalid token"))
	}

	if userId == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(newMsgResponse("unauthorized"))
	}

	c.SetUserContext(context.WithValue(c.Context(), models.UserIdCtxKey, strconv.Itoa(userId)))

	return c.Next()
}
