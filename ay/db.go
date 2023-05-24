package ay

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func GetDB() (err error) {
	InitializeRedis()
	var option gorm.Dialector

	switch Yaml.GetString("sql.type") {
	case "mysql":
		log.Println(123)
		dsn := Yaml.GetString("sql.user") + ":" + Yaml.GetString("sql.password") + "@tcp(" + Yaml.GetString("sql.localhost") + ":" + Yaml.GetString("sql.port") + ")/" + Yaml.GetString("sql.database") + "?charset=utf8mb4&parseTime=true&loc=Local"
		option = mysql.New(mysql.Config{DSN: dsn})
	case "sqlite":
		//option = sqlite.Open(Yaml.GetString("sql.localhost"))
	default:

	}

	if Db, err = gorm.Open(option, &gorm.Config{}); err != nil {
		return
	}

	return
}
