package loader

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

import _ "github.com/go-sql-driver/mysql" //Import driver

//ConnectDB initializes a db connection based on settings file.
func ConnectDB() *sql.DB {
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	pass := viper.GetString("db.pass")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, user)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to db: %s", err.Error())
	}
	return db
}
