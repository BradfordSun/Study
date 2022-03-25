package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://aws.amazon.com/premiumsupport/knowledge-center/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 选择器不对的： rds部分(5)和ec2部分(2)，/cn/premiumsupport/knowledge-center/acm-troubleshoot-caa-errors/有2个a，(1)
	// 确实都是以/cn开头的，只有一个以https开头，这样的话应该有2498-5-2-1=2490个文章（算上https的）
	// 英文版2513-5-2-1=2505(算上https的)
	// 直接拿href和a.href相比，
	// 1. 有2个linux-static-hostname-suse，因为1个p里有2个a
	// 2. 多了个#EC2_Windows
	// 3. elb-fix-failing-health-checks-alb 一个p里面有2个a，误加了上面的elb-ecs-tasks-improperly-replaced到这个p里.这导致这里读了2个elb-ecs-tasks-improperly-replaced，下面的方法也是读了2个，但是读不出来elb-fix-failing-health-checks-alb
	// 4. premiumsupport/knowledge-center/opensearch-restore-data同理，但这个是因为上面的/elasticsearch-delegate-access-iam和/opensearch-delegate-access-iam/重复
	// 5.#Amazon_Aurora #Amazon_RDS_for_MySQL_and_MariaDB #Amazon_RDS_for_Oracle #Amazon_RDS_for_PostgreSQL #Amazon_RDS_for_SQL_Server都能读出来，下面的只能读到#Amazon_Aurora
	// 6. rds-postgresql-tune-query-performance/这里面有3个，就特么离谱
	// 7. simple-resource-record-route53-cli两个放到了一个p里。。导致下面的没有
	// 8. /sns-sms-delivery-delays/重复的，同理
	//9. security-network-acl-vpc-endpoint/五个在一个p，你敢信？导致下面直接少了4个
	var testslice []string
	doc.Find(".lb-col.lb-tiny-24.lb-mid-18 .lb-rtxt p a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		url, _ := s.Attr("href")
		if strings.HasPrefix(url, "/premiumsupport/knowledge-center/") || strings.HasPrefix(url, " https://aws.amazon.com/premiumsupport/knowledge-center/") {
			testslice = append(testslice, url)
		}
		//fmt.Printf("%v\n", testslice)
	})
	fmt.Printf("%v\n", testslice)
	fmt.Println(len(testslice))
	// 这个只有2478 但英文版有2496

	//doc.Find(".lb-col.lb-tiny-24.lb-mid-18 .lb-rtxt p").Each(func(i int, s *goquery.Selection) {
	//	// For each item found, get the title
	//	//title := s.Find("a").Text()
	//	url, _ := s.Find("a").Attr("href")
	//	fmt.Printf("%s\n", url)
	//})
}

func main() {
	ExampleScrape()
}
