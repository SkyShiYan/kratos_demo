package data

import (
	"context"
	"helloworld2/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

type Articles struct {
	gorm.Model
	Title   string
	Content string
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper("data/greeter", logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) (int, error) {
	s := &Articles{
		Title:   g.Hello,
		Content: g.Hello + "-Content",
	}
	result := r.data.db.Create(&s)
	// 返回数据插入的主键
	return int(s.ID), result.Error
}

func (r *greeterRepo) UpdateGreeter(g *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) GetGreeter(ctx context.Context, g *biz.Greeter) (string, error) {
	s := &Articles{
		Title: g.Hello,
	}
	result := r.data.db.Where("title = ?", g.Hello).First(&s)
	return s.Content, result.Error
}
