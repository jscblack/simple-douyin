package handler_test

import (
	"testing"

	"simple-douyin/api/biz/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestDefault(t *testing.T) {
	h := server.Default()
	h.GET("/", handler.Default)
	w := ut.PerformRequest(h.Engine, "GET", "/", nil, ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "<!DOCTYPE html> <html> <head> <meta charset=\"utf-8\"> <title>Simple-Douyin</title> </head> <body> <h1>Hello, Hertz!</h1> <h1>Hello, Simple-Douyin!</h1> </body> </html>", string(resp.Body()))
	assert.DeepEqual(t, "text/html; charset=utf-8", string(resp.Header.Get("Content-Type")))
}

func TestDefault_UnknownRoute(t *testing.T) {
	h := server.Default()
	h.GET("/index", handler.Default)
	w := ut.PerformRequest(h.Engine, "GET", "/index", nil, ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "<!DOCTYPE html> <html> <head> <meta charset=\"utf-8\"> <title>Simple-Douyin</title> </head> <body> <h1>Hello, Hertz!</h1> <h1>Hello, Simple-Douyin!</h1> </body> </html>", string(resp.Body()))
	assert.DeepEqual(t, "text/html; charset=utf-8", string(resp.Header.Get("Content-Type")))
}

func TestDefault_HealthCheck(t *testing.T) {
	h := server.Default()
	h.GET("/health", handler.Default)
	w := ut.PerformRequest(h.Engine, "GET", "/health", nil, ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "<!DOCTYPE html> <html> <head> <meta charset=\"utf-8\"> <title>Simple-Douyin</title> </head> <body> <h1>Hello, Hertz!</h1> <h1>Hello, Simple-Douyin!</h1> </body> </html>", string(resp.Body()))
	assert.DeepEqual(t, "text/html; charset=utf-8", string(resp.Header.Get("Content-Type")))
}
