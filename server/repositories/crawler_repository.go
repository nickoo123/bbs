package repositories

import (
	"bbs/model"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"gorm.io/gorm"
)

var CrawlerRepository = newCrawlerRepository()

func newCrawlerRepository() *crawlerRepository {
	return &crawlerRepository{}
}

type crawlerRepository struct {
}

func (r *crawlerRepository) Get(db *gorm.DB, id int64) *model.Crawler {
	ret := &model.Crawler{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *crawlerRepository) Take(db *gorm.DB, where ...interface{}) *model.Crawler {
	ret := &model.Crawler{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *crawlerRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.Crawler) {
	cnd.Find(db, &list)
	return
}

func (r *crawlerRepository) FindOne(db *gorm.DB, cnd *sqls.Cnd) *model.Crawler {
	ret := &model.Crawler{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *crawlerRepository) FindPageByParams(db *gorm.DB, params *params.QueryParams) (list []model.Crawler, paging *sqls.Paging) {
	return r.FindPageByCnd(db, &params.Cnd)
}

func (r *crawlerRepository) FindPageByCnd(db *gorm.DB, cnd *sqls.Cnd) (list []model.Crawler, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Crawler{})

	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *crawlerRepository) Count(db *gorm.DB, cnd *sqls.Cnd) int64 {
	return cnd.Count(db, &model.Crawler{})
}

func (r *crawlerRepository) Create(db *gorm.DB, t *model.Crawler) (err error) {
	err = db.Create(t).Error
	return
}

func (r *crawlerRepository) Update(db *gorm.DB, t *model.Crawler) (err error) {
	err = db.Save(t).Error
	return
}

func (r *crawlerRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Crawler{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *crawlerRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Crawler{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *crawlerRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Crawler{}, "id = ?", id)
}
