package repository

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/atadzan/audio-library/internal/pkg/models"
)

const (
	passwordHashSalt = "helloWorld"
)

func getUserId(ctx context.Context) int {
	rawValue := ctx.Value(models.UserIdCtxKey)
	if rawValue != nil {
		userId, err := strconv.Atoi(rawValue.(string))
		if err != nil {
			return 0
		}
		return userId
	}
	return 0
}

func generatePasswordHash(password *string) error {

	// convert password string to slice of bytes
	passwordBytes := []byte(*password)

	// creating sha-512 header
	sha512Header := sha512.New()

	// convert passwordSalt string to slice of bytes
	saltBytes := []byte(passwordHashSalt)

	// Append passwordSalt to password
	passwordBytes = append(passwordBytes, saltBytes...)

	// write password to sha-512 header
	if _, err := sha512Header.Write(passwordBytes); err != nil {
		return err
	}

	// get sha-512 hashed password
	hashedPassword := sha512Header.Sum(nil)

	// convert hashed password to HEX string
	*password = hex.EncodeToString(hashedPassword)

	return nil

}

func getOffset(page, limit int) int {
	return (page - 1) * limit
}

func addStorageDomain(path *string) {
	if path == nil || len(*path) == 0 {
		*path = ""
	}
	*path = fmt.Sprintf("http://localhost:9001/tracks/%s", *path) // localhost:9001 API of minIO storage and `tracks` is a bucket
}
