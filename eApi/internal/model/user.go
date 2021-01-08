package model

import (
	ms "github.com/fuwensun/goms/pkg/misc"
)

//
type User struct {
	Uid  int64  `redis:"uid" validate:"required,gte=0"`
	Name string `redis:"name" validate:"required,min=1,max=18"`
	Sex  int64  `redis:"sex" validate:"required,gte=1,lte=2"`
}

// for test
//
func GetUser() *User {
	return &User{
		Uid:  ms.GetUid(),
		Name: ms.GetName(),
		Sex:  ms.GetSex(),
	}
}

// for cache
var expire int64 = 10

//
func GetExpire() int64 {
	return expire
}

//
func SetExpire(time int64) {
	expire = time
}
