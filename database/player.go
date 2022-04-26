package database

import "time"

type Player struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Password    string
	Score       int `gorm:"default:0"`
	DisplayName string
}

func (p Player) PrettyName() string {
	if p.DisplayName != "" {
		return p.DisplayName
	}

	return p.Name
}

type PlayerVersion struct {
	ID         uint `gorm:"primaryKey"`
	Player     Player
	PlayerID   uint
	Version    int
	Entrypoint string
	IsLatest   bool
	CreatedAt  time.Time
}
