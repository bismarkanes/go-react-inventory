package customtools

import (
	"fmt"
	model "itemservice/model"

	"gorm.io/gorm"
)

func initialSeed(db *gorm.DB) {
	// 1. check if item 1 is existed
	items := []model.Item{}
	if result := db.Where("id = ?", 1).Find(&items); result.Error != nil {
		fmt.Println("Db error : ", result.Error.Error())
		return
	}

	if len(items) == 0 {
		item := model.Item{
			ID:    1,
			Stock: 2500,
		}

		// 2. create the item 1
		if result := db.Create(&item); result.Error != nil {
			fmt.Println("Db error : ", result.Error.Error())
			return
		}
	}
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&model.Item{}, &model.ReservationItem{})
	initialSeed(db)
}
