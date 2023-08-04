namespace go favorite

include "./common.thrift"

// 点赞操作
struct FavoriteActionRequest {
    1: required string token (api.query = "token"),          // 用户鉴权token
    2: required i64 video_id (api.query = "video_id"),       // 视频id
    3: required i32 action_type (api.query = "action_type"), // 1-点赞，2-取消点赞
}

struct FavoriteActionResponse {
    1: required i32 status_code,   // 0-成功，1-失败
    2: optional string status_msg, // 错误信息
}

// 点赞列表
struct FavoriteListRequest {
    1: required i64 user_id (api.query = "user_id"), // 用户id
    2: required string token (api.query = "token"),  // 用户鉴权token
}

struct FavoriteListResponse {
    1: required i32 status_code,               // 0-成功，1-失败
    2: optional string status_msg,             // 错误信息
    3: required list<common.Video> video_list, // 视频列表
}

service FavoriteService {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req) (api.post = "/douyin/favorite/action"),
    FavoriteListResponse FavoriteList(1: FavoriteListRequest req) (api.post = "/douyin/favorite/list"),
}
