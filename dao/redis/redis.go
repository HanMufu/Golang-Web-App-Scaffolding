package redis

import (
	"context"
	"fmt"
	"go-web-app/settings"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

var (
	rdb *redis.Client
)

// 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	//rdb = redis.NewClient(&redis.Options{
	//	Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
	//	Password: viper.GetString("redis.password"), // no password set
	//	DB:       viper.GetInt("redis.db"),          // use default DB
	//	PoolSize: viper.GetInt("redis.pool_size"),   // 连接池大小
	//})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return
}

func Close() {
	_ = rdb.Close()
}

//func V8Example() {
//	ctx := context.Background()
//	if err := initClient(); err != nil {
//		return
//	}
//
//	err := rdb.Set(ctx, "key", "value", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	val, err := rdb.Get(ctx, "key").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("key", val)
//
//	val2, err := rdb.Get(ctx, "key2").Result()
//	if err == redis.Nil {
//		fmt.Println("key2 does not exist")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("key2", val2)
//	}
//	// Output: key value
//	// key2 does not exist
//}
