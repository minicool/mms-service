package model

import (
	"github.com/minicool/mms-service/src/db"
	"github.com/golang/glog"
)

type User struct {
//	gorm.Model
	ID int 				`gorm:"primary_key;not null;"`
	LoginName string 	`gorm:"type:varchar(50);not null;"`
	Psw	string			`gorm:"type:varchar(100);not null;"`
	RealName string		`gorm:"type:varchar(50);"`
	Gender int			`gorm:"type:tinyint;"`
	Mobile string		`gorm:"type:varchar(20);not null;"`
	Email string		`gorm:"type:varchar(100);"`

	Dept Dept 			`gorm:"ForeignKey:DeptId;"`
	DeptId int			`gorm:"type:int;"`

	Role Role			`gorm:"ForeignKey:RoleId;"`
	RoleId int			`gorm:"type:int;not null;"`
	Status int			`gorm:"type:tinyint;"`
	Skin string			`gorm:"type:varchar(45);"`
	DepMan int 			`gorm:"type:int;"`
	ZcTypeMan int 		`gorm:"type:int;"`
	RecordMan int		`gorm:"type:int;"`

}

func init() {
	//	db.RegisterMigration(`DROP TABLE IF EXISTS "users"`)
	db.RegisterModel(&User{})
}

func CreateTable_User(){
	err := db.Conn.HasTable(&User{})
	if err == false{
		db.Conn.CreateTable(&User{})
		glog.Info("create table user")
	}else{
		glog.Info("user table is exists")
	}
}

func GetAllUser() ([]User, error) {
	var user []User
	res := db.Conn.Find(&user)
	return user,res.Error
}

func Add_User(loginName string,Psw	string,RealName string,Gender int,Mobile string,
	Email string,DeptId int,RoleId int,Status int,Skin string,DepMan int,ZcTypeMan int,
	RecordMan int){

	dept,err := FindDeptOne(DeptId)
	if err != nil {
		return
	}

	role,err1 := Find_Role_One(RoleId)
	if err1 != nil {
		return
	}

	user := User{LoginName:loginName,Psw:Psw,RealName:RealName,Gender:Gender,Mobile:Mobile,Email:Email,
	Dept:*dept,Role:*role,Status:Status,Skin:Skin,DepMan:DepMan,ZcTypeMan:ZcTypeMan,RecordMan:RecordMan}

	if db.Conn.NewRecord(user) == false {
		glog.Error("DeviceRegedit primary key is blank")
	}else{
		db.Conn.Create(&user)
	}
}