package top

import (
	"encoding/json"
	"fmt"
	"github.com/shuyi-tangerine/csdn/gen-go/tangerine/csdn"
)

type SaveArticleHttpRequest struct {
	*Article
	Source    string `json:"source"`
	IsNew     int64  `json:"is_new"`
	ReadType  string `json:"readType"`
	PubStatus string `json:"pub_status"`
}

func NewSaveArticleHttpRequest(req *csdn.SaveArticleRequest) (saveArticleHttpRequest *SaveArticleHttpRequest) {
	return &SaveArticleHttpRequest{
		Article: &Article{
			ID:               fmt.Sprintf("%d", req.ArticleId),
			ArticleID:        fmt.Sprintf("%d", req.ArticleId),
			Title:            req.Title,
			MarkdownContent:  req.Content,
			Content:          req.Content,
			Level:            "1",
			Tags:             "",
			Status:           2,
			Categories:       "",
			Type:             "original",
			OriginalLink:     "",
			AuthorizedStatus: false,
			Description:      "",
			ResourceUrl:      "",
			NotAutoSaved:     "1",
			CoverImages:      []string{},
			CoverType:        1,
			VoteID:           0,
			ResourceID:       "",
		},
		Source:    "pc_mdeditor",
		IsNew:     1,
		ReadType:  "public",
		PubStatus: req.PubStatus,
	}
}

func NewSaveArticleHttpRequestByArticle(req *Article, pubStatus string) (saveArticleHttpRequest *SaveArticleHttpRequest) {
	return &SaveArticleHttpRequest{
		Article:   req,
		Source:    "pc_mdeditor",
		IsNew:     1,
		ReadType:  "public",
		PubStatus: pubStatus,
	}
}

type SaveArticleHttpResponseBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewSaveArticleHttpResponseBody(rawBody string) (body *SaveArticleHttpResponseBody, err error) {
	var receiveBody SaveArticleHttpResponseBody
	err = json.Unmarshal([]byte(rawBody), &receiveBody)
	if err != nil {
		return
	}
	return &receiveBody, nil
}

type GetArticleHttpRequest struct {
	ID        string `json:"id"`
	ModelType string `json:"model_type"`
}

type GetArticleHttpResponseBody struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    *Article `json:"data"`
}

func NewGetArticleHttpResponseBody(rawBody string) (body *GetArticleHttpResponseBody, err error) {
	var receiveBody GetArticleHttpResponseBody
	err = json.Unmarshal([]byte(rawBody), &receiveBody)
	if err != nil {
		return
	}
	return &receiveBody, nil
}

type Article struct {
	ID               string   `json:"id"`
	ArticleID        string   `json:"article_id"`
	Title            string   `json:"title"`
	MarkdownContent  string   `json:"markdowncontent"`
	Content          string   `json:"content"`
	Level            string   `json:"level"`
	Tags             string   `json:"tags"`
	Status           int64    `json:"status"`
	Categories       string   `json:"categories"`
	Type             string   `json:"type"`
	OriginalLink     string   `json:"original_ink"`
	AuthorizedStatus bool     `json:"authorized_status"`
	Description      string   `json:"Description"`
	ResourceUrl      string   `json:"resource_url"`
	NotAutoSaved     string   `json:"not_auto_saved"`
	CoverImages      []string `json:"cover_images"`
	CoverType        int64    `json:"cover_type"`
	VoteID           int64    `json:"vote_id"`
	ResourceID       string   `json:"resource_id"`
}
