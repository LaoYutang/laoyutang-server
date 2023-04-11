package utils

import (
	"context"
	"time"

	"github.com/laoyutang/laoyutang-server/modules/db"
	"github.com/redis/go-redis/v9"
)

func UseRedis[T any](key string, getData func() (T, error), expires time.Duration) (res T, errOut error) {
	ctx := context.Background()
	redisStr := ""

	redisStr, errOut = db.Redis.Get(ctx, key).Result()
	if errOut == redis.Nil {
		res, errOut = getData()
		if errOut != nil {
			return
		}
		db.Redis.Set(ctx, key, ToJson(res), expires)
	} else if errOut != nil {
		return
	} else {
		errOut = ParseJson(redisStr, &res)
		if errOut != nil {
			return
		}
	}

	return
}
