namespace go favorite

include "./common.thrift"

// 点赞操作
struct FavoriteAddActionRequest {
    1: required i64 user_id,  // 用户id
    2: required i64 video_id, // 视频id
}

struct FavoriteAddActionResponse {
    1: required i32 status_code,   // 0-成功，1-失败
    2: optional string status_msg, // 错误信息
}

// 销赞操作
struct FavoriteDelActionRequest {
    1: required i64 user_id,  // 用户id
    2: required i64 video_id, // 视频id
}

struct FavoriteDelActionResponse {
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
    3: optional list<common.Video> video_list, // 视频列表
}

//内部rpc
// 获取点赞数
struct UserFavorCountRequest {
    1: required i64 user_id,
}

struct UserFavorCountResponse {
    1: required i32 status_code,   // 0:成功 other:失败
    2: optional string status_msg,
    3: optional i64 favor_count,  //点赞数
}

// 获取视频被赞数
struct VideoFavoredCountRequest {
    1: required i64 video_id,
}

struct VideoFavoredCountResponse {
    1: required i32 status_code,    // 0:成功 other:失败
    2: optional string status_msg,
    3: optional i64 favored_count, //被赞数
}

// 获取用户被赞数
struct UserFavoredCountRequest {
    1: required i64 user_id,
}

struct UserFavoredCountResponse {
    1: required i32 status_code,    // 0:成功 other:失败
    2: optional string status_msg,
    3: optional i64 favored_count, //被赞数
}

// 获取点赞关系(user_id是否点赞video_id)
struct IsFavorRequest {
    1: required i64 user_id,
    2: required i64 video_id,
}

struct IsFavorResponse {
    1: required i32 status_code,   // 0:成功 other:失败
    2: optional string status_msg,
    3: required bool is_favorite,     //是否点赞
}

service FavoriteService {
    FavoriteAddActionResponse FavoriteAddAction(1: FavoriteAddActionRequest req),
    FavoriteDelActionResponse FavoriteDelAction(1: FavoriteDelActionRequest req),
    FavoriteListResponse FavoriteList(1: FavoriteListRequest req),
    UserFavorCountResponse UserFavorCount(1: UserFavorCountRequest req),
    VideoFavoredCountResponse VideoFavoredCount(1: VideoFavoredCountRequest req),
    UserFavoredCountResponse UserFavoredCount(1: UserFavoredCountRequest req),
    IsFavorResponse IsFavor(1: IsFavorRequest req), 
}