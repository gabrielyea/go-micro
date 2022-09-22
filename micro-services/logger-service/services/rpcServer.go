package services

import (
	"context"
	"fmt"
	"log"
	"logger-service/config"
	"logger-service/models"
	"logger-service/repo"
	"net"
	"net/rpc"
)

type RPCserviceInt interface {
	Listen(config.Config)
	LogInfo(models.RPCPayload, *string) error
}

type RPCServer struct {
	r repo.LoggerRepoInterface
}

func NewRpcServer(r repo.LoggerRepoInterface) RPCserviceInt {
	return &RPCServer{r}
}

func (s *RPCServer) Listen(conf config.Config) {
	log.Println("Starting RPC server on ", conf.RpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("logger-service%s", conf.RpcPort))
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

func (s *RPCServer) LogInfo(payload models.RPCPayload, res *string) error {
	ctx := context.TODO()
	var log models.LogEntry
	log.Name = payload.Name
	log.Data = payload.Data
	err := s.r.Insert(ctx, log)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return err
	}
	*res = "Processed payload by RPC " + payload.Name
	return nil
}
