/*
IDL 注解说明
https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/usage/annotation/
*/
namespace go user

include "./common.thrift"

// 用户注册
struct UserRegisterRequest {
    1: required string username (api.query = "username"), // 注册用户名，最长32个字符
    2: required string password (api.query = "password"), // 密码，最长32个字符
}

struct UserRegisterResponse {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: required i64 user_id,       // 用户id
    4: required string token,      // 用户鉴权token
}

// 用户登录
struct UserLoginRequest {
    1: required string username (api.query = "username"), // 登录用户名
    2: required string password (api.query = "password"), // 登录密码
}

struct UserLoginResponse {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: required i64 user_id,       // 用户id
    4: required string token,      // 用户鉴权token
}

// 用户信息
struct UserInfoRequest {
    1: required i64 user_id (api.query = "user_id"), // 用户id
    2: required string token (api.query = "token"),  // 用户鉴权token
}

struct UserInfoResponse {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: required common.User user,  // 用户信息
}

service UserService {
    UserRegisterResponse UserRegister(1: UserRegisterRequest req) (api.post = "/douyin/user/register/"),
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post = "/douyin/user/login/"),
    UserInfoResponse UserInfo(1: UserInfoRequest req) (api.get = "/douyin/user/"),
}
