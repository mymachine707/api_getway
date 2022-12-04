package models

import "time"

// Content ...
type Content struct {
	Title string `json:"title" binding:"required" `
	Body  string `json:"body" binding:"required"`
}

// Article ...
type Article struct {
	ID        string     `json:"id"`
	Content              // Promoted fields
	AuthorID  string     `json:"author_id" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// CreateArticleModul ...
type CreateArticleModul struct {
	Content         // Promoted fields
	AuthorID string `json:"author_id" binding:"required"`
}

// PackedArticleModel ...
type PackedArticleModel struct {
	ID        string     `json:"id" binding:"required"`
	Content              // Promoted fields
	Author    Author     `json:"author_id" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// UpdateArticleModul ...
type UpdateArticleModul struct {
	ID      string `json:"id" binding:"required"`
	Content        // Promoted fields
}
