package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (s *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 0, nil
}

func (S *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

func (s *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}