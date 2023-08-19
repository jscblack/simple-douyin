namespace go relation

include "./common.thrift"

//引入User

// 关注操作
struct RelationAddRequest {
    1: required i64 user_id,    //己方用户id
    2: required i64 to_user_id, // 对方用户id
//调用双方的UpdateUserCounter
}

struct RelationAddResponse {
    1: required i32 status_code,   // 0:成功 other:失败
    2: optional string status_msg,
//调用双方的UpdateUserCounter
}

//取关操作
struct RelationRemoveRequest {
    1: required i64 user_id,    //己方用户id
    2: required i64 to_user_id, // 对方用户id
}

struct RelationRemoveResponse {
    1: required i32 status_code,   // 0:成功 other:失败
    2: optional string status_msg,
}

//get关注列表
struct RelationFollowListRequest {
    1: required i64 user_id,
    2: required i64 from_user_id, 
}

struct RelationFollowListResponse {
    1: required i32 status_code,               // 0:成功 other:失败
    2: optional string status_msg,
    3: optional list<common.User> follow_list, //关注列表   调用UserInfo
}

//get粉丝列表
struct RelationFollowerListRequest {
    1: required i64 user_id,
    2: required i64 from_user_id, 
}

struct RelationFollowerListResponse {
    1: required i32 status_code,                 // 0:成功 other:失败
    2: optional string status_msg,
    3: optional list<common.User> follower_list, //粉丝列表
}

struct FriendUser {
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
    12: optional string message,         // 最近一条消息
    13: required i64 msg_type,           // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

//get好友列表
struct RelationFriendListRequest {
    1: required i64 user_id,
    2: required i64 from_user_id,
}

struct RelationFriendListResponse {
    1: required i32 status_code,              // 0:成功 other:失败
    2: optional string status_msg,
    3: optional list<FriendUser> friend_list, //好友列表
}

//内部rpc，针对relation表的信息请求
// 获取关注数
struct RelationFollowCountRequest {
    1: required i64 user_id,
}
struct RelationFollowCountResponse {
    1: required i32 status_code, // 0:成功 other:失败
    2: optional string status_msg,
    3: optional i64 follow_count, //关注数
}

// 获取粉丝数
struct RelationFollowerCountRequest {
    1: required i64 user_id,
}
struct RelationFollowerCountResponse {
    1: required i32 status_code,   // 0:成功 other:失败
    2: optional string status_msg,
    3: optional i64 follower_count, //粉丝数
}

// 获取关注关系(user_id是否关注to_user_id)
struct RelationIsFollowRequest {
    1: required i64 user_id,
    2: required i64 to_user_id,
}
struct RelationIsFollowResponse {
    1: required i32 status_code,   // 0:成功 other:失败
    2: optional string status_msg,
    3: required bool is_follow, //是否关注
}


service RelationService {
    RelationAddResponse RelationAdd(1: RelationAddRequest request),
    RelationRemoveResponse RelationRemove(1: RelationRemoveRequest request),
    RelationFollowListResponse RelationFollowList(1: RelationFollowListRequest request),
    RelationFollowerListResponse RelationFollowerList(1: RelationFollowerListRequest request),
    RelationFriendListResponse RelationFriendList(1: RelationFriendListRequest request),
    RelationFollowCountResponse RelationFollowCount(1: RelationFollowCountRequest request),
    RelationFollowerCountResponse RelationFollowerCount(1: RelationFollowerCountRequest request),
    RelationIsFollowResponse RelationIsFollow(1: RelationIsFollowRequest request),
}
