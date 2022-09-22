package repo

import (
	"context"
	"logger-service/logs"
	"logger-service/models"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	LogEntry models.LogEntry
	Repo     LoggerRepoInterface
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
