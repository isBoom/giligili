package task

import "singo/cache"

func RestarDailyRank() error {
	return cache.RedisClient.Del(cache.DailyRankKey).Err()
}
