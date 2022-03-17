package examples

import (
	"context"
	"time"

	goredis "github.com/go-redis/redis/v8"
	"github.com/goodboy3/referenceMem/basic"
	"github.com/goodboy3/referenceMem/plugin/redisClient"
)

//redis example
func Redis_run() {
	if redisClient.GetInstance() != nil {
		key := redisClient.GetInstance().GenKey("foo")
		redisClient.GetInstance().Set(context.Background(), key, "redis-bar", 100*time.Second)
		str, err := redisClient.GetInstance().Get(context.Background(), "redis-foo").Result()
		if err != nil && err != goredis.Nil {
			basic.Logger.Errorln(err)
		}
		basic.Logger.Debugln(str)
	}
}
