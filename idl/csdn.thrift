include "base.thrift"

namespace go tangerine.csdn

const string PubStatusDraft = "draft"

enum SaveArticleAction {
    All = 0,
    OnlyContent = 1
}

struct SaveArticleRequest {
    1: required i64 articleId
    2: required string content
    3: required string title
    4: required string pubStatus  // draft=草稿
    5: optional SaveArticleAction action = 0
    255: optional base.RPCRequest Base
}

struct SaveArticleResponse {
    255: required base.RPCResponse Base
}

struct ArticleInfo {
    1: required string content
}

struct GetArticleRequest {
    1: required i64 articleId
    255: optional base.RPCRequest Base
}

struct GetArticleResponse {
    1: optional ArticleInfo articleInfo
    255: required base.RPCResponse Base
}

service CSDNHandler {
    // 保存文章
    SaveArticleResponse SaveArticle(1:SaveArticleRequest req)
    // 查询文章
    GetArticleResponse GetArticle(1:GetArticleRequest req)
}

