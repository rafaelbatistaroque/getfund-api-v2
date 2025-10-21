package custom_cache_fixture

import (
	"context"
	config_redis_cache "getfund-api-v2/internal/config/redis_cache"
	"getfund-api-v2/internal/shared/cache"
	"getfund-api-v2/test/helper/settings_spy"
)

func NewSut() (cache.Service, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	return config_redis_cache.New(ctx, &settings_spy.ApplicationSettingsSpy{}), cancel
}
