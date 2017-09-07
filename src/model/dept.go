package model

import (
	"github.com/minicool/mms-service/src/db"
	"github.com/golang/glog"
)

type Dept struct {
	ID int 				`gorm:"primary_key;not null;"`
	Pid int				`gorm:"type:int;not null;"`
	DeptName string		`gorm:"type:varchar(100);not null;"`
	DeptLeader string	`gorm:"type:varchar(50);"`
	Phone string		`gorm:"type:varchar(30);"`
	SortNum int			`gorm:"type:smallint;"`
	Tip string			`gorm:"type:varchar(300)"`
	Status int			`gorm:"type:tinyint;"`
	Location []Location `gorm:"ForeignKey:DeptId"`
}

func init() {
	//	db.RegisterMigration(`DROP TABLE IF EXISTS "users"`)
	db.RegisterModel(&Dept{})
}

func (Dept) TableName() string {
	return "companydept"
}

func CreateTable_Dept(){
	err := db.Conn.HasTable(&Dept{})
	if err == false{
		db.Conn.CreateTable(&Dept{})
		glog.Info("create table dept")
	}else{
		glog.Info("dept table is exists")
	}
}

func GetAllDept() ([]Dept, error) {
	var depts []Dept
	res := db.Conn.Find(&depts)
	return depts,res.Error
}

func FindDeptOne(id int) (*Dept, error) {
	dept := &Dept{ID: id}
	res := db.Conn.First(&dept, "id=?", id)
	err := res.Error
	if err != nil {
		glog.Errorf("Error finding dept with id %d: %v", id, res.Error.Error())
	}
	return dept, err
}

func Add_Dept_One(pId int,deptName string,deptLeader string,Phone string,sortNum int,
	tip string,status int){
	dept := Dept{Pid:pId,DeptName:deptName,DeptLeader:deptLeader,Phone:Phone,SortNum:sortNum,
	Tip:tip,Status:status}
	if db.Conn.NewRecord(dept) == false {
		glog.Error("dept primary key is blank")
	}else{
		db.Conn.Create(&dept)
	}
}

func Add_DeptLoaction(id int,location []*Location){
	dept,err := FindDeptOne(id)
	if err != nil {
		return
	}
	for _,value := range location{
		dept.Location = append(dept.Location, *value)
	}
	db.Conn.Update(&dept)
}

