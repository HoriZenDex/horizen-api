package models

// Post model with minimal fields.
type Post struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Description string `json:"description"`
	Content     string `json:"content"` // Image URL
	Tag         string `json:"tag"`
	OwnerID     uint   `json:"owner_id"`
}
