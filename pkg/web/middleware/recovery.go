package middleware

import (
	"app_blade/pkg/logging"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewLoggingRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				if brokenPipe {
					logging.Default().ErrorfContext(ctx.Request.Context(),
						"[URL]: %s [ERROR]: %v [REQUEST]: %s",
						ctx.Request.URL.String(),
						err, string(httpRequest),
					)
					// If the connection is dead, we can't write a status to it.
					ctx.Error(err.(error)) // nolint: errcheck
					ctx.Abort()
					return
				}
				_, file, line, _ := runtime.Caller(3)
				logging.Default().ErrorfContext(ctx.Request.Context(),
					"[Recovery from panic %s] [URL]: %s [ERROR]: %v [REQUEST]: %s",
					fmt.Sprintf("%s:%d", file, line),
					ctx.Request.URL.String(),
					err,
					string(httpRequest),
				)
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}
