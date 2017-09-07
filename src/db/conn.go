package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/golang/glog"
	"github.com/minicool/mms-service/src/conf"
	"strconv"
)

var Conn *gorm.DB

func connect(host string, database string, user string, pass string) (db *gorm.DB, err error) {

		dbConnString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, database)
		db, err = gorm.Open("mysql", dbConnString)
		if err != nil {
			glog.Errorf("failed to connect database %s",dbConnString)
		}else{
			glog.Errorf("connect database %s sucessed",dbConnString)
		}
		Conn = db
	return Conn,err
}

func GetDBConnection() *gorm.DB {
	config,err := conf.LoadConfig()
	if err != nil {
		return nil
	}
	address_port := strconv.Itoa(config.Db_conf.Port)
	address := config.Db_conf.Host + ":" + address_port

	conn, err := connect(address, config.Db_conf.Db, config.Db_conf.User, config.Db_conf.Pass)
	if err != nil {
		glog.Error(err.Error())
	}
	conn.SingularTable(true)
	AutoMigrate()
	return conn
}