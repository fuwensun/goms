package http

import (
	"net/http"

	m "github.com/gomsx/goms/eTest/internal/model"
	e "github.com/gomsx/goms/pkg/err"
	"github.com/gomsx/goms/pkg/id"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
	"github.com/unknwon/com"
)

// createUser create user.
func (s *Server) createUser(ctx *gin.Context) {
	svc := s.svc
	name := com.StrTo(ctx.PostForm("name")).String()
	sex := com.StrTo(ctx.PostForm("sex")).MustInt64()
	// 记录参数
	log.Infof("start to create user, name:%v, sex:%v", name, sex)

	user := &m.User{}
	user.Uid = id.GenUid()
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, GetValidateError(err))
		// 记录异常
		log.Infof("failed to validate data, user: %v, error: %v", *user, err)
		return
	}
	// 记录中间结果
	log.Infof("succeeded to create data, user: %v", *user)

	if err := svc.CreateUser(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Infof("failed to create user, error: %v", err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	// 记录返回结果
	log.Infof("succeeded to create user, user: %v", *user)
	return
}

// readUser read user.
func (s *Server) readUser(ctx *gin.Context) {
	svc := s.svc
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	if uid == 0 {
		uid = com.StrTo(ctx.Query("uid")).MustInt64()
	}
	log.Infof("start to read user, uid: %v", uid)

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ctx.JSON(http.StatusBadRequest, GetValidateError(err))
		log.Infof("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return
	}
	log.Infof("succeeded to create data, uid: %v", user.Uid)

	user, err := svc.ReadUser(ctx, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Infof("failed to read user, error: %v", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{ //read ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Infof("succeeded to read user, user: %v", *user)
	return
}

// updateUser update user.
func (s *Server) updateUser(ctx *gin.Context) {
	svc := s.svc
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	if uid == 0 {
		uid = com.StrTo(ctx.PostForm("uid")).MustInt64()
	}
	name := com.StrTo(ctx.PostForm("name")).String()
	sex := com.StrTo(ctx.PostForm("sex")).MustInt64()
	log.Infof("start to update user, name: %v, sex %v", name, sex)

	user := &m.User{}
	user.Uid = uid
	user.Name = name
	user.Sex = sex

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, GetValidateError(err))
		log.Infof("failed to validate data, user: %v, error: %v", *user, err)
		return
	}
	log.Infof("succeeded to create data, user: %v", *user)

	err := svc.UpdateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Infof("failed to update user, error: %v", err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{}) //update ok
	log.Infof("succeeded to update user, user: %v", *user)
	return
}

// deleteUser delete user.
func (s *Server) deleteUser(ctx *gin.Context) {
	svc := s.svc
	uid := com.StrTo(ctx.Param("uid")).MustInt64()
	log.Infof("start to delete user, uid: %v", uid)

	user := &m.User{}
	user.Uid = uid

	validate := validator.New()
	if err := validate.StructPartial(user, "Uid"); err != nil {
		ctx.JSON(http.StatusBadRequest, GetValidateError(err))
		log.Infof("failed to validate data, uid: %v, error: %v", user.Uid, err)
		return
	}
	log.Infof("succeeded to create data, uid: %v", user.Uid)

	err := svc.DeleteUser(ctx, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Infof("failed to delete user, error: %v", err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{}) //delete ok
	log.Infof("succeeded to delete user, user: %v", *user)
	return
}

// GetValidateError get validate error.
func GetValidateError(err error) *map[string]interface{} {
	ev := err.(validator.ValidationErrors)[0]
	field := ev.StructField()
	value := ev.Value()

	em := make(map[string]interface{})
	em["error"] = e.UserEcodeMap[field]
	em[field] = value
	return &em
}
