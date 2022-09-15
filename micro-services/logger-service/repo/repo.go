package repo

import (
	"context"
	"fmt"
	"log"
	"logger-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoggerRepoInterface interface {
	Insert(context.Context, models.LogEntry) error
	All(context.Context) ([]*models.LogEntry, error)
	GetById(context.Context, string) (*models.LogEntry, error)
	DropCollection(context.Context) error
	UpdateCollection(context.Context, *models.LogEntry) (*mongo.UpdateResult, error)
}

type logger struct {
	db *mongo.Client
}

func NewLoggerRepo(db *mongo.Client) LoggerRepoInterface {
	return &logger{db}
}

func (r *logger) Insert(ctx context.Context, entry models.LogEntry) error {
	collection := r.db.Database("logs").Collection("logs")

	_, err := collection.InsertOne(context.TODO(), models.LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *logger) All(ctx context.Context) ([]*models.LogEntry, error) {
	var logs []*models.LogEntry

	collection := r.db.Database("logs").Collection("logs")

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item models.LogEntry

		err := cursor.Decode(&item)
		if err != nil {
			log.Print(err.Error())
			return nil, err
		} else {
			logs = append(logs, &item)
		}
	}
	return logs, nil
}

func (r *logger) GetById(ctx context.Context, id string) (*models.LogEntry, error) {
	var res models.LogEntry
	collection := r.db.Database("logs").Collection("logs")

	fmt.Printf("id: %v\n", id)
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": docId}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *logger) DropCollection(ctx context.Context) error {
	collection := r.db.Database("logs").Collection("logs")

	if err := collection.Drop(ctx); err != nil {
		return err
	}
	return nil
}

func (r *logger) UpdateCollection(ctx context.Context, ref *models.LogEntry) (*mongo.UpdateResult, error) {
	collection := r.db.Database("logs").Collection("logs")
	docId, err := primitive.ObjectIDFromHex(ref.ID)
	if err != nil {
		return nil, err
	}
	res, err := collection.UpdateOne(ctx,
		bson.M{"_id": docId},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "name", Value: ref.Name},
				{Key: "data", Value: ref.Data},
				{Key: "updated_at", Value: time.Now()},
			}},
		},
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
