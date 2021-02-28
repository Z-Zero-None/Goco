package service

import (
	"context"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"

	"Goco/global"
)

type Service struct {
	ctx context.Context
	engine *gorm.DB
}

func NewService(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.engine=otgorm.WithContext(svc.ctx, global.DBEngine)
	return svc
}
