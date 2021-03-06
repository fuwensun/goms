package dao

import (
	"context"
	"fmt"
	"log"

	m "github.com/gomsx/goms/eRedis/internal/model"
)

// Create create user
func (d *dao) CreateUser(ctx context.Context, user *m.User) error {
	if err := d.createUserDB(ctx, user); err != nil {
		err = fmt.Errorf("create user in db: %w", err)
		return err
	}
	return nil
}

// ReadUser read user
// Cache Aside 读策略
func (d *dao) ReadUser(ctx context.Context, uid int64) (*m.User, error) {
	// 1.先读 cache
	exist, err := d.existUserCC(ctx, uid)
	if err != nil {
		// 查询 cache 失败，返回 err
		return nil, err
	}
	if exist {
		// 查询 cache 成功，存在条目
		user, err := d.getUserCC(ctx, uid)
		if err != nil {
			// 读 cache 失败，返回 err
			err = fmt.Errorf("get user from cc: %w", err)
			return nil, err
		}
		// 读 cache 成功，返回 user
		return user, nil
	}

	// 2.再读 DB (cache 没命中)
	user, err := d.readUserDB(ctx, uid)
	if err != nil {
		// 读 DB 失败，返回 err
		err = fmt.Errorf("read user from db: %w", err)
		return nil, err
	}

	// 3.最后写 cache
	if err = d.setUserCC(ctx, user); err != nil {
		// 读 DB 成功，回种 cache 失败，返回 err
		log.Printf("faild to set user cc")
		err = fmt.Errorf("set user to cc: %w", err)
		return nil, err
	}

	// 读 DB 成功，回种 cache 成功，返回 user
	return user, nil
}

// UpdateUser update user
// Cache Aside 写策略(更新)
func (d *dao) UpdateUser(ctx context.Context, user *m.User) error {
	// 1.先更新 DB
	if err := d.updateUserDB(ctx, user); err != nil {
		err = fmt.Errorf("update user in db: %w", err)
		return err
	}
	// 2.再删除 cache
	if err := d.delUserCC(ctx, user.Uid); err != nil {
		// 缓存过期
		log.Printf("cache expiration, uid=%v, err=%v", user.Uid, err)
		err = fmt.Errorf("delete user in cc: %w", err)
		return err
	}
	return nil
}

// DeleteUser delete user
// Cache Aside 写策略(删除)
func (d *dao) DeleteUser(ctx context.Context, uid int64) error {
	// 1.先删除 DB
	if err := d.deleteUserDB(ctx, uid); err != nil {
		err = fmt.Errorf("delete user in db: %w", err)
		return err
	}
	// 2.再删除 cache
	if err := d.delUserCC(ctx, uid); err != nil {
		// 缓存过期
		log.Printf("cache expiration, uid=%v, err=%v", uid, err)
		err = fmt.Errorf("del user in cc: %w", err)
		return err
	}
	return nil
}
