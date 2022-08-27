package dao

import (
	"context"
	"github.com/coding-standard/golang-project-layout/internal/model"
)

type DemoDbDao interface {
	DemoDb(ctx context.Context, demoDb string) (*model.DemoDb, error)
}
