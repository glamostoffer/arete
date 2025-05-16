package cache

import "github.com/redis/go-redis/v9"

type cache struct {
	cl *redis.Client
}

func New(cl *redis.Client) *cache {
	return &cache{
		cl: cl,
	}
}
