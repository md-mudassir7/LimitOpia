package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type RedisLimiter struct {
	client   *redis.Client
	key      string
	limit    int
	interval time.Duration
}

func NewRedisLimiter(client *redis.Client, key string, limit int, interval time.Duration) *RedisLimiter {
	return &RedisLimiter{
		client:   client,
		key:      key,
		limit:    limit,
		interval: interval,
	}
}

func (rl *RedisLimiter) Allow(ctx context.Context) bool {
	pipe := rl.client.TxPipeline()
	pipe.Incr(ctx, rl.key)
	pipe.Expire(ctx, rl.key, rl.interval)
	cmds, _ := pipe.Exec(ctx)

	if count := cmds[0].(*redis.IntCmd).Val(); count > int64(rl.limit) {
		return false
	}
	return true
}
