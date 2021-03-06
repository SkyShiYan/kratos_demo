package service

import (
	"context"

	v1 "helloworld2/api/helloworld/v1"
	"helloworld2/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper("service/greeter", logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	id, err := s.uc.Create(ctx, &biz.Greeter{Hello: in.GetName()})
	if err != nil {
		s.log.Errorf("插入数据失败 %v", err)
	} else {
		s.log.Infof("插入数据成功 %v", id)
	}
	// s.log.Infof("SayHello Received: %v", in.GetName())
	// s.uc.Update(&biz.Greeter{Hello: "ssss"})
	// if in.GetName() == "error" {
	// return nil, errors.NotFound("ErrorReason", in.GetName())
	// }
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) GetTest(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	content, err := s.uc.Get(ctx, &biz.Greeter{Hello: in.GetName()})
	if err == gorm.ErrRecordNotFound {
		return &v1.HelloReply{Message: "用户未找到"}, nil
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName() + "-Content:" + content}, err
}

// 保存数据
func (s *GreeterService) SetTest(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	content, err := s.uc.Get(ctx, &biz.Greeter{Hello: in.GetName()})
	if err == gorm.ErrRecordNotFound {
		return &v1.HelloReply{Message: "用户未找到"}, nil
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName() + "-Content:" + content}, err
}
