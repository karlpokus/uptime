package service

import (
	"context"
	"time"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func mongoCall(s Service) error {
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	mongoConnString := s.Url
	client, err := mongo.Connect(ctx, mongoConnString)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}
	return nil
}
