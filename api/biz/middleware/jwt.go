package middleware

import (
	"context"
	"time"

	client "simple-douyin/api/biz/client"
	"simple-douyin/api/biz/model/common"
	"simple-douyin/api/biz/model/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "simple-douyin-jwt-identity"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "simple-douyin-jwt",
		Key:           []byte("1qaz0plm"),
		Timeout:       10 * time.Minute,
		MaxRefresh:    5 * time.Minute,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		IdentityKey:   IdentityKey,
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req user.UserLoginRequest
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			resp := new(user.UserLoginResponse)
			err = client.UserLogin(ctx, &req, resp)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			if resp.StatusCode != 0 {
				return nil, jwt.ErrFailedAuthentication
			}
			c.Set("temp_use", resp.UserID) // set user id to context(only for temp use)
			return resp.UserID, nil        // return user id a int64
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			userID, exist := c.Get("temp_use")
			c.Set("temp_use", nil) // clear it after use
			resp := new(user.UserLoginResponse)
			if !exist {
				resp.StatusCode = 57001
				resp.StatusMsg = new(string)
				*resp.StatusMsg = "Unauthorized"
				c.JSON(consts.StatusOK, resp)
				return
			}
			resp.StatusCode = 0
			resp.UserID = userID.(int64)
			resp.Token = token
			c.JSON(consts.StatusOK, resp)
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			userid, _ := claims[IdentityKey].(int64)
			return &common.User{
				ID: userid,
			}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(consts.StatusUnauthorized, utils.H{
				"status_code": 57001,
				"status_msg":  "Unauthorized",
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
