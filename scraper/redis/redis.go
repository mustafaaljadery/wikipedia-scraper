package redis

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func Client()(*redis.Client){
	godotenv.Load()
	redisUrl := os.Getenv("REDIS_URL")

	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}
	
	client := redis.NewClient(opt)
	return client
}