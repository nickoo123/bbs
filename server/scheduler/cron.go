package scheduler

import (
	"bbs/model"
	"github.com/mlogclub/simple/common/dates"
	"strings"

	//"bbs/model"
	//"bbs/model/constants"
	"bbs/pkg/config"
	xhttp "bbs/pkg/librarys/net/http"
	"bbs/pkg/sitemap"
	//"bbs/spam"
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/mlogclub/simple/sqls"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"bbs/services"
)

const (
	// 请求失败重试次数
	RETRY = 5
)

func Start() {
	c := cron.New()

	// crawler url data
	addCronFunc(c, "@every 10s", func() {
		userIds := config.Instance.CrawlerUser.UserId
		rand.Seed(time.Now().Unix())
		rands := rand.Int63n(config.Instance.CrawlerUser.UserSize)
		userId, _ := strconv.ParseInt(fmt.Sprintf("%v", userIds[rands]), 10, 64)
		craw := services.CrawlerService.FindOne(sqls.NewCnd().Where("status=?", 1))
		if craw != nil {
			doc, _, err := newHtml(craw.Url, craw.Charset, true)
			if err != nil {
				return
			}
			doc.Find(craw.CrawlerSelector).Each(func(i int, s *goquery.Selection) {
				link, ok := s.Find(craw.CrawlerSubSelector).Attr(craw.CrawlerSubAttr)
				if ok && !strings.Contains(link, craw.CrawlerSubExcludeAttr) {
					res := services.CrawlerContentService.FindOne(sqls.NewCnd().Where("link=?", link))
					if res == nil {
						title, _ := s.Find(craw.CrawlerSubSelector).Html()
						t := &model.CrawlerContent{
							CrawlerId:  craw.Id,
							Title:      title,
							Link:       link,
							Status:     1,
							CreateTime: dates.NowTimestamp(),
						}
						services.CrawlerContentService.Create(t)
					}
				}
			})
			news := services.CrawlerContentService.FindOne(sqls.NewCnd().Where("status=?", 1))
			if news != nil {
				time.Sleep(20)
				subdoc, _, _ := newHtml(news.Link, craw.Charset, true)
				cont, err := subdoc.Find(craw.CrawlerLinkSelector).Html()
				if err == nil {
					fmt.Println(userId, cont)
					//form := model.CreateArticleForm{
					//	Title:       news.Title,
					//	Summary:     summary,
					//	Content:     content,
					//	ContentType: constants.ContentTypeMarkdown,
					//	Tags:        tags,
					//}
					//
					//// 检测内容是否存在
					//user := services.UserService.Get(userId)
					//if err := spam.CheckArticle(user, form); err != nil {
					//	return
					//}
					//
					//// 发布内容
					//services.ArticleService.Publish(userId, form)
					//
					//// 更新采集状态
					//newconts := make(map[string]interface{})
					//newconts["Status"]		= 2
					//newconts["CrawlerNum"]	= int8(news.CrawlerNum)+1
					//services.CrawlerContentService.Updates(news.Id, newconts)
				}
			}
		}
	})

	// Generate RSS
	addCronFunc(c, "@every 30m", func() {
		services.ArticleService.GenerateRss()
		services.TopicService.GenerateRss()
		services.ProjectService.GenerateRss()
	})

	// Generate sitemap
	addCronFunc(c, "* 0 4 ? * *", func() {
		sitemap.Generate()
	})

	c.Start()
}

func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		logrus.Error(err)
	}
}

// 网页请求，失败重试
// 返回goquery
func newHtml(rawurl, charset string, isRedirect bool) (*goquery.Document, *http.Response, error) {
	var res []byte
	var resp *http.Response
	var body io.Reader
	var err error

	conf := &xhttp.ClientConfig{
		Timeout:   10 * time.Second,
		Dial:      10 * time.Second,
		KeepAlive: 60 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 存在重定向是否直接跳转
			if isRedirect {
				return nil
			}
			return http.ErrUseLastResponse
		},
	}
	c := xhttp.NewClient(conf)

	// 失败重试
	for i := 0; i < RETRY; i++ {
		//c.SetProxy(this.proxyFunc())
		res, resp, err = c.Get(context.TODO(), rawurl, nil)
		if err == nil {
			break
		}

		if resp != nil && (resp.StatusCode == 301 || resp.StatusCode == 302) {
			break
		}

		logrus.Errorf("请求失败：", rawurl, err)

		// 休眠10ms，防止采集速度过快被屏蔽
		time.Sleep(time.Duration(10) * time.Millisecond)
	}

	if err != nil {
		return nil, resp, err
	}

	body = bytes.NewReader(res)

	// 编码转换
	if charset != "UTF-8" {
		enc := mahonia.NewDecoder("GB18030")
		body = enc.NewReader(body)
	}

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, resp, err
	}

	return doc, resp, nil
}
