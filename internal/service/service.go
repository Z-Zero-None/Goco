package service

import (
	"context"

	otgorm "github.com/eddycjy/opentracing-gorm"

	"Goco/global"
	"Goco/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func NewService(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.NewDao(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
