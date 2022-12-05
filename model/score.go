package model

import (
	"os"

	"gorm.io/gorm"
)

type Score struct {
	GameTitle string `json:"game_title" gorm:"column:game_title;not null"`
	GameType  string `json:"game_type" gorm:"column:game_type;not null"`
	UserID    string `json:"user_id" gorm:"column:user_id;not null"`
	UserName  string `json:"user_name" gorm:"column:user_name;not null"`
	UserScore int64  `json:"user_score" gorm:"column:user_score;default:0"`
}

type GormScore struct {
	gorm.Model
	Score
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Board to `test_leader_board`
func (GormScore) TableName() string {
	name := os.Getenv("DB_TABLE_NAME")
	if name == "" {
		name = "test_table"
	}
	return name
}
