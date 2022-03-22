package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseUrlList(t *testing.T) {
	contents, err := ioutil.ReadFile("urllist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseUrlList(contents)
	const resultSize = 2409
	expectUrls := []string{
		"/cn/premiumsupport/knowledge-center/?nc2=h_m_ma",
		"/cn/premiumsupport/knowledge-center/view-activate-credits/",
		"/cn/premiumsupport/knowledge-center/activate-business-support-charge/",
	}
	expectArticles := []string{
		"知识中心",
		"我收到一封电子邮件，里面提供了 AWS Activate 创建者或组合包信息。我可以从何处找到我的 AWS 促销积分？",
		"我拥有 AWS Activate 组合资源包，为什么还向我收取 AWS 商业支持费用？",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d link, but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s, but was %s", i, url, result.Requests[i].Url)
		}
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d link, but had %d", resultSize, len(result.Items))
	}
	for i, item := range expectArticles {
		if result.Items[i].(string) != item {
			t.Errorf("expected item #%d: %s, but was %s", i, item, result.Items[i].(string))
		}
	}
}
