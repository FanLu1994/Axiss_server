package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var GlobalDb *gorm.DB

func init() {
	//globalConfig := config.GetGlobalConfig()

	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/Axiss?charset=utf8mb4&parseTime=True&loc=Local",
	//	globalConfig.Mysql.User,
	//	globalConfig.Mysql.Password,
	//	globalConfig.Mysql.Addr)

	//dsn := "myuser:123123123@tcp(110.42.182.92:3307)/Axiss?charset=utf8mb4&parseTime=True&loc=Local"

	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("axiss.db"), &gorm.Config{})
	if err != nil {
		//log.Errorln("无法连接到数据库",dsn)
		log.Errorln("无法连接到数据库:", "sqlite")
		panic("无法连接到数据库")
	}
	err = db.AutoMigrate(&Feed{}, &BenchMark{})
	if err != nil {
		log.Errorln("数据库迁移失败:", "sqlite")
		panic("数据库迁移失败")
	}
	GlobalDb = db
}
