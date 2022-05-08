package mysql

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

type MysqlDB struct {
	*xorm.Engine
}

var test *MysqlDB

func InitMySQL() {
	test = initDB(getArgs("test"))
	//test.AutoMigrate(&model.User{})
}

func initDB(args string) *MysqlDB {
	db, err := xorm.NewEngine("mysql", args)
	if err != nil {
		panic("Fail to open database.")
	}
	err = db.DB().Ping()
	if err != nil {
		panic("Fail to connect database.")
	}
	fmt.Println("Successfully connected!")
	return &MysqlDB{
		Engine: db,
	}
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
