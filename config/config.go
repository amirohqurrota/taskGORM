package config

import (
	"GORM-2/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	//declare struct config &variable connectionString
	var err error
	dsn := "root:22juli1998@tcp(0.0.0.0:3306)/alterralearn?parseTime=True"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&models.User{})
}

// func migrate {

// 	GormDB.AutoMigrate(&article{})
// 	GormDB.AutoMigrate(&trial{})
// 	// initeMigrate()
// }
