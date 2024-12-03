package rediscacheservicefixture

import (
	"context"
	"getfund-api-v2/internal/shared/service/cacheservice"
	"getfund-api-v2/test/spyshared/settingsspy"
)

func NewSut() (cacheservice.Cache, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	return cacheservice.New(ctx, &settingsspy.ApplicationSettingsSpy{}), cancel
}
