package uptime

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func mongoCall(s *Service) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.Url))
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
	return client.Ping(ctx, nil)
}
