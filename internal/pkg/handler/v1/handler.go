package v1

import (
	"github.com/atadzan/audio-library/internal/pkg/repository"
	"github.com/atadzan/audio-library/internal/pkg/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Handler struct {
	repo    *repository.Repository
	storage storage.Storage
}

func NewHandler(repo *repository.Repository, storage storage.Storage) *Handler {
	return &Handler{
		repo:    repo,
		storage: storage,
	}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	app.Use(
		logger.New(), // for logging requests
	)
	// group endpoints with versioning
	base := app.Group("/v1")
	{
		auth := base.Group("/auth")
		{
			auth.Post("/signUp", h.signUp)
			auth.Post("/signIn", h.signIn)
		}

		authorizedRoutes := base.Group("", h.middlewareUserIdentify)
		{
			authorizedRoutes.Get("/profile", h.getProfile)

			tracks := authorizedRoutes.Group("/tracks")
			{
				tracks.Post("/upload", h.uploadTrack)
				tracks.Post("", h.createTrack)
				tracks.Get("", h.listTracks)
				tracks.Post("/:id/like", h.likeTrack)
				tracks.Post("/:id/revertLike", h.revertLike)
				tracks.Get("/favourite", h.listFavouriteTracks)
			}
		}
	}
}
