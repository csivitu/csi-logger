package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/csivitu/csi-logger/helpers"
	"github.com/csivitu/csi-logger/initializers"
	"github.com/csivitu/csi-logger/models"
	"github.com/redis/go-redis/v9"
)


func GetFromCache(key string, ctx context.Context) ([]models.Log, error) {
	data, err := initializers.RedisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("item not found in cache")
		}
		go helpers.LogServerError("Error Getting from cache", err, "")
		return nil, fmt.Errorf("error getting from cache")
	}
	var logs []models.Log
	err = json.Unmarshal([]byte(data), &logs)
	if err != nil {
		go helpers.LogServerError("Error unmarshalling data from cache", err, "")
		return nil, fmt.Errorf("error unmarshalling data from cache")
	}
	return logs, nil
}

func SetToCache(key string, data []models.Log, ctx context.Context) error {

	jsonData, err := json.Marshal(data)
    if err != nil {
        go helpers.LogServerError("Error marshalling data to JSON", err, "")
        return fmt.Errorf("error marshalling data to JSON")
    }

	if err := initializers.RedisClient.Set(ctx, key, jsonData, initializers.CacheExpirationTime).Err(); err != nil {
		go helpers.LogServerError("Error Setting to cache", err, "")
		return fmt.Errorf("error setting to cache")
	}
	return nil
}

func FlushCache(ctx context.Context) error {
	if err := initializers.RedisClient.FlushAll(ctx).Err(); err != nil {
		go helpers.LogServerError("Error Flushing Cache", err, "")
		return fmt.Errorf("error flushing cache")
	}
	return nil
}