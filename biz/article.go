package biz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/shuyi-tangerine/csdn/gen-go/tangerine/csdn"
	"github.com/shuyi-tangerine/csdn/top"
	"io"
	"net/http"
)

type ArticleService struct {
	GetURL            string
	SaveURL           string
	DefaultHttpClient http.Client
}

func NewArticleService() (articleService *ArticleService) {
	return &ArticleService{
		GetURL:            "https://bizapi.csdn.net/blog-console-api/v3/editor/getArticle?id=%s&model_type=%s",
		SaveURL:           "https://bizapi.csdn.net/blog-console-api/v3/mdeditor/saveArticle",
		DefaultHttpClient: http.Client{},
	}
}

// SaveDraftContent 仅仅保存内容
func (m *ArticleService) SaveDraftContent(ctx context.Context, articleID string, content string) (err error) {
	article, err := m.Get(ctx, &top.GetArticleHttpRequest{
		ID:        articleID,
		ModelType: "",
	})
	if err != nil {
		return
	}

	article.ID, article.ArticleID = articleID, articleID
	article.MarkdownContent, article.Content = content, content
	req := top.NewSaveArticleHttpRequestByArticle(article, csdn.PubStatusDraft)
	err = m.Save(ctx, req)
	return
}

func (m *ArticleService) Save(ctx context.Context, req *top.SaveArticleHttpRequest) (err error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, m.SaveURL, bytes.NewBuffer(bts))
	if err != nil {
		return
	}

	// 从文件中取出 header，这里可以考虑加个单机缓存，或者换种方式存取
	headers, err := top.GetHttpHeaders(top.HttpHeaderTypeSaveArticle)
	if err != nil {
		return
	}
	for k, v := range headers {
		httpReq.Header.Set(k, v)
	}
	httpResp, err := m.DefaultHttpClient.Do(httpReq)
	if err != nil {
		return
	}

	// 保存接口，应该只存在 200 ok 和其它的不成功状态
	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code[%d]!=200, status=%s", httpResp.StatusCode, httpResp.Status)
	}

	rawBody, err := io.ReadAll(httpResp.Body)
	httpResp.Body.Close()
	if err != nil {
		return
	}

	body, err := top.NewSaveArticleHttpResponseBody(string(rawBody))
	if err != nil {
		return
	}

	// 成功
	if body.Code == 200 {
		return
	}

	return fmt.Errorf("save error, raw_body:%s", string(rawBody))
}

func (m *ArticleService) Get(ctx context.Context, req *top.GetArticleHttpRequest) (article *top.Article, err error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, m.buildGetURL(ctx, req.ID, req.ModelType), nil)
	if err != nil {
		return
	}

	// 从文件中取出 header，这里可以考虑加个单机缓存，或者换种方式存取
	headers, err := top.GetHttpHeaders(top.HttpHeaderTypeGetArticle)
	if err != nil {
		return
	}
	for k, v := range headers {
		httpReq.Header.Set(k, v)
	}
	httpResp, err := m.DefaultHttpClient.Do(httpReq)
	if err != nil {
		return
	}

	// 保存接口，应该只存在 200 ok 和其它的不成功状态
	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code[%d]!=200, status=%s", httpResp.StatusCode, httpResp.Status)
	}

	rawBody, err := io.ReadAll(httpResp.Body)
	httpResp.Body.Close()
	if err != nil {
		return
	}

	body, err := top.NewGetArticleHttpResponseBody(string(rawBody))
	if err != nil {
		return
	}

	// 成功
	if body.Code == 200 {
		return body.Data, nil
	}

	return nil, fmt.Errorf("get error, raw_body:%s", string(rawBody))
}

func (m *ArticleService) buildGetURL(ctx context.Context, articleID string, modelType string) string {
	return fmt.Sprintf(m.GetURL, articleID, modelType)
}
