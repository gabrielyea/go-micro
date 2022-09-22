package services

import (
	"context"
	"fmt"
	"log"
	"logger-service/config"
	"logger-service/logs"
	"logger-service/models"
	"logger-service/repo"
	"net"

	"google.golang.org/grpc"
)

type GrpcServerInt interface {
	Listen(config.Config)
	WriteLog(context.Context, *logs.LogRequest) (*logs.LogResponse, error)
}

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Repo repo.LoggerRepoInterface
}

func NewGrpcServer(r repo.LoggerRepoInterface) GrpcServerInt {
	return &LogServer{
		Repo: r,
	}
}

func (l *LogServer) Listen(conf config.Config) {
	fmt.Printf("conf.GrpcPort: %v\n", conf.GrpcPort)
	lis, err := net.Listen("tcp", conf.GrpcPort)
	if err != nil {
		log.Fatalf("Failed to listen for: %v", err)
		return
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, l)
	log.Printf("gRpc server started on port %s", conf.GrpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to listen")
	}

}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	logEntry := models.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Repo.Insert(ctx, logEntry)
	if err != nil {
		res := logs.LogResponse{Result: "Failed"}
		return &res, err
	}

	res := &logs.LogResponse{Result: "Logged"}
	return res, nil
}
