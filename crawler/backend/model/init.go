package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/lexkong/log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/olivere/elastic.v5"
)

var DB *Database
var Client *elastic.Client

type Database struct {
	Self *gorm.DB
}


func (db *Database) Init() {
	var err error
	DB = &Database{
		Self: GetSelfDB(),
	}
	Client ,err = elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
}

func (db *Database) Close() {
	DB.Self.Close()
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// 建立数据库连接
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	// 用于设置最大打开的连接数，默认为0不限制，设置最大的连接数可以避免并发太高
	// 导致连接mysql出现 too many connections 的错误
	//db.DB().SetMaxOpenConns(20000)
	// 用于设置闲置的连接数，设置闲置连接数则当开启的一个连接使用完成后可以放在连接池里等待下次使用
	db.DB().SetMaxIdleConns(0)
}

// 使用了cli
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}






