package parser

import (
	"crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseArticle(t *testing.T) {
	contents, err := ioutil.ReadFile("article_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseArticle(contents)
	if len(result.Items) != 1 {
		t.Errorf("result should contain1 1element,but was %v", result.Items)
	}
	profile := result.Items[0].(model.Profile)
	profile.UpdatedAt = "2021 年 5 月 28 日"
	expected := model.Profile{
		UpdatedAt: "2021 年 5 月 28 日",
		Title:     " 如何使用 AWS Lambda 集成或 Lambda 授权方排除 API Gateway HTTP API 的权限错误？",
	}
	if profile != expected {
		t.Errorf("wrong,get %v", profile)
	}
}
