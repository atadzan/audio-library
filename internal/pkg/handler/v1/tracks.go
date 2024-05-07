package v1

import (
	"log"

	"github.com/atadzan/audio-library/internal/pkg/models"
	"github.com/atadzan/audio-library/internal/pkg/repository"
	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) uploadTrack(c *fiber.Ctx) error {
	inputFile, err := c.FormFile("file")
	if err != nil || inputFile == nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errInvalidInputParamsMsg))
	}

	file, err := inputFile.Open()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errInvalidInputParamsMsg))
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Println(err)
		}
	}()

	path, err := h.storage.Upload(c.Context(), inputFile.Filename, file)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMsgResponse(errInternalServerMsg))
	}

	return c.Status(fiber.StatusOK).JSON(map[string]string{"message": successMsg, "path": path})
}

func (h *Handler) createTrack(c *fiber.Ctx) error {
	var input models.CreateTrack
	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errInvalidInputParamsMsg))
	}

	if err := input.Validate(); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(err.Error()))
	}

	if err := h.repo.Tracks.Save(c.UserContext(), input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMsgResponse(errInternalServerMsg))
	}

	return c.Status(fiber.StatusOK).JSON(newMsgResponse(successMsg))
}

func (h *Handler) listTracks(c *fiber.Ctx) error {
	var input models.PaginationParams
	if err := c.QueryParser(&input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errInvalidInputParamsMsg))
	}

	input.Validate()

	tracks, err := h.repo.Tracks.List(c.UserContext(), input)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMsgResponse(errInternalServerMsg))
	}

	return c.Status(fiber.StatusOK).JSON(tracks)
}

func (h *Handler) likeTrack(c *fiber.Ctx) error {
	var input models.TrackId
	if err := c.ParamsParser(&input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errInvalidInputParamsMsg))
	}
	if err := h.repo.Tracks.Like(c.UserContext(), input.Id); err != nil {
		log.Println(err)
		switch {
		case errors.Is(err, repository.ErrOperationFailed):
			return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errOperationFailedMsg))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(newMsgResponse(errInternalServerMsg))
		}
	}
	return c.Status(fiber.StatusOK).JSON(successMsg)
}

func (h *Handler) revertLike(c *fiber.Ctx) error {
	var input models.TrackId
	if err := c.ParamsParser(&input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errInvalidInputParamsMsg))
	}
	if err := h.repo.Tracks.RevertLike(c.UserContext(), input.Id); err != nil {
		log.Println(err)
		switch {
		case errors.Is(err, repository.ErrOperationFailed):
			return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errOperationFailedMsg))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(newMsgResponse(errInternalServerMsg))
		}
	}
	return c.Status(fiber.StatusOK).JSON(successMsg)
}

func (h *Handler) listFavouriteTracks(c *fiber.Ctx) error {
	var input models.PaginationParams
	if err := c.QueryParser(&input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(newMsgResponse(errInvalidInputParamsMsg))
	}

	input.Validate()

	tracks, err := h.repo.Tracks.LikedList(c.UserContext(), input)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(newMsgResponse(errInternalServerMsg))
	}

	return c.Status(fiber.StatusOK).JSON(tracks)
}
