package model

type Model struct {
	Id        int   `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt int64 `gorm:"not null"`
	UpdatedAt int64 `gorm:"not null"`
	DeletedAt int64 `gorm:"not null"`
}
