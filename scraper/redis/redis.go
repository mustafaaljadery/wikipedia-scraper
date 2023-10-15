package redis

import (
	"github.com/redis/go-redis/v9"
)

func Client()(*redis.Client){
	opt, err := redis.ParseURL("mask")
	if err != nil {
		panic(err)
	}
	
	client := redis.NewClient(opt)
	return client
}