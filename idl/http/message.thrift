namespace go message

include "./common.thrift"

//聊天记录
struct MessageChatRequest {
    1: required string token (api.query = "token"),
    2: required i64 to_user_id (api.query = "to_user_id"),
    3: required i64 pre_msg_time (api.query = "pre_msg_time"),
}

struct MessageChatResponse {
    1: required i32 status_code,               //状态码: 0:成功 other:失败
    2: optional string status_msg,             //错误信息
    3: optional list<common.Message> messages, //消息列表
}

//
struct MessageActionRequest {
    1: required string token (api.query = "token"),
    2: required i64 to_user_id (api.query = "to_user_id"),
    3: required i32 action_type (api.query = "action_type"), //1:发送消息
    4: required string content (api.query = "content"),      //消息内容
}

struct MessageActionResponse {
    1: required i32 status_code,   //状态码: 0:成功 other:失败
    2: optional string status_msg,
}

service MessageService {
    MessageChatResponse MessageChat(1: MessageChatRequest request) (api.get = "/douyin/message/chat/"),
    MessageActionResponse MessageAction(1: MessageActionRequest request) (api.post = "/douyin/message/action/"),
}
