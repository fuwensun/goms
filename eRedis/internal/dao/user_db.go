package dao

import (
	"context"
	"fmt"

	m "github.com/aivuca/goms/eRedis/internal/model"
	"github.com/qiniu/x/log"
)

const (
	_createUser = "INSERT INTO user_table(uid,name,sex) VALUES(?,?,?)"
	_readUser   = "SELECT uid,name,sex FROM user_table WHERE uid=?"
	_updateUser = "UPDATE user_table SET name=?,sex=? WHERE uid=?"
	_deleteUser = "DELETE FROM user_table WHERE uid=?"
)

//
func (d *dao) createUserDB(c context.Context, user *m.User) error {
	db := d.db
	result, err := db.Exec(_createUser, user.Uid, user.Name, user.Sex)
	if err != nil {
		err = fmt.Errorf("db exec insert: %w", err)
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("db rows affected: %w", err)
		return err
	}
	log.Printf("db insert user = %v", *user)
	return nil
}

func (d *dao) readUserDB(c context.Context, uid int64) (*m.User, error) {
	db := d.db
	user := &m.User{}
	rows, err := db.Query(_readUser, uid)
	if err != nil {
		err = fmt.Errorf("db query: %w", err)
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&user.Uid, &user.Name, &user.Sex); err != nil {
			err = fmt.Errorf("db rows scan: %w", err)
			return nil, err
		}
		if rows.Next() {
			// uid 重复
			log.Printf("db read multiple uid")
		}
		log.Printf("db read user = %v", *user)
		return user, nil
	}
	//not found
	log.Printf("db not found user,uid = %v", user.Uid)
	return user, nil
}

func (d *dao) updateUserDB(c context.Context, user *m.User) error {
	db := d.db
	result, err := db.Exec(_updateUser, user.Name, user.Sex, user.Uid)
	if err != nil {
		err = fmt.Errorf("db exec update: %w", err)
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("db rows affected: %w", err)
		return err
	}
	log.Printf("db update user = %v", *user)
	return nil
}

func (d *dao) deleteUserDB(c context.Context, uid int64) error {
	db := d.db
	result, err := db.Exec(_deleteUser, uid)
	if err != nil {
		err = fmt.Errorf("db exec delete: %w", err)
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("db rows affected: %w", err)
		return err
	}
	log.Printf("db delete user, uid = %v", uid)
	return nil
}