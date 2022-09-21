package repo

import (
	"context"
	"log"
	"logger-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type RPCRepoInterfcate interface {
	LogInfo(models.RPCPayload, *string) error
}

type RPCrepo struct {
	db *mongo.Client
}

func NewRPCRepo(db *mongo.Client) RPCRepoInterfcate {
	return &RPCrepo{db}
}

func (t *RPCrepo) Test(arg1 models.RPCPayload, arg2 *string) error {
	log.Println("It worked!!", arg1)
	return nil
}

func (r *RPCrepo) LogInfo(payload models.RPCPayload, res *string) error {
	collection := r.db.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), models.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	})
	if err != nil {
		return err
	}

	*res = "Processed payload by RPC " + payload.Name
	return nil
}
