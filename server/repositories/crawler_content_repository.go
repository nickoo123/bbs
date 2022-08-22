package repositories

import (
	"bbs/model"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"gorm.io/gorm"
)

var CrawlerContentRepository = newCrawlerContentRepository()

func newCrawlerContentRepository() *crawlerContentRepository {
	return &crawlerContentRepository{}
}

type crawlerContentRepository struct {
}

func (r *crawlerContentRepository) Get(db *gorm.DB, id int64) *model.CrawlerContent {
	ret := &model.CrawlerContent{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *crawlerContentRepository) Take(db *gorm.DB, where ...interface{}) *model.CrawlerContent {
	ret := &model.CrawlerContent{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *crawlerContentRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.CrawlerContent) {
	cnd.Find(db, &list)
	return
}

func (r *crawlerContentRepository) FindOne(db *gorm.DB, cnd *sqls.Cnd) *model.CrawlerContent {
	ret := &model.CrawlerContent{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *crawlerContentRepository) FindPageByParams(db *gorm.DB, params *params.QueryParams) (list []model.CrawlerContent, paging *sqls.Paging) {
	return r.FindPageByCnd(db, &params.Cnd)
}

func (r *crawlerContentRepository) FindPageByCnd(db *gorm.DB, cnd *sqls.Cnd) (list []model.CrawlerContent, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.CrawlerContent{})

	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *crawlerContentRepository) Count(db *gorm.DB, cnd *sqls.Cnd) int64 {
	return cnd.Count(db, &model.CrawlerContent{})
}

func (r *crawlerContentRepository) Create(db *gorm.DB, t *model.CrawlerContent) (err error) {
	err = db.Create(t).Error
	return
}

func (r *crawlerContentRepository) Update(db *gorm.DB, t *model.CrawlerContent) (err error) {
	err = db.Save(t).Error
	return
}

func (r *crawlerContentRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.CrawlerContent{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *crawlerContentRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.CrawlerContent{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *crawlerContentRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.CrawlerContent{}, "id = ?", id)
}
