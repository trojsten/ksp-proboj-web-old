package database

type Map struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Args string
}
