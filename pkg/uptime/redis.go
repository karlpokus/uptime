package uptime

import (
	"context"
	"time"
	"github.com/go-redis/redis"
)

func redisCall(s *Service) error {
	opts := &redis.Options{
		Addr: s.Url,
		Password: s.Pwd, // empty string is ignored
		DB: 0,
	}
	ctx, _ := context.WithTimeout(context.Background(), 3 * time.Second)
	client := redis.NewClient(opts).WithContext(ctx)
	defer client.Close()
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
