package repository

import (
	"context"

	"github.com/atadzan/audio-library/internal/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Users interface {
	Save(ctx context.Context, params models.SignUp) (int, error)
	Get(ctx context.Context) (models.UserProfile, error)
	GetUserId(ctx context.Context, params models.SignIn) (userId int, err error)
}

type Tracks interface {
	Save(ctx context.Context, params models.CreateTrack) error
	List(cxt context.Context, params models.PaginationParams) ([]models.Track, error)
	Like(ctx context.Context, trackId int) error
	RevertLike(ctx context.Context, trackId int) error
	LikedList(ctx context.Context, params models.PaginationParams) ([]models.Track, error)
}

type Repository struct {
	Users  *UsersRepo
	Tracks *TracksRepo
}

func New(poolConn *pgxpool.Pool) *Repository {
	return &Repository{
		Users:  NewUsersRepo(poolConn),
		Tracks: NewTracksRepo(poolConn),
	}
}
