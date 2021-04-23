package dao

import (
	"context"
	"helloworld2/internal/model"
)

type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	// Article(c context.Context, id int64) (*model.Article, error)
	//新增接口
	AddUser(c context.Context, nickname string, age int32) (user *model.User, err error)
	UpdateUser(c context.Context, uid int64, nickname string, age int32) (row int64, err error)
	GetUser(c context.Context, uid int64) (user *model.User, err error)
	GetUserList(c context.Context) (userlist []*model.User, err error)
}
