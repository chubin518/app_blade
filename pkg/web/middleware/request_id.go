package middleware

import (
	"app_blade/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const REQUEST_ID = "X-Request-Id"

func NewRequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Next()
			return
		}
		requestId := ctx.GetHeader(REQUEST_ID)
		if requestId == "" {
			id, err := uuid.NewV7()
			if err != nil {
				id = uuid.New()
			}
			requestId = id.String()
			ctx.Request.Header.Set(REQUEST_ID, requestId)
		}
		ctx.Header(REQUEST_ID, requestId)
		traceCtx := logging.Default().WithContext(ctx.Request.Context(), map[string]any{"traceId": requestId})
		ctx.Request = ctx.Request.WithContext(traceCtx)
		ctx.Next()
	}
}
