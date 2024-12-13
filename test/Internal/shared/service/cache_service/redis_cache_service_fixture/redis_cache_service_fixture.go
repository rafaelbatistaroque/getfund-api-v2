package redis_cache_service_fixture

import (
	"context"
	"getfund-api-v2/internal/shared/service/cache_service"
	redisconfig "getfund-api-v2/pkg/redis"
	"getfund-api-v2/test/helper/settings_spy"
)

func NewSut() (cache_service.Cache, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	redisClient := redisconfig.New(ctx, &settings_spy.ApplicationSettingsSpy{})
	return cache_service.New(redisClient, ctx), cancel
}
