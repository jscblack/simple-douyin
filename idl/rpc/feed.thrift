/*
IDL 注解说明
https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/usage/annotation/
*/
namespace go feed

include "./common.thrift"

// 视频流拉取
struct FeedRequest {
    1: optional i64 user_id, // 用户id
    2: optional i64 latest_time, // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
}

struct FeedResponse {
    1: required i32 status_code,      // 状态码，0-成功，其他值-失败
    2: optional string status_msg,    // 返回状态描述
    3: optional list<common.Video> video_list, // 视频列表
    4: optional i64 next_time,        // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

service FeedService {
    FeedResponse Feed(1: FeedRequest req),
}
