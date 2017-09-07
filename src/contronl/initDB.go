package contronl

import (
	"github.com/minicool/mms-service/src/model"
	//"github.com/minicool/mms-service/src/db"
)

func InitDb() {
	model.CreateTable_Dept()
	model.CreateTable_Role()
	model.CreateTable_Dept()
	model.CreateTable_Location()
}

func Test(){
	model.Add_Role_One("测试人员","描述",0)
	model.Add_User("wangzhenhua","123456","振华",
		1,"13000000","123@245.com",3,1,3,"123",
		2,3,3)

	location1,err :=  model.Find_Loaction_one(" '小会议室'")
	if err != nil {
		return
	}
	location2,err :=  model.Find_Loaction_one(" '大会议室'")
	if err != nil {
		return
	}
	model.Add_DeptLoaction(1,[]*model.Location{location1,location2})

	model.Add_DeptLoaction(2,[]*model.Location{location1})
}
