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

//内部rpc
// 获取点赞数
struct FavorCountRequest {
    1: required i64 user_id,
}

struct FavorCountResponse {
    1: required i32 status_code,   // 0:成功 other:失败
    2: optional string status_msg,
    3: optional i64 favor_count,  //点赞数
}

// 获取被赞数
struct FavoredCountRequest {
    1: required i64 user_id,
}

struct FavoredCountResponse {
    1: required i32 status_code,    // 0:成功 other:失败
    2: optional string status_msg,
    3: optional i64 favored_count, //被赞数
}

// 获取点赞关系(user_id是否点赞vedio_id)
struct IsFavorRequest {
    1: required i64 user_id,
    2: required i64 vedio_id,
}

struct IsFavorResponse {
    1: required i32 status_code,   // 0:成功 other:失败
    2: optional string status_msg,
    3: required bool is_favor,     //是否点赞
}

service FavoriteService {
    AddFavoriteActionResponse AddFavoriteAction(1: AddFavoriteActionRequest req),
    DelFavoriteActionResponse DelFavoriteAction(1: DelFavoriteActionRequest req),
    FavoriteListResponse FavoriteList(1: FavoriteListRequest req),
    FavorCountResponse FavorCount(1: FavorCountRequest req),
    FavoredCountResponse FavoredCount(1: FavoredCountRequest req),
    IsFavorResponse IsFavor(1: IsFavorRequest req), 
}