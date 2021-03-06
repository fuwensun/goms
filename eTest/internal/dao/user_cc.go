package dao

import (
	"context"
	"fmt"
	"strconv"

	m "github.com/gomsx/goms/eTest/internal/model"

	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

// existUserCC check user from cache.
func (d *dao) existUserCC(ctx context.Context, uid int64) (bool, error) {
	cc := d.redis
	key := getRedisKey(uid)
	exist, err := redis.Bool(cc.Do("EXISTS", key))
	if err != nil {
		err = fmt.Errorf("cc do EXISTS: %w", err)
		return exist, err
	}
	log.Debugf("cc %v exist user, uid: %v", exist, uid)
	return exist, nil
}

// setUserCC set user to cache.
func (d *dao) setUserCC(ctx context.Context, user *m.User) error {
	cc := d.redis
	key := getRedisKey(user.Uid)
	if _, err := cc.Do("HMSET", redis.Args{}.Add(key).AddFlat(user)...); err != nil {
		err = fmt.Errorf("cc do HMSET: %w", err)
		return err
	}
	log.Debugf("cc set user: %v", *user)
	return nil
}

// getUserCC get user from cache.
func (d *dao) getUserCC(ctx context.Context, uid int64) (*m.User, error) {
	cc := d.redis
	user := &m.User{}
	key := getRedisKey(uid)
	value, err := redis.Values(cc.Do("HGETALL", key))
	if err != nil {
		err = fmt.Errorf("cc do HGETALL: %w", err)
		return user, err
	}
	if err = redis.ScanStruct(value, user); err != nil {
		err = fmt.Errorf("cc ScanStruct: %w", err)
		return user, err
	}
	log.Debugf("cc get user: %v", *user)
	return user, nil
}

// delUserCC delete user from cache.
func (d *dao) delUserCC(ctx context.Context, uid int64) error {
	cc := d.redis
	key := getRedisKey(uid)
	if _, err := cc.Do("DEL", key); err != nil {
		err = fmt.Errorf("cc do DEL: %w", err)
		return err
	}
	log.Debugf("cc delete user, uid: %v", uid)
	return nil
}

//
func getRedisKey(uid int64) string {
	return "uid#" + strconv.FormatInt(uid, 16)
}
