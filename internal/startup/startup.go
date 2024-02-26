package startup

import (
	"app_blade/pkg/app"
	"app_blade/pkg/web"

	"github.com/google/wire"
)

// StartupConfig 初始化服务
func StartupConfig(webServer *web.WebServer) app.InitializeRunner {
	return func() []app.AppRunner {
		return []app.AppRunner{
			webServer,
		}
	}
}

var ProviderSet = wire.NewSet(StartupConfig)
