namespace go common

struct User {
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
}

struct Video {
    1: required i64 id,             // 视频唯一标识
    2: required User author,        // 视频作者信息
    3: required string play_url,    // 视频播放地址
    4: required string cover_url,   // 视频封面地址
    5: required i64 favorite_count, // 视频的点赞总数
    6: required i64 comment_count,  // 视频的评论总数
    7: required bool is_favorite,   // true-已点赞，false-未点赞
    8: required string title,       // 视频标题
}

struct Message {
    1: required i64 id,             // 消息id
    2: required i64 to_user_id,     // 该消息接收者的id
<<<<<<< Updated upstream
    3: required i64 from_user_id,   // 该消息发送者的id
=======
    3: required i64 user_id,        // 该消息发送者的id
>>>>>>> Stashed changes
    4: required string content,     // 消息内容
    5: optional string create_time, // 消息创建时间
}

struct Comment {
    1: required i64 id,             // 视频评论id
    2: required User user,          // 评论用户信息
    3: required string content,     // 评论内容
    4: required string create_date, // 评论发布日期，格式 mm-dd
}
