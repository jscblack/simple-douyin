package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Default .
func Default(ctx context.Context, c *app.RequestContext) {
	// <h1> Hello, Hertz! </h1>
	// <h1> Hello, Simple-Douyin! </h1>
	c.String(consts.StatusOK, "<!DOCTYPE html> <html> <head> <meta charset=\"utf-8\"> <title>Simple-Douyin</title> </head> <body> <h1>Hello, Hertz!</h1> <h1>Hello, Simple-Douyin!</h1> </body> </html>")
	c.Header("Content-Type", "text/html; charset=utf-8")
	// c.JSON(consts.StatusOK, utils.H{
	// 	"message": "Hello, Hertz!",
	// })
}
