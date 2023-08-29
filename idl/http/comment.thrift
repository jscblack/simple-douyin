namespace go comment

include "./common.thrift"

// 评论操作
struct CommentActionRequest {
    1: required string token (api.query = "token"),               // 用户鉴权token
    2: required i64 video_id (api.query = "video_id"),            // 视频id
    3: required i32 action_type (api.query = "action_type"),      // 1-发布评论，2-删除评论
    4: optional string comment_text (api.query = "comment_text"), // 用户填写的评论内容，在action_type=1的时候使用
    5: optional i64 comment_id (api.query = "comment_id"),        // 要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse {
    1: required i32 status_code,        // 0-成功，1-失败
    2: optional string status_msg,      // 错误信息
    3: optional common.Comment comment, // 评论成功返回评论内容，不需要重新拉取整个列表
}

// 评论列表
struct CommentListRequest {
    1: required string token (api.query = "token"),    // 用户鉴权token
    2: required i64 video_id (api.query = "video_id"), // 视频id
}

struct CommentListResponse {
    1: required i32 status_code,                   // 状态码，0-成功，其他值-失败
    2: optional string status_msg,                 // 返回状态描述
    3: optional list<common.Comment> comment_list, // 评论列表
}

service CommentService {
    CommentActionResponse CommentAction(1: CommentActionRequest req) (api.post = "/douyin/comment/action/"),
    CommentListResponse CommentList(1: CommentListRequest req) (api.get = "/douyin/comment/list/"),
}