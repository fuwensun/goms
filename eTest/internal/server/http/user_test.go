package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	m "github.com/aivuca/goms/eTest/internal/model"
	"github.com/aivuca/goms/eTest/internal/service/mock"

	. "bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}
	ctx := gomock.Any()
	errt := errors.New("error")

	router := gin.New()
	router.Use(setRequestId()) //request_id
	router.POST("/user", srv.createUser)

	Convey("createUser should respond http.StatusCreated", t, func() {
		//mock
		user := m.GetUser()
		Patch(m.GetUid, func() int64 {
			return user.Uid
		})
		svcm.EXPECT().
			CreateUser(ctx, user).
			Return(nil)

		//构建请求数据
		v := url.Values{}
		v.Set("name", user.Name)
		v.Set("sex", m.StrInt(user.Sex))
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("POST", "/user", reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))
		fmt.Println(" ==>", string(body))

		//解析 resp 到 map
		rm := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &rm)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", rm)

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusCreated)
		So(rm["name"], ShouldEqual, user.Name)
		So(rm["sex"], ShouldEqual, float64(user.Sex))
	})

	Convey("createUser should respond http.StatusBadRequest", t, func() {
		//mock
		user := m.GetUser()
		Patch(m.GetUid, func() int64 {
			return user.Uid
		})
		user.Sex = m.GetSexBad()

		//构建请求数据
		v := url.Values{}
		v.Set("name", user.Name)
		v.Set("sex", m.StrInt(user.Sex))
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("POST", "/user", reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))
		fmt.Println(" ==>", string(body))

		//解析 resp 到 map
		rm := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &rm)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", rm)

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
	})

	Convey("createUser should respond http.StatusInternalServerError", t, func() {
		//mock
		user := m.GetUser()
		Patch(m.GetUid, func() int64 {
			return user.Uid
		})
		svcm.EXPECT().
			CreateUser(ctx, user).
			Return(errt)

		//构建请求数据
		v := url.Values{}
		v.Set("name", user.Name)
		v.Set("sex", m.StrInt(user.Sex))
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("POST", "/user", reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(" ==>", resp.StatusCode)
		fmt.Println(" ==>", resp.Header.Get("Content-Type"))
		fmt.Println(" ==>", string(body))

		//解析 resp 到 map
		rm := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &rm)
		if err != nil {
			panic(err)
		}
		fmt.Println(" ==>", rm)

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestReadUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}
	ctx := gomock.Any()
	errt := errors.New("error")

	router := gin.New()
	router.Use(setRequestId())
	router.GET("/user/:uid", srv.readUser)

	Convey("readUser should respond http.StatusOK", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			ReadUser(ctx, user.Uid).
			Return(user, nil)

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/user/"+m.StrInt(user.Uid), nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		//解析 resp 到 map
		rm := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &rm)
		if err != nil {
			panic(err)
		}
		// fmt.Println(" ==>", rm)

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusOK)
		So(rm["uid"], ShouldEqual, float64(user.Uid))
		So(rm["name"], ShouldEqual, user.Name)
		So(rm["sex"], ShouldEqual, float64(user.Sex))
	})

	Convey("readUser should respond http.StatusBadRequest", t, func() {
		//mock
		user := m.GetUser()
		user.Uid = m.GetUidBad()

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/user/"+m.StrInt(user.Uid), nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		//解析 resp 到 map
		rm := make(map[string]interface{}, 4)
		err := json.Unmarshal([]byte(string(body)), &rm)
		if err != nil {
			panic(err)
		}
		// fmt.Println(" ==>", rm)

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
		So(rm["Uid"], ShouldEqual, float64(user.Uid))
	})

	Convey("readUser should respond http.StatusInternalServerError", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			ReadUser(ctx, user.Uid).
			Return(user, errt)

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/user/"+m.StrInt(user.Uid), nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
	})
}
func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}
	ctx := gomock.Any()
	errt := errors.New("error")

	router := gin.New()
	router.Use(setRequestId())
	router.PUT("/user/:uid", srv.updateUser)

	Convey("updateUser should respond http.StatusNoContent", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			UpdateUser(ctx, user).
			Return(nil)

		//构建请求数据
		v := url.Values{}
		v.Set("uid", m.StrInt(user.Uid))
		v.Set("name", user.Name)
		v.Set("sex", m.StrInt(user.Sex))
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/user/"+m.StrInt(user.Uid), reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusNoContent)
	})

	Convey("updateUser should respond http.StatusBadRequest", t, func() {
		//mock
		user := m.GetUser()
		user.Uid = m.GetUidBad()

		//构建请求数据
		v := url.Values{}
		v.Set("uid", m.StrInt(user.Uid))
		v.Set("name", user.Name)
		v.Set("sex", m.StrInt(user.Sex))
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/user/"+m.StrInt(user.Uid), reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
	})

	Convey("updateUser should respond http.StatusInternalServerError", t, func() {
		//mock
		user := m.GetUser()
		svcm.EXPECT().
			UpdateUser(ctx, user).
			Return(errt)

		//构建请求数据
		v := url.Values{}
		v.Set("uid", m.StrInt(user.Uid))
		v.Set("name", user.Name)
		v.Set("sex", m.StrInt(user.Sex))
		reader := ioutil.NopCloser(strings.NewReader(v.Encode()))

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/user/"+m.StrInt(user.Uid), reader)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		//断言
		So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)

	srv := Server{svc: svcm}
	ctx := gomock.Any()
	errt := errors.New("error")

	router := gin.New()
	router.Use(setRequestId())
	router.DELETE("/user/:uid", srv.deleteUser)

	Convey("deleteUser should respond http.StatusNoContent", t, func() {
		//mock
		uid := m.GetUid()
		svcm.EXPECT().
			DeleteUser(ctx, uid).
			Return(nil)

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/user/"+m.StrInt(uid), nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		// 断言
		So(resp.StatusCode, ShouldEqual, http.StatusNoContent)
	})

	Convey("deleteUser should respond http.StatusBadRequest", t, func() {
		//mock
		uid := m.GetUidBad()
		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/user/"+m.StrInt(uid), nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		// 断言
		So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
	})

	Convey("deleteUser should respond http.StatusInternalServerError", t, func() {
		//mock
		uid := m.GetUid()
		svcm.EXPECT().
			DeleteUser(ctx, uid).
			Return(errt)

		//构建请求
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("DELETE", "/user/"+m.StrInt(uid), nil)

		//发起req
		router.ServeHTTP(w, r)
		resp := w.Result()

		// 断言
		So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
	})
}
