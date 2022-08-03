package mssql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var DB *gorm.DB

//connect database
func ConnectDB(configPath string) *gorm.DB {
	var err error
	viper.AddConfigPath(configPath)
	if err = viper.ReadInConfig(); err != nil {
		log.Errorln("Fatal Error Config File : ", err)
	}
	conString := fmt.Sprintf("server=%s ; user=%s ; password=%s ; port=%s ; database=%s",
		viper.GetString("mssql.server"),
		viper.GetString("mssql.user"),
		viper.GetString("mssql.password"),
		viper.GetString("mssql.port"),
		viper.GetString("mssql.database"))
	log.Infoln(conString)
	database, err := gorm.Open(viper.GetString("mssql.databaseType"), conString)
	if err != nil {
		log.Errorln("Failed to connect database : ", conString)
		log.Errorln("Error : ", err)
	} else {
		log.Infoln("Database Connected!")
	}
	DB = database
	return DB
}
