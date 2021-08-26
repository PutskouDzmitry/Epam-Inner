package data

import (
	"fmt"
	"gorm.io/gorm"
)

type EntityData struct {
	db *gorm.DB
}

func NewEntityData(db *gorm.DB) *EntityData {
	return &EntityData{db: db}
}

func (e EntityData) GetEntityDb(to string, from string, space string) int {
	var a int = -1
	query := fmt.Sprintf("timeUTC > ? AND ? > timeUTC AND spaceId = ?")
	e.db.Table("visits_scud").Select("id").Where(query, to, from, space).First(&a)
	return a
}
