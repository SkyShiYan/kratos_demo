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
	r.log.Infof("asdasdad---CreateGreeter--%v", g.Hello)
	s := &Articles{
		Title:   g.Hello,
		Content: g.Hello + "-Content",
	}
	r.data.db.Create(s)
	return 0, nil
}

func (r *greeterRepo) UpdateGreeter(g *biz.Greeter) error {
	r.log.Infof("asdasdad---UpdateGreeter--%v", g.Hello)
	return nil
}
