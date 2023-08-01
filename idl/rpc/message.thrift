namespace go message

include "./common.thrift"
//引入Message

//聊天记录
struct MessageChatRequest { 
    1: required i64 user_id, //己方用户id
    2: required i64 to_user_id, //对方用户id
    3: required i64 pre_msg_time, //上次最新消息的时间
}

struct MessageChatResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: optional list<common.Message> messages, //聊天记录列表
}

//聊天操作
struct MessageSendRequest {
    1: required i64 user_id,    //发送消息的用户id
    2: required i64 to_user_id, //接收消息的用户id
    3: required string content, //消息内容
}

struct MessageSendResponse {
    1: required i32 status_code, //0:成功 other:失败
    2: optional string status_msg,
}

//查询与该用户聊天双方的最新消息
struct MessageGetUserLastMessageRequest {
    1: required i64 user_id,  //调用者用户id
    2: required i64 to_user_id,  //调用者好友的用户id , 去db中查询双方的最新消息
}
struct MessageGetUserLastMessageResponse {
    1: required i32 status_code, //0:成功 other:失败
    2: optional string status_msg,
    3: optional common.Message message, //最新消息
}


service MessageService {
    MessageChatResponse MessageChat(1: MessageChatRequest request),
//    MessageActionResponse MessageAction(1: MessageActionRequest request),
    MessageSendResponse MessageSend(1: MessageSendRequest request),
    MessageGetUserLastMessageResponse MessageGetUserLastMessage(1: MessageGetUserLastMessageRequest request),
}