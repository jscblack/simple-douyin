namespace go comment

include "./common.thrift"

// 评论操作
struct AddCommentActionRequest {
    1: required i64 user_id,         // 用户id
    2: required i64 video_id,        // 视频id
    3: optional string comment_text, // 用户填写的评论内容
}

struct AddCommentActionResponse {
    1: required i32 status_code,        // 0-成功，1-失败
    2: optional string status_msg,      // 错误信息
    3: optional common.Comment comment, // 评论成功返回评论内容
}

// 销评操作
struct DelCommentActionRequest {
    1: optional i64 comment_id, // 要删除的评论id
}

struct DelCommentActionResponse {
    1: required i32 status_code,   // 0-成功，1-失败
    2: optional string status_msg, // 错误信息
}

// 评论列表
struct CommentListRequest {
    1: required i64 user_id,  // 用户id
    2: required i64 video_id, // 视频id
}

struct CommentListResponse {
    1: required i32 status_code,                   // 状态码，0-成功，其他值-失败
    2: optional string status_msg,                 // 返回状态描述
    3: required list<common.Comment> comment_list, // 评论列表
}

//内部rpc
// 获取被评论数
struct CommentCountRequest {
    1: required i64 video_id, // 视频id
}

struct CommentCountResponse {
    1: required i32 status_code,   // 0:成功 other:失败
    2: optional string status_msg,
    3: optional i64 comment_count,   //被评论数
}

service CommentService {
    AddCommentActionResponse AddCommentAction(1: AddCommentActionRequest req),
    DelCommentActionResponse DelCommentAction(1: DelCommentActionRequest req),
    CommentListResponse CommentList(1: CommentListRequest req),
    CommentCountResponse CommentCount(1: CommentCountRequest req),
}