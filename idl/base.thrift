namespace go base

struct RPCRequest {

}

struct RPCResponse {
    1: required i64 code
    2: required string message
}