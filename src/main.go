package main

import (
	"github.com/golang/glog"
	"net"
	"os"
	"flag"
	"github.com/minicool/mms-service/src/conf"
	"strconv"
	"google.golang.org/grpc"
	pb "github.com/minicool/mms-service/proto"
	"github.com/minicool/mms-service/src/network"
	"github.com/minicool/mms-service/src/contronl"
	"github.com/minicool/mms-service/src/db"
)

func main() {
	//Init the command-line flags.
	flag.Parse()
	// user temp logs
	defer glog.Infof("Temp folder for log files: %s\n", os.TempDir())
	// Flushes all pending log I/O.
	defer glog.Flush()
	//
	conn := db.GetDBConnection()
	defer conn.Close()
	contronl.InitDb()
	contronl.Test()

	//
	config,err := conf.LoadConfig()
	if err != nil {
		return
	}
	address_port := strconv.Itoa(config.Ser_conf.Port)
	address := config.Ser_conf.Host + ":" + address_port
	glog.Infof("address %s %d",config.Ser_conf.Host,config.Ser_conf.Port)

	listen, err := net.Listen("tcp", address)
	if err != nil {
		glog.Infof("Failed to listen: %v", err)
	}

	// 实例化grpc Server, 并开启TLS认证
	s := grpc.NewServer(/*grpc.Creds(creds)*/)
	// 注册HelloService
	pb.RegisterRFIDDeviceServerServer(s, network.NewServer())


	glog.Errorf("Listen on " + address + " with TLS")

	//	grpclog.Infoln("Listen on " + address + " with TLS")

	s.Serve(listen)
}
