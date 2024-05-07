package v1

import (
	"log"

	"github.com/atadzan/audio-library/internal/pkg/models"
	"github.com/atadzan/audio-library/internal/pkg/repository"
	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) signUp(c *fiber.Ctx) error {
	var input models.SignUp
	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errInvalidInputParamsMsg))
	}
	userId, err := h.repo.Users.Save(c.Context(), input)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMsgResponse(errInternalServerMsg))
	}

	token, err := generateToken(userId)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMsgResponse(errInternalServerMsg))
	}

	return c.Status(fiber.StatusOK).JSON(newSuccessAuthResp(token))
}

func (h *Handler) signIn(c *fiber.Ctx) error {
	var input models.SignIn

	userId, err := h.repo.Users.GetUserId(c.Context(), input)
	if err != nil {
		log.Println(err)
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errUserNotFoundMsg))
		default:
			return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errInternalServerMsg))
		}
	}

	token, err := generateToken(userId)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMsgResponse(errInternalServerMsg))
	}

	return c.Status(fiber.StatusOK).JSON(newSuccessAuthResp(token))
}

func (h *Handler) getProfile(c *fiber.Ctx) error {
	profile, err := h.repo.Users.Get(c.UserContext())
	if err != nil {
		log.Println(err)
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errUserNotFoundMsg))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(newMsgResponse(errInternalServerMsg))
		}
	}

	return c.Status(fiber.StatusOK).JSON(profile)
}
