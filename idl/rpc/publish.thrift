/*
IDL 注解说明
https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/usage/annotation/
*/
namespace go publish

include "./common.thrift"

struct PublishActionRequest {
    1: required i64 user_id,  // 用户id
    2: required binary data,  // 视频数据
    3: required string title, // 视频标题
}

struct PublishActionResponse {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
}

struct PublishListRequest {
    1: required i64 user_id, // 用户id
}

struct PublishListResponse {
    1: required i32 status_code,               // 状态码，0-成功，其他值-失败
    2: optional string status_msg,             // 返回状态描述
    3: optional list<common.Video> video_list, // 用户发布的视频列表
}

struct PublishVideoInfoRequest {
    1: optional i64 user_id,  // 发送请求的用户id
    2: required i64 video_id, // 视频id
}

struct PublishVideoInfoResponse {
    1: required i32 status_code,    // 状态码，0-成功，其他值-失败
    2: optional string status_msg,  // 返回状态描述
    3: required common.Video video, // 视频信息
}

struct PublishWorkCountRequest {
    1: required i64 user_id, // 发送请求的用户id
}

struct PublishWorkCountResponse {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: optional i64 work_count,    // 视频作品数量
}

service PublishService {
    PublishActionResponse PublishAction(1: PublishActionRequest req),          // 处理登录用户视频上传
    PublishListResponse PublishList(1: PublishListRequest req),                // 列出登录用户上传视频
    PublishVideoInfoResponse PublishVideoInfo(1: PublishVideoInfoRequest req), // 接口，根据视频Id返回视频信息
    PublishWorkCountResponse PublishWorkCount(1: PublishWorkCountRequest req), // 接口，根据用户Id返回作品数量
}
