package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/csivitu/csi-logger/cache"
	"github.com/csivitu/csi-logger/models"
	"github.com/csivitu/csi-logger/schemas"
)

func FindCache(resourceID string, urlParams schemas.LogFetchSchema, ctx context.Context) ([]models.Log, error) {

	cacheKey := fmt.Sprintf("%s-%d-%d", resourceID , urlParams.Limit, urlParams.Page)

	cachedResult, err := cache.GetFromCache(cacheKey, ctx)
	if err == nil {
		return cachedResult, nil
	}
	return nil, errors.New("cache miss")
}