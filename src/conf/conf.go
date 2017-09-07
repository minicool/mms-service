package conf

import (
	"encoding/json"
	"io/ioutil"
	"github.com/golang/glog"
//	"strconv"
)

type Ser_conf struct{
	Host string
	Port int
}

type Db_conf struct {
	Host string
	Port int
	Db string
	User string
	Pass string
}

type Config struct {
	Ser_conf Ser_conf
	Db_conf Db_conf
}

var Json_config= Config{}

func readFile(filename string) (Config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		glog.Errorf("ReadFile: ", err.Error())
		return Json_config, err
	}

	if err := json.Unmarshal(bytes, &Json_config); err != nil {
		glog.Errorf("Unmarshal: ", err.Error())
		return Json_config, err
	}

	return Json_config, nil
}

//load config
func LoadConfig() (Config,error) {
	//config read file

	Json_config, err := readFile("config/config.json")
	if err != nil {
		glog.Error("config readFile: ", err.Error())
		return Json_config,err
	}else{
		glog.Error(" load config success")
	}
	//address_port := strconv.Itoa(config["client_conf"].Port)
	//service.address = config["client_conf"].Host + ":" + address_port

	return Json_config,nil
}
