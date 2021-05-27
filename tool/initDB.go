package tool

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"writer/model"
)

var (
	OrmDB *gorm.DB
)

func InitDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/writehelp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	_ = db.AutoMigrate(&model.Author{}, &model.Book{}, &model.Chapter{},
		&model.ChapterSetting{}, &model.Rolesetting{}, &model.Inspiration{}, &model.PlotSetting{})

	OrmDB = db

}

func GetOrmDB() *gorm.DB {
	return OrmDB
}
