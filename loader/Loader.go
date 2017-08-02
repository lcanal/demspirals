package loader

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL Driver for GORM
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	cache "github.com/patrickmn/go-cache"
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
	dbDriver := viper.GetString("db.driver")
	if dbDriver == "mysql" {
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

	//Using local sqlite.
	db, err := gorm.Open("sqlite3", "demspirals.db")
	if err != nil {
		log.Fatalf("Error creating demspirals.db: %s\n", err.Error())
	}
	return db

}

//ReadFromCache reads from cachestore using the key
func ReadFromCache(key string) (interface{}, bool) {
	var cachedobjs map[string]cache.Item

	f, err := os.Open("cache/playerstats")
	if err != nil {
		log.Printf("Error opening cache file %s:%v\n", "/tmp/cache", err)
		return nil, false
	}
	gob.Register(cachedobjs)
	dec := gob.NewDecoder(f)
	err = dec.Decode(&cachedobjs)
	if err != nil {
		log.Printf("Error decoding file:%s\n", err.Error())
		return nil, false
	}
	c := cache.NewFrom(5*time.Minute, 10*time.Minute, cachedobjs)

	return c.Get(key)
}

//WriteToCache writes to cachestore (file) based on key.
func WriteToCache(key string, obj interface{}) {
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set(key, obj, cache.DefaultExpiration)

	f, err := os.Create("cache/playerstats")
	if err != nil {
		log.Printf("Error making cache file %s:%s", "/tmp/cache", err.Error())
		return
	}

	//gob.Register(obj)
	enc := gob.NewEncoder(f)
	err = enc.Encode(c.Items())
	if err != nil {
		log.Printf("Error encoding var: %s", err.Error())
		return
	}
	f.Close()
}
