package services

import (
	"fmt"
	"log"
	"logger-service/config"
	"logger-service/repo"
	"net"
	"net/rpc"
)

type RPCserviceInt interface {
	Listen(config.Config)
}

type rPCServer struct {
	r repo.RPCRepoInterfcate
}

func NewRpcServer(r repo.RPCRepoInterfcate) RPCserviceInt {
	return &rPCServer{r}
}

func (s *rPCServer) Listen(conf config.Config) {
	log.Println("Starting RPC server on ", conf.RpcPort)
	listen, err := net.Listen("tcp", "logger-service:5001")
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}
	defer listen.Close()

	for {
		rpcCon, err := listen.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("rpcCon: %v\n", rpcCon)
		go rpc.ServeConn(rpcCon)
	}
}
