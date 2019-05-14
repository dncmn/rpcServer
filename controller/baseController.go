package controller

import "rpcServer/utils/logging"

type (
	User          struct{}
	Server        struct{}
	StreamService struct{}
)

var (
	logger = logging.GetLogger()
)
