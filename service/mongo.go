package service

import (
	"context"
	"time"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func mongoCall(s Service) error {
	ctx, _ := context.WithTimeout(context.Background(), 3 * time.Second)
	client, err := mongo.Connect(ctx, s.Url)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
	return client.Ping(ctx, nil)
}
