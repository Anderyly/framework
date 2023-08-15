package dal

type Model struct {
	Id        int64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt int64 `gorm:"not null" json:"created_at"`
	UpdatedAt int64 `gorm:"not null" json:"updated_at"`
	DeletedAt int64 `gorm:"not null" json:"deleted_at"`
}
