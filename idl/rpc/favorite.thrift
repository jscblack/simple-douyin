namespace go favorite

include "./common.thrift"

// 点赞操作
struct AddFavoriteActionRequest {
    1: required i64 user_id,  // 用户id
    2: required i64 video_id, // 视频id
}

struct AddFavoriteActionResponse {
    1: required i32 status_code,   // 0-成功，1-失败
    2: optional string status_msg, // 错误信息
}

// 销赞操作
struct DelFavoriteActionRequest {
    1: required i64 user_id,  // 用户id
    2: required i64 video_id, // 视频id
}

struct DelFavoriteActionResponse {
    1: required i32 status_code,   // 0-成功，1-失败
    2: optional string status_msg, // 错误信息
}

// 点赞列表
struct FavoriteListRequest {
    1: required i64 user_id, // 用户id
}

struct FavoriteListResponse {
    1: required i32 status_code,               // 0-成功，1-失败
    2: optional string status_msg,             // 错误信息
    3: required list<common.Video> video_list, // 视频列表
}

service FavoriteService {
    AddFavoriteActionResponse AddFavoriteAction(1: AddFavoriteActionRequest req),
    DelFavoriteActionResponse DelFavoriteAction(1: DelFavoriteActionRequest req),
    FavoriteListResponse FavoriteList(1: FavoriteListRequest req),
}