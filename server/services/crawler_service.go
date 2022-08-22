package services

import (
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"

	"bbs/model"
	"bbs/repositories"
)

var CrawlerService = newCrawlerService()

func newCrawlerService() *crawlerService {
	return &crawlerService{}
}

type crawlerService struct {
}

func (s *crawlerService) Get(id int64) *model.Crawler {
	return repositories.CrawlerRepository.Get(sqls.DB(), id)
}

func (s *crawlerService) Take(where ...interface{}) *model.Crawler {
	return repositories.CrawlerRepository.Take(sqls.DB(), where...)
}

func (s *crawlerService) Find(cnd *sqls.Cnd) []model.Crawler {
	return repositories.CrawlerRepository.Find(sqls.DB(), cnd)
}

func (s *crawlerService) FindOne(cnd *sqls.Cnd) *model.Crawler {
	return repositories.CrawlerRepository.FindOne(sqls.DB(), cnd)
}

func (s *crawlerService) FindPageByParams(params *params.QueryParams) (list []model.Crawler, paging *sqls.Paging) {
	return repositories.CrawlerRepository.FindPageByParams(sqls.DB(), params)
}

func (s *crawlerService) FindPageByCnd(cnd *sqls.Cnd) (list []model.Crawler, paging *sqls.Paging) {
	return repositories.CrawlerRepository.FindPageByCnd(sqls.DB(), cnd)
}

func (s *crawlerService) Create(t *model.Crawler) error {
	return repositories.CrawlerRepository.Create(sqls.DB(), t)
}

func (s *crawlerService) Update(t *model.Crawler) error {
	return repositories.CrawlerRepository.Update(sqls.DB(), t)
}

func (s *crawlerService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.CrawlerRepository.Updates(sqls.DB(), id, columns)
}

func (s *crawlerService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.CrawlerRepository.UpdateColumn(sqls.DB(), id, name, value)
}

func (s *crawlerService) Delete(id int64) {
	repositories.CrawlerRepository.Delete(sqls.DB(), id)
}
