package main

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/csdn/biz"
	"github.com/shuyi-tangerine/csdn/gen-go/base"
	"github.com/shuyi-tangerine/csdn/gen-go/tangerine/csdn"
	"github.com/shuyi-tangerine/csdn/top"
)

type CSDNHandler struct {
	ArticleService *biz.ArticleService
}

func NewCSDNHandler() (csdnHandler *CSDNHandler) {
	return &CSDNHandler{
		ArticleService: biz.NewArticleService(),
	}
}

func (m *CSDNHandler) SaveArticle(ctx context.Context, req *csdn.SaveArticleRequest) (resp *csdn.SaveArticleResponse, err error) {
	resp = csdn.NewSaveArticleResponse()
	resp.Base = base.NewRPCResponse()
	resp.Base.Code = 0
	if req.Action == csdn.SaveArticleAction_OnlyContent {
		err = m.ArticleService.SaveDraftContent(ctx, fmt.Sprintf("%d", req.ArticleId), req.Content)
		if err != nil {
			resp.Base.Code = -2
			resp.Base.Message = err.Error()
			err = nil
			return
		}
		return
	}
	saveReq := top.NewSaveArticleHttpRequest(req)
	err = m.ArticleService.Save(ctx, saveReq)
	if err != nil {
		resp.Base.Code = -1
		resp.Base.Message = err.Error()
		err = nil
		return
	}
	return
}

func (m *CSDNHandler) GetArticle(ctx context.Context, req *csdn.GetArticleRequest) (resp *csdn.GetArticleResponse, err error) {
	resp = csdn.NewGetArticleResponse()
	resp.Base = base.NewRPCResponse()
	resp.Base.Code = 0
	article, err := m.ArticleService.Get(ctx, &top.GetArticleHttpRequest{
		ID:        fmt.Sprintf("%d", req.ArticleId),
		ModelType: "",
	})
	if err != nil {
		resp.Base.Code = -1
		resp.Base.Message = err.Error()
		err = nil
		return
	}
	resp.ArticleInfo = csdn.NewArticleInfo()
	resp.ArticleInfo.Content = article.MarkdownContent
	return
}
