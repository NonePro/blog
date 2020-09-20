package models

import (
	"fmt"
	"github.com/NonePro/blog/pkg/setting"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	section, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "fail to get section 'database': %v ", err)
	}

	dbType = section.Key("TYPE").String()
	dbName = section.Key("NAME").String()
	user = section.Key("USER").String()
	password = section.Key("PASSWORD").String()
	host = section.Key("HOST").String()
	tablePrefix = section.Key("TABLE_PREFIX").String()

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbName))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
