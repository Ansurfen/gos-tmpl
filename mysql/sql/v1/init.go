package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type MysqlDB struct {
	*sql.DB
}

var test *MysqlDB

func InitMySQL() {
	test.DB = initDB(getArgs("test"))
}

func initDB(args string) *sql.DB {
	db, err := sql.Open("mysql", args)
	if err != nil {
		fmt.Println("Fail to open database.")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Fail to connect database.")
	}
	fmt.Println("Successfully connected!")
	return db
}

func getArgs(database string) string {
	dir, _ := os.Getwd()
	conf := viper.New()
	conf.SetConfigName("config")
	conf.SetConfigType("yml")
	conf.AddConfigPath(dir + "/db/mysql")
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	host := conf.GetString(database + ".host")
	port := conf.GetString(database + ".port")
	username := conf.GetString(database + ".username")
	password := conf.GetString(database + ".password")
	charset := conf.GetString(database + ".charset")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
}
