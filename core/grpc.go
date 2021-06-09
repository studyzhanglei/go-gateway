package core

import (
	"fmt"
	"go-gateway/global"
	"google.golang.org/grpc"
	grpcpool "go-gateway/utils/grpc"
)


func GetGrpcClient2(port string) (conn *grpc.ClientConn ) {
	var err error
	conn, err = grpc.Dial(":" + port, grpc.WithInsecure())
	if err != nil {
		global.LOG.Error(fmt.Sprintf("grpc.Dial err: %s", err.Error()))
	}

	return conn
}

func GetGrpcClient(port string) (p *grpcpool.Pool ) {

	p, err := grpcpool.New(func() (*grpc.ClientConn, error) {
		return grpc.Dial(":" + port, grpc.WithInsecure())
	}, 20, 30, 0)


	if err != nil {
		panic(fmt.Errorf("grpc init failed : %s \n", err))
	}

	fmt.Println(p.Available(), p.Capacity())

	return p
}
