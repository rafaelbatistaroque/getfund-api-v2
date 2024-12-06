package rediscacheservicefixture

import (
	"context"
	"getfund-api-v2/internal/shared/service/cacheservice"
	redisconfig "getfund-api-v2/pkg/redis"

	"getfund-api-v2/test/helper/settingsspy"
)

func NewSut() (cacheservice.Cache, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	redisClient := redisconfig.New(ctx, &settingsspy.ApplicationSettingsSpy{})
	return cacheservice.New(redisClient, ctx), cancel
}
