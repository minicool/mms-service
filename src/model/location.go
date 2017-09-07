package model

import (
	"github.com/minicool/mms-service/src/db"
	"github.com/golang/glog"
)

type Location struct {
	ID int			`gorm:"primary_key;not null;"`
	DeptId int		`gorm:"type:int;not null;"`
	LocSN string	`gorm:"column:LocSN;type:varchar(30);not null;"`
	LocName string	`gorm:"column:LocName;type:varchar(50);not null;"`
	Tip string		`gorm:"type:varchar(100);"`
	Stauts int		`gorm:"type:tinyint;"`

}

func init() {
	db.RegisterModel(&Dept{})
}

func CreateTable_Location(){
	err := db.Conn.HasTable(&Location{})
	if err == false{
		db.Conn.CreateTable(&Location{})
		glog.Info("create table location")
	}else{
		glog.Info("location table is exists")
	}
}

func Find_Loaction_one(locationName string) (*Location,error){
	loc := &Location{LocName: locationName}
	res := db.Conn.First(&loc, "locName=?", locationName)
	err := res.Error
	if err != nil {
		glog.Errorf("Error finding dept with locName %s: %v", locationName, res.Error.Error())
	}
	return loc, err
}

