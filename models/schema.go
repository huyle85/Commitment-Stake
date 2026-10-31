package models

import "time"

type User struct {
	WalletAddress string    `gorm:"type:varchar(42);primaryKey"`
	Name          string    `gorm:"type:varchar(100)"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

type Event struct {
	ID            uint      `gorm:"primaryKey"`
	Name          string    `gorm:"type:varchar(255);not null"`
	DepositAmount float64   `gorm:"type:decimal(10,2);not null"`
	StartTime     time.Time `gorm:"not null"`
	EndTime       time.Time `gorm:"not null"`
	Status        string    `gorm:"type:enum('upcoming', 'ongoing', 'ended', 'distributed');default:'upcoming'"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

type Registration struct {
	ID            uint      `gorm:"primaryKey"`
	EventID       uint      `gorm:"not null;uniqueIndex:idx_event_user"`
	WalletAddress string    `gorm:"type:varchar(42);not null;uniqueIndex:idx_event_user"`
	Status        string    `gorm:"type:enum('pending', 'staked', 'checked_in', 'no_show');default:'pending'"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}