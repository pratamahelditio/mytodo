package domain

type Todo struct {
	ID    int64  `gorm:"primary_key;auto_increment;not_null"`
	Title string `gorm:"type:varchar(255)"`
}
