package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase)

type User struct {
	Uid      int32
	Nickname string
	Age      int32
	Uptime   int32
	Addtime  int32
}
