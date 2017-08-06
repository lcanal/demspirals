package loader

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //To Support MySQL Driver
	"github.com/spf13/viper"
)

//ConnectDB initializes a db connection based on settings file.
func ConnectDB() *sql.DB {
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	pass := viper.GetString("db.pass")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, user)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to db: %s\n", err.Error())
	}
	return db
}

//GormConnectDB initializes a GORM DB ORM connection
func GormConnectDB() *gorm.DB {

	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	pass := viper.GetString("db.pass")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", user, pass, host, port, user)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting gorm to db: %s\n", err.Error())
	}
	return db

}
