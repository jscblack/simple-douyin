namespace go pong

struct PingReq {
    1: required string ping_time,
}
struct PongResp {
    1: required string pong_time,
}
service PongService {
    PongResp Pong(1: required PingReq req),
}
