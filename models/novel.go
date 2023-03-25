package models

type Novel struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Synopsis  string `json:"synopsis"`
	Cover     string `json:"cover"`
	Thumbnail string `json:"thumbnail"`
	WordCount int    `json:"wordCount"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
	DeletedAt int    `json:"deletedAt"`
}
