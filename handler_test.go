package main

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/csdn/gen-go/tangerine/csdn"
	"testing"
)

func TestCSDNHandler_SaveArticle(t *testing.T) {
	handler := NewCSDNHandler()
	req := &csdn.SaveArticleRequest{
		ArticleId: 139729092,
		Content:   "测试内容6",
		Title:     "《周边商品价格1》",
		PubStatus: csdn.PubStatusDraft,
		Action:    csdn.SaveArticleAction_OnlyContent,
		Base:      nil,
	}
	resp, err := handler.SaveArticle(context.Background(), req)
	fmt.Println(resp, err)
}

func TestCSDNHandler_GetArticle(t *testing.T) {
	handler := NewCSDNHandler()
	req := &csdn.GetArticleRequest{
		ArticleId: 139729092,
		Base:      nil,
	}
	resp, err := handler.GetArticle(context.Background(), req)
	fmt.Println(resp, err)
}
