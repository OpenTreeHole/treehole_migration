package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"treehole_migration/config"
)

var DB *gorm.DB

var gormConfig = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	},
}

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.Config.DbUrl), gormConfig)
	if err != nil {
		panic(err)
	}

	//DB = DB.Debug()

	err = DB.AutoMigrate(
		&AnonynameMapping{},
		&FloorLike{},
		&FloorHistory{},
	)
	if err != nil {
		panic(err)
	}
}
