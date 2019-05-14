package controller

import (
	"context"
	"errors"
	pb "rpcServer/pb"
	"rpcServer/service"
	"rpcServer/utils"
)

// Register:user register
func (u *User) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	var (
		resp = &pb.RegisterReply{}
		err  error
	)
	// param check
	if utils.IsStringEmpty(in.Username) || utils.IsStringEmpty(in.Password) {
		logger.Info("param is empty")
		return resp, errors.New("param is empty")
	}

	if resp, err = service.UserRegister(in); err != nil {
		logger.Error(err)
		return resp, err
	}
	return resp, nil
}

func (u *User) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}

func (u *User) UserByUID(ctx context.Context, in *pb.UserByUIDRequest) (*pb.UserByUIDReply, error) {
	return &pb.UserByUIDReply{}, nil
}
