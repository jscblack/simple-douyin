namespace go relation

include "./common.thrift"

//引入User

// 关注操作
struct RelationActionRequest {
    1: required string token (api.query = "token"),          //鉴权token
    2: required i64 to_user_id (api.query = " to_user_id"),  // 对方用户id
    3: required i32 action_type (api.query = "action_type"), // 1:关注 2:取消关注
}

struct RelationActionResponse {
    1: required i32 status_code,   // 0:成功 other:失败
    2: optional string status_msg,
}

//关注列表
struct RelationFollowListRequest {
    1: required i64 user_id (api.query = "user_id"),
    2: required string token (api.query = "token"),

}

struct RelationFollowListResponse {
    1: required i32 status_code,             // 0:成功 other:失败
    2: optional string status_msg,
    3: optional list<common.User> user_list,
}

//粉丝列表
struct RelationFollowerListRequest {
    1: required i64 user_id (api.query = "user_id"),
    2: required string token (api.query = "token"),
}

struct RelationFollowerListResponse {
    1: required i32 status_code,             // 0:成功 other:失败
    2: optional string status_msg,
    3: optional list<common.User> user_list,
}

//好友列表
struct FriendUser {
    //extend User
    1: required i64 id,                  // 用户id
    2: required string name,             // 用户名称
    3: optional i64 follow_count,        // 关注总数
    4: optional i64 follower_count,      // 粉丝总数
    5: required bool is_follow,          // true-已关注，false-未关注
    6: optional string avatar,           // 用户头像
    7: optional string background_image, // 用户个人页顶部大图
    8: optional string signature,        // 个人简介
    9: optional i64 total_favorited,     // 获赞数量
    10: optional i64 work_count,         // 作品数量
    11: optional i64 favorite_count,     // 点赞数量
    //
    12: optional string message, // 最近一条消息
    13: required i64 msgType,    // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

struct RelationFriendListRequest {
    1: required i64 user_id (api.query = "user_id"),
    2: required string token (api.query = "token"),
}

struct RelationFriendListResponse {
    1: required i32 status_code,              // 0:成功 other:失败
    2: optional string status_msg,
    3: optional list<FriendUser> user_list,
}

service RelationService {
    RelationActionResponse RelationAction(1: RelationActionRequest request) (api.post = "/douyin/relation/action/"),
    RelationFollowListResponse RelationFollowList(1: RelationFollowListRequest request) (api.get = "/douyin/relation/follow/list/"),
    RelationFollowerListResponse RelationFollowerList(1: RelationFollowerListRequest request) (api.get = "/douyin/relation/follower/list/"),
    RelationFriendListResponse RelationFriendList(1: RelationFriendListRequest request) (api.get = "/douyin/relation/friend/list/"),
}
