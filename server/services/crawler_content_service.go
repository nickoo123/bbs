package services

import (
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"

	"bbs/model"
	"bbs/repositories"
)

var CrawlerContentService = newCrawlerContentService()

func newCrawlerContentService() *crawlerContentService {
	return &crawlerContentService{}
}

type crawlerContentService struct {
}

func (s *crawlerContentService) Get(id int64) *model.CrawlerContent {
	return repositories.CrawlerContentRepository.Get(sqls.DB(), id)
}

func (s *crawlerContentService) Take(where ...interface{}) *model.CrawlerContent {
	return repositories.CrawlerContentRepository.Take(sqls.DB(), where...)
}

func (s *crawlerContentService) Find(cnd *sqls.Cnd) []model.CrawlerContent {
	return repositories.CrawlerContentRepository.Find(sqls.DB(), cnd)
}

func (s *crawlerContentService) FindOne(cnd *sqls.Cnd) *model.CrawlerContent {
	return repositories.CrawlerContentRepository.FindOne(sqls.DB(), cnd)
}

func (s *crawlerContentService) FindPageByParams(params *params.QueryParams) (list []model.CrawlerContent, paging *sqls.Paging) {
	return repositories.CrawlerContentRepository.FindPageByParams(sqls.DB(), params)
}

func (s *crawlerContentService) FindPageByCnd(cnd *sqls.Cnd) (list []model.CrawlerContent, paging *sqls.Paging) {
	return repositories.CrawlerContentRepository.FindPageByCnd(sqls.DB(), cnd)
}

func (s *crawlerContentService) Create(t *model.CrawlerContent) error {
	return repositories.CrawlerContentRepository.Create(sqls.DB(), t)
}

func (s *crawlerContentService) Update(t *model.CrawlerContent) error {
	return repositories.CrawlerContentRepository.Update(sqls.DB(), t)
}

func (s *crawlerContentService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.CrawlerContentRepository.Updates(sqls.DB(), id, columns)
}

func (s *crawlerContentService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.CrawlerContentRepository.UpdateColumn(sqls.DB(), id, name, value)
}

func (s *crawlerContentService) Delete(id int64) {
	repositories.CrawlerContentRepository.Delete(sqls.DB(), id)
}
