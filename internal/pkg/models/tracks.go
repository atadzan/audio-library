package models

import (
	"fmt"
	"time"
)

type CreateTrack struct {
	Title  string `json:"title" `
	Artist string `json:"artist"`
	Genre  string `json:"genre" `
	Path   string `json:"path" `
}

type Track struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Genre     string    `json:"genre" `
	Path      string    `json:"path" `
	Uploader  string    `json:"uploader"`
	CreatedAt time.Time `json:"releaseDate"`
}

type PaginationParams struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

type TrackId struct {
	Id int `params:"id"`
}

func (t *CreateTrack) Validate() error {
	if len(t.Title) == 0 {
		return fmt.Errorf("required field title is empty")
	}
	if len(t.Path) == 0 {
		return fmt.Errorf("required field path is empty")
	}
	if len(t.Genre) == 0 {
		return fmt.Errorf("required field genre is empty")
	}
	if len(t.Artist) == 0 {
		return fmt.Errorf("required field genre is empty")
	}
	return nil
}

func (p *PaginationParams) Validate() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10 // default limit is 10
	}
}
