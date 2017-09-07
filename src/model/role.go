package model

import (
	"github.com/minicool/mms-service/src/db"
	"github.com/golang/glog"
)

type Role struct {
	ID int 				`gorm:"primary_key;not null;"`
	RoleName string		`gorm:"type:varchar(50);not null;"`
	RoleDesc string		`gorm:"type:varchar(200);not null;"`
	Status int			`gorm:"type:tinyint;"`
}

func init() {
	db.RegisterModel(&Role{})
}

func CreateTable_Role(){
	err := db.Conn.HasTable(&Role{})
	if err == false{
		db.Conn.CreateTable(&Role{})
		glog.Info("create table dept")
	}else{
		glog.Info("dept table is exists")
	}
}

func Add_Role_One(rolename string,roledesc string,status int){
	role := Role{RoleName:roledesc,RoleDesc:roledesc,Status:status}
	if db.Conn.NewRecord(role) == false {
		glog.Error("role primary key is blank")
	}else{
		db.Conn.Create(&role)
	}
}

func Find_Role_One(roleId int) (*Role, error){
	role := &Role{ID: roleId}
	res := db.Conn.First(&role, "id=?", roleId)
	err := res.Error
	if err != nil {
		glog.Errorf("Error finding role with id %s: %v", roleId, res.Error.Error())
	}
	return role, err
}