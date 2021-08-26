package data

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type EntityData struct {
	db *gorm.DB
}

func NewEntityData(db *gorm.DB) *EntityData{
	return &EntityData{db: db}
}

func (e EntityData) GetEntityDb(to string, from string, space string) int {
	var a int = -1
	quere := fmt.Sprintf("%v >")
	var timeUtc time.Time
	e.db.Table("visits_scud").Select("")
	e.db.Table("visits_scud").Select("*").Where("").First(&str)
	fmt.Println(str)
	return a
}