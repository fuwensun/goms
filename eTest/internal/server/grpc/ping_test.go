package grpc

import (
	"testing"

	"github.com/gomsx/goms/eTest/api"
	m "github.com/gomsx/goms/eTest/internal/model"
	"github.com/gomsx/goms/eTest/internal/service/mock"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svcm := mock.NewMockSvc(ctrl)
	//
	srv := Server{svc: svcm}
	ping := &m.Ping{Type: "grpc"}
	want := &m.Ping{Type: "grpc", Count: 5}

	Convey("Ping with message", t, func() {
		//mock
		svcm.EXPECT().
			HandPing(ctxa, ping).
			Return(want, nil)
		//构建 req
		req := &api.Request{
			Message: "xxx",
		}
		//发起 req
		res, err := srv.Ping(ctxb, req)
		//断言
		So(err, ShouldEqual, nil)
		So(res.Message, ShouldEqual, m.MakePongMsg(req.Message))
		So(res.Count, ShouldEqual, want.Count)
	})

	Convey("Ping when service error", t, func() {
		//mock
		svcm.EXPECT().
			HandPing(ctxa, ping).
			Return(want, errx)
		//构建 req
		req := &api.Request{}
		//发起 req
		_, err := srv.Ping(ctxb, req)
		//断言
		So(err, ShouldEqual, errx)
	})
}
