package model

import "strconv"

type User struct {
	Uid  int64  `redis:"uid" validate:"required,gte=0"`
	Name string `redis:"name" validate:"required,min=1,max=18"`
	Sex  int64  `redis:"sex" validate:"required,gte=1,lte=2"`
}

//
func GetRedisKey(uid int64) string {
	return "uid#" + strconv.FormatInt(uid, 10)
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
