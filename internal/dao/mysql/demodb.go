package mysql

import (
	"context"
	"github.com/coding-standard/golang-project-layout/internal/model"
	"gorm.io/gorm"
)

type DemoDbDao struct {
	Db *gorm.DB
}

func (d DemoDbDao) DemoDb(ctx context.Context, demoDb string) (*model.DemoDb, error) {
	dDb := &model.DemoDb{
		DemoDb: demoDb,
	}
	// Change GORM logic here
	if err := d.Db.Where("demo_db LIKE ?", "%"+demoDb+"%").Find(&dDb).Error; err != nil {
		return nil, err
	}
	return dDb, nil
}

func NewDemoDbDao(d *gorm.DB) *DemoDbDao {
	return &DemoDbDao{d}
}
