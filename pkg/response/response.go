package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (r *Response) WithCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) WithMessage(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) WithData(data any) *Response {
	r.Data = data
	return r
}

func (r *Response) JSON(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, r)
}

func Result(code int, message string) *Response {
	return &Response{Code: code, Message: message}
}

func OK() *Response {
	return Result(http.StatusOK, "OK")
}

func BadRequest() *Response {
	return Result(http.StatusBadRequest, "Bad Request")
}

func Unauthorized() *Response {
	return Result(http.StatusUnauthorized, "Unauthorized")
}

func Forbidden() *Response {
	return Result(http.StatusForbidden, "Forbidden")
}

func NotFound() *Response {
	return Result(http.StatusNotFound, "Not Found")
}

func InternalServerError() *Response {
	return Result(http.StatusInternalServerError, "Internal Server Error")
}

func MethodNotAllowed() *Response {
	return Result(http.StatusMethodNotAllowed, "Method Not Allowed")
}

func ServiceUnavailable() *Response {
	return Result(http.StatusServiceUnavailable, "Service Unavailable")
}

func ServiceError() *Response {
	return Result(10001, "Service Error")
}
