// Code generated by hertz generator.

package publish

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	apiLog "github.com/prometheus/common/log"
	"io"
	"simple-douyin/api/biz/client"
	mw "simple-douyin/api/biz/middleware"
	bizPublish "simple-douyin/api/biz/model/publish"
	kitexPublish "simple-douyin/kitex_gen/publish"
)

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var bizReq bizPublish.PublishActionRequest
	resp := new(bizPublish.PublishActionResponse)
	apiLog.Info("PublishAction begin.")

	// err = c.BindAndValidate(&bizReq)
	//if err != nil {
	//	c.String(consts.StatusBadRequest, err.Error())
	//	return
	//}
	// apiLog.Info("After bind.")

	bizReq.Token = string(c.FormValue("token"))
	apiLog.Info("token: ", bizReq.Token)

	if bizReq.Token == "" {
		apiLog.Info(err)
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "Unauthorized"
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	title := c.FormValue("title")
	if title == nil {
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "Empty video title"
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	bizReq.Title = string(title)
	apiLog.Info("title: ", string(title))

	fileHeader, err := c.FormFile("data")
	if err != nil {
		apiLog.Info(err)
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		apiLog.Info(err)
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	bizReq.Data, err = io.ReadAll(file)
	if err != nil {
		apiLog.Info(err)
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	apiLog.Info("data length: ", len(bizReq.Data))

	// 该接口需要登录态，需要确认具体身份，仅在路由时鉴权即可
	// 通过中间件获取用户id
	apiLog.Info("Getting userId")
	loggedClaims, exist := c.Get("JWT_PAYLOAD")
	if !exist {
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "Unauthorized"
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	userId := int64(loggedClaims.(jwt.MapClaims)[mw.JwtMiddleware.IdentityKey].(float64))
	apiLog.Info("userId: ", userId)

	req := kitexPublish.PublishActionRequest{
		UserId: userId,
		Data:   bizReq.Data,
		Title:  bizReq.Title,
	}

	apiLog.Info("Publish Action.")
	resp, err = client.PublishAction(ctx, &req)
	apiLog.Info("After Publish Action.")

	if err != nil {
		apiLog.Info(err)
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var bizReq bizPublish.PublishListRequest
	resp := new(bizPublish.PublishListResponse)
	err = c.BindAndValidate(&bizReq)
	if err != nil {
		apiLog.Info(err)
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	// 该接口需要登录态，需要确认具体身份，仅在路由时鉴权即可
	req := kitexPublish.PublishListRequest{
		UserId: bizReq.UserID,
	}

	resp, err = client.PublishList(ctx, &req)
	if err != nil {
		apiLog.Info(err)
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}
