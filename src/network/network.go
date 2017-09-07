package network

import (
	"sync"
	pb "github.com/minicool/mms-service/proto"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	//"github.com/minicool/mms-service/src/model"
)

type rfidDeviceServer struct{
	employers []*pb.DeviceRegeditRequest
	m      sync.RWMutex
	//emplyer data
	address string
}

// HelloService Hello服务
var RFIDDeviceServer = rfidDeviceServer{}

func NewServer() *rfidDeviceServer {
	s := new(rfidDeviceServer)
	//	s.loadFeatures(*jsonDBFile)
	//	s.employers = make(map[uint64][]*pb.Employer)
	//	s.employers = make([]*pb.Employer,4)
	return s
}

func (service *rfidDeviceServer) DeviceRegedit(ctx context.Context, in *pb.DeviceRegeditRequest) (*pb.DeviceRegeditResponse, error) {
	glog.Errorf("DeviceRegedit %s",in)
	//model.AddDeviceRegedit(in.DeviceName,int32(in.DeviceType),in.IpAdd,in.IpPort,in.MacAdd)
	var resp = &pb.DeviceRegeditResponse{}
	resp.Success = true
	resp.DeviceRegedid = 1000
	resp.Errormsg = &pb.ErrorMessage{pb.ERRORCODE_ERRORCODE_NONE,"执行成功"}
	return resp, nil
}

func (service *rfidDeviceServer)RFIDprintWriteDataStream(in *pb.RFIDprintWriteDataStream_Request,stream pb.RFIDDeviceServer_RFIDprintWriteDataStreamServer) error{
	glog.Errorf("RFIDprint_writeData_stream %s",in)
	var resp = &pb.RFIDprintWriteDataStream_Response{}
	resp.Data = &pb.RFIDprint_Data{AssetsId:"1234",CompanyName:"测试公司",DevpementNmae:"开发公司",
		AssetsName:"资产名称",AssetsType:"电脑",AssetsModel:"测试型号",Url:"www.baidu.com",Date:"2017-04-13"}
	resp.Errormsg = &pb.ErrorMessage{pb.ERRORCODE_ERRORCODE_NONE,"执行成功"}
	resp.Success = true
	for true{
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}
