package models

type Novel struct {
	ID        uint
	Name      string
	Author    string
	Synopsis  string
	Cover     string
	Thumbnail string
	WordCount string
	CreatedAt int
	UpdatedAt int
	DeletedAt int
}
