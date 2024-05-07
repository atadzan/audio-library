package repository

import (
	"context"
	"log"

	"fmt"
	"time"

	"github.com/atadzan/audio-library/internal/pkg/models"
	"github.com/go-errors/errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UsersRepo struct {
	db *pgxpool.Pool
}

func NewUsersRepo(pool *pgxpool.Pool) *UsersRepo {
	return &UsersRepo{
		db: pool,
	}
}

func (u *UsersRepo) Save(ctx context.Context, params models.SignUp) (userId int, err error) {
	if err = generatePasswordHash(&params.Password); err != nil {
		return
	}
	query := fmt.Sprintf(`INSERT INTO %s(fullname, login, password_hash, created_at) VALUES($1, $2, $3, $4) RETURNING id`, usersTable)
	if err = u.db.QueryRow(ctx, query, params.Fullname, params.Login, params.Password, time.Now()).Scan(&userId); err != nil {
		return
	}

	return

}
func (u *UsersRepo) Get(ctx context.Context) (profile models.UserProfile, err error) {
	query := fmt.Sprintf(`SELECT fullname, login, created_at FROM %s WHERE id=$1`, usersTable)
	if err = u.db.QueryRow(ctx, query, getUserId(ctx)).Scan(&profile.Fullname, &profile.Login, &profile.RegisteredAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return profile, errors.New(ErrNotFound)
		}
		return
	}
	return
}

func (u *UsersRepo) GetUserId(ctx context.Context, params models.SignIn) (userId int, err error) {
	if err = generatePasswordHash(&params.Password); err != nil {
		return 0, err
	}
	query := fmt.Sprintf(`SELECT id FROM %s WHERE login=$1 AND password_hash=$2`, usersTable)
	if err = u.db.QueryRow(ctx, query, params.Login, params.Password).Scan(&userId); err != nil {
		log.Println(err)
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, errors.New(ErrNotFound)
		}
		return
	}
	return
}
