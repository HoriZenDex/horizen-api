package models

// Like model.
type Like struct {
	ID     uint `json:"id" gorm:"primaryKey;autoIncrement"`
	PostID uint `json:"post_id"`
	UserID uint `json:"user_id"`
}
