package database

import "time"

type Player struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Password string
	Score    int
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
