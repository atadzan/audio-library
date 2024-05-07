package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/atadzan/audio-library/internal/pkg/models"
	"github.com/go-errors/errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TracksRepo struct {
	db *pgxpool.Pool
}

func NewTracksRepo(pool *pgxpool.Pool) *TracksRepo {
	return &TracksRepo{
		db: pool,
	}
}

func (t *TracksRepo) Save(ctx context.Context, params models.CreateTrack) error {
	query := fmt.Sprintf(`INSERT INTO %s(title, artist, genre, path, uploader_id, created_at) 
											VALUES($1, $2, $3, $4, $5, $6)`, tracksTable)
	rows, err := t.db.Exec(ctx, query, params.Title, params.Artist, params.Genre, params.Path, getUserId(ctx), time.Now())
	if err != nil {
		return err
	}
	if rows.RowsAffected() == 0 {
		return errors.New(ErrOperationFailed)
	}

	return nil
}

func (t *TracksRepo) List(ctx context.Context, params models.PaginationParams) (tracks []models.Track, err error) {
	query := fmt.Sprintf(`SELECT t.id, t.title, t.artist, t.genre, t.path, u.fullname, t.created_at FROM %s t
                                                              JOIN %s u ON u.id = t.uploader_id
                                                                        LIMIT $1 OFFSET $2`, tracksTable, usersTable)
	rows, err := t.db.Query(ctx, query, params.Limit, getOffset(params.Page, params.Limit))
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var track models.Track
		if err = rows.Scan(&track.Id, &track.Title, &track.Artist, &track.Genre, &track.Path, &track.Uploader, &track.CreatedAt); err != nil {
			return
		}
		addStorageDomain(&track.Path)
		tracks = append(tracks, track)
	}

	if len(tracks) == 0 {
		return []models.Track{}, nil
	}

	return
}

func (t *TracksRepo) Like(ctx context.Context, trackId int) error {
	query := fmt.Sprintf(`INSERT INTO %s(track_id, user_id, created_at) VALUES($1, $2, $3) ON CONFLICT(track_id, user_id) DO NOTHING`, likedTracksTable)
	row, err := t.db.Exec(ctx, query, trackId, getUserId(ctx), time.Now())
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return errors.New(ErrOperationFailed)
	}
	return nil
}

func (t *TracksRepo) RevertLike(ctx context.Context, trackId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE track_id=$1 AND user_id=$2`, likedTracksTable)
	row, err := t.db.Exec(ctx, query, trackId, getUserId(ctx))
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return errors.New(ErrOperationFailed)
	}
	return nil
}

func (t *TracksRepo) LikedList(ctx context.Context, params models.PaginationParams) (tracks []models.Track, err error) {
	query := fmt.Sprintf(`SELECT t.id, t.title, t.artist, t.genre, t.path, u.fullname, t.created_at FROM %s lt
                                                              JOIN %s t ON lt.track_id = t.id 
															  JOIN %s u ON u.id = t.uploader_id
                                                              WHERE lt.user_id=$1
                                                              LIMIT $2 OFFSET $3`, likedTracksTable, tracksTable, usersTable)
	rows, err := t.db.Query(ctx, query, getUserId(ctx), params.Limit, getOffset(params.Page, params.Limit))
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var track models.Track
		if err = rows.Scan(&track.Id, &track.Title, &track.Artist, &track.Genre, &track.Path, &track.Uploader, &track.CreatedAt); err != nil {
			return
		}
		addStorageDomain(&track.Path)
		tracks = append(tracks, track)
	}

	if len(tracks) == 0 {
		return []models.Track{}, nil
	}

	return
}
