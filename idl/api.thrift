namespace go api

/*-----------------------------------基础接口-----------------------------------*/

/*-----------------------------------互动接口-----------------------------------*/

// 赞操作
struct FavoriteActionRequest {
    1: string token (api.query = "token"),
    2: i64 video_id (api.query = "video_id"),
    3: i32 action_type (api.query = "action_type"),
}

struct FavoriteActionResponse {
    1: i32 status_code,
    2: string status_msg,
}

// 喜欢列表
struct FavoriteListRequest {
    1: i64 user_id (api.query = "user_id"),
    2: string token (api.query = "token"),
}

struct FavoriteListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: list<Video> video_list,
}

service FavoriteService {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req) (api.post = "/douyin/favorite/action/"),
    FavoriteListResponse FavoriteList(1: FavoriteListRequest req) (api.get = "/douyin/favorite/list/"),
}

// 评论操作
struct Comment {
    1: i64 id,             // 视频评论id
    2: User user,          // 评论用户信息
    3: string content,     // 评论内容
    4: string create_date, // 评论发布日期，格式 mm-dd
}

struct CommentActionRequest {
    1: string token (api.query = "token"),
    2: i64 video_id (api.query = "video_id"),
    3: i32 action_type (api.query = "action_type"),
    4: string comment_text (api.query = "comment_text"),
    5: i64 comment_id (api.query = "comment_id"),
}

struct CommentActionResponse {
    1: i32 status_code,
    2: string status_msg,
    3: Comment comment,
}

// 评论列表
struct CommentListRequest {
    1: string token (api.query = "token"),
    2: i64 video_id (api.query = "video_id"),
}

struct CommentListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: list<Comment> comment_list,
}

service CommentService {
    CommentActionResponse CommentAction(1: CommentActionRequest req) (api.post = "/douyin/comment/action/"),
    CommentListResponse CommentList(1: CommentListRequest req) (api.get = "/douyin/comment/list/"),
}

/*-----------------------------------社交接口-----------------------------------*/

// 关注操作
struct RelationActionRequest {
    1: string token (api.query = "token"),
    2: i64 to_user_id (api.query = "to_user_id"),
    3: i32 action_type (api.query = "action_type"),
}

struct RelationActionResponse {
    1: i32 status_code,
    2: string status_msg,
}

// 关注列表
struct RelationFollowListRequest {
    1: i64 user_id (api.query = "user_id"),
    2: string token (api.query = "token"),
}

struct RelationFollowListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: list<User> user_list,
}

// 粉丝列表
struct RelationFollowerListRequest {
    1: i64 user_id (api.query = "user_id"),
    2: string token (api.query = "token"),
}

struct RelationFollowerListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: list<User> user_list,
}

// 好友列表
struct RelationFriendListRequest {
    1: i64 user_id (api.query = "user_id"),
    2: string token (api.query = "token"),
}

struct RelationFriendListResponse {
    1: i32 status_code,
    2: string status_msg,
    3: list<User> user_list,
}

service RelationService {
    RelationActionResponse RelationAction(1: RelationActionRequest req) (api.post = "/douyin/relation/action/"),
    RelationFollowListResponse RelationFollowList(1: RelationFollowListRequest req) (api.get = "/douyin/relation/follow/list/"),
    RelationFollowerListResponse RelationFollowerList(1: RelationFollowerListRequest req) (api.get = "/douyin/relation/follower/list/"),
    RelationFriendListResponse RelationFriendList(1: RelationFriendListRequest req) (api.get = "/douyin/relation/friend/list/"),
}

// 发送消息
struct MessageActionRequest {
    1: string token (api.query = "token"),
    2: i64 to_user_id (api.query = "to_user_id"),
    3: i32 action_type (api.query = "action_type"),
    4: string content (api.query = "content"),
}

struct MessageActionResponse {
    1: i32 status_code,
    2: string status_msg,
}

struct MessageChatRequest {
    1: string token (api.query = "token"),
    2: i64 to_user_id (api.query = "to_user_id"),
    3: i64 pre_msg_time (api.query = "pre_msg_time"),
}

struct MessageChatResponse {
    1: i32 status_code,
    2: string status_msg,
    3: list<Message> message_list,
}

service MeassgeService {
    MessageActionResponse MessageAction(1: MessageActionRequest req) (api.post = "/douyin/message/action/"),
    MessageChatResponse MessageChat(1: MessageChatRequest req) (api.get = "/douyin/message/chat/"),
}
