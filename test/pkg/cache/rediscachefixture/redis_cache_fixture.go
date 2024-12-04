package rediscacheservicefixture

import (
	"context"
	"getfund-api-v2/internal/pkg/cache"
	"getfund-api-v2/test/spyshared/settingsspy"
)

func NewSut() (cache.Cache, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	return cache.New(ctx, &settingsspy.ApplicationSettingsSpy{}), cancel
}
