package service

import (
	pb "rpcServer/pb"
	"rpcServer/utils"
)

func UserRegister(body *pb.RegisterRequest) (resp *pb.RegisterReply, err error) {
	resp = &pb.RegisterReply{}
	logger.Infof("username=%v,password=%v,phoneNum=%v,country=%v",
		body.Username, body.Password, body.PhoneNum, body.Country)
	resp.Uid = utils.EncodeMD5(body.Password)
	return
}
