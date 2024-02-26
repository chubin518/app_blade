package middleware

import (
	"app_blade/pkg/logging"
	"bytes"
	"io"
	"net/http"
	"path/filepath"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

type ExcludedPatterns []*regexp.Regexp

func NewLoggingRequest(patterns ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Next()
			return
		}
		start := time.Now()
		path := ctx.Request.URL.Path
		if pathMatch(path, patterns...) {
			ctx.Next()
			return
		}
		reqBuffer := make([]byte, 0)
		if ctx.Request.Body != http.NoBody {
			body, err := ctx.GetRawData()
			if err != nil {
				reqBuffer = []byte(err.Error())
			} else {
				reqBuffer = body
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		}

		logging.Default().InfofContext(ctx.Request.Context(),
			"Request Starting [%s %s] Request Body [%s]",
			ctx.Request.Method,
			ctx.Request.URL.String(),
			string(reqBuffer),
		)

		respBuffer := new(bytes.Buffer)
		ctx.Writer = &responseWriter{Writer: io.MultiWriter(ctx.Writer, respBuffer), ResponseWriter: ctx.Writer}

		// Process request
		ctx.Next()

		cost := time.Since(start)
		if len(ctx.Errors) > 0 {
			logging.Default().ErrorfContext(ctx.Request.Context(),
				"Error [%v] Request Finished [%s %s] Status [%d] Response Body [%s] Time [%d] ms",
				ctx.Errors.ByType(gin.ErrorTypePrivate).String(),
				ctx.Request.Method,
				ctx.Request.URL.String(),
				ctx.Writer.Status(),
				respBuffer.String(),
				cost.Milliseconds(),
			)
		} else {
			logging.Default().InfofContext(ctx.Request.Context(),
				"Request Finished [%s %s] Status [%d] Response Body [%s] Time [%d] ms",
				ctx.Request.Method,
				ctx.Request.URL.String(),
				ctx.Writer.Status(),
				respBuffer.String(),
				cost.Milliseconds(),
			)
		}
	}
}

type responseWriter struct {
	io.Writer
	gin.ResponseWriter
}

func (w responseWriter) Write(p []byte) (int, error) {
	return w.Writer.Write(p)
}

func pathMatch(path string, patterns ...string) bool {
	if len(path) == 0 || len(patterns) == 0 {
		return false
	}
	for _, pattern := range patterns {
		if ok, _ := filepath.Match(pattern, pattern); ok {
			return true
		}
	}
	return false
}
