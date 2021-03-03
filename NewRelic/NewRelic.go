package NewRelic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"os"
)

type NewRelicMiddleware struct {
	License, Env string
	TracerEnable bool
}

func (r *NewRelicMiddleware) Middleware() gin.HandlerFunc {
	// Relic Middleware
	AppName := fmt.Sprintf(`Koala Listing Api Auth %s`, r.Env)
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(AppName),
		newrelic.ConfigLicense(r.License),
		newrelic.ConfigDistributedTracerEnabled(r.TracerEnable),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if err != nil {
		panic(err)
	}
	// End Relic Middleware

	// returning Gin Middleware with
	return nrgin.Middleware(app)
}
