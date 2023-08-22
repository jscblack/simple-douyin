namespace go user

include "./common.thrift"

// 用户注册
struct UserRegisterRequest {
    1: required string username, // 注册用户名，最长32个字符
    2: required string password, // 密码，最长32个字符
}

struct UserRegisterResponse {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: required i64 user_id,       // 用户id
}

// 用户登录
struct UserLoginRequest {
    1: required string username, // 登录用户名
    2: required string password, // 登录密码
}

struct UserLoginResponse {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: required i64 user_id,       // 用户id
}

// 用户信息
struct UserInfoRequest {
    1: optional i64 user_id,    // 发送请求的用户id
    2: required i64 to_user_id, // 要查询的用户id
}

struct UserInfoResponse {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: required common.User user,  // 用户信息
}

# // 更新用户统计信息
# // 在redis中更新对应统计量
# struct UpdateUserCounterRequest {
#     1: required i64 user_id,    // 用户id
#     2: required string counter, // 计数器
#     3: required bool increment, // 是否增加
# }

# struct UpdateUserCounterResponse {
#     1: required i32 status_code,   // 状态码，0-成功，其他值-失败
#     2: optional string status_msg, // 返回状态描述
# }

service UserService {
    UserRegisterResponse UserRegister(1: UserRegisterRequest req),
    UserLoginResponse UserLogin(1: UserLoginRequest req),
    UserInfoResponse UserInfo(1: UserInfoRequest req),
    # UpdateUserCounterResponse UpdateUserCounter(1: UpdateUserCounterRequest req),
}
