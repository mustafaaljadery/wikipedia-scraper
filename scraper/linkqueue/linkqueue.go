package linkqueue

import (
	"context"
	r "scraper/redis"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

type Queue [] interface {}

func (q *Queue) LoadQueue(){
	client = r.Client()
	ctx := context.Background()

	iter := client.Scan(ctx, 0, "", 0).Iterator()

	for iter.Next(ctx) {
		key := iter.Val()
		value, _ := client.Get(ctx, key).Result()

		if (value == "not scraped"){
			*q = append(*q, key)
		}
	}
}

func (q *Queue) Enqueue(elem interface{}) {
	*q = append(*q, elem)
	ctx := context.Background()
	str, _ := elem.(string)
	client.Set(ctx, str, "not scraped", 0).Err()
}

func (q *Queue) Dequeue() interface{} {
	ctx := context.Background()
    if len(*q) == 0 {
        return nil
    }
    elem := (*q)[0]
	str, _ := elem.(string)
	client.Set(ctx, str, "scraped", 0).Err()
    *q = (*q)[1:]
    return elem
}

func (q *Queue) GetNext() interface {}{
	elem := (*q)[0]
	return elem
}