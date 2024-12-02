package rediscachefixture

import (
	"context"
	"getfund-api-v2/internal/shared/service/cacheservice"
	"getfund-api-v2/test/spyshared/settingsspy"
)

func NewSut() cacheservice.Cache {
	return cacheservice.New(context.Background(), &settingsspy.ApplicationSettingsSpy{})
}
