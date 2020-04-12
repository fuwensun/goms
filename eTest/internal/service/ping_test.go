package service

import (
	"context"
	"testing"

	"github.com/fuwensun/goms/eTest/internal/dao/mock"
	. "github.com/fuwensun/goms/eTest/internal/model"

	"github.com/golang/mock/gomock"

	. "github.com/smartystreets/goconvey/convey"
)

//http
func TestUpdateHttpPingCount(t *testing.T) {
	Convey("TestUpdateHttpPingCount should return nil", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		daom := mock.NewMockDao(ctrl)
		svc := service{dao: daom}

		var pc PingCount = 2
		daom.EXPECT().
			UpdatePingCount(gomock.Any(), HTTP, pc).
			Return(nil)

		err := svc.UpdateHttpPingCount(context.Background(), pc)
		So(err, ShouldBeNil)
	})
}

func TestReadHttpPingCount(t *testing.T) {
	Convey("TestReadHttpPingCount ", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		daom := mock.NewMockDao(ctrl)
		svc := service{dao: daom}

		Convey("for succ", func() {
			var want PingCount = 2
			daom.EXPECT().
				ReadPingCount(gomock.Any(), HTTP).
				Return(want, nil)

			got, err := svc.ReadHttpPingCount(context.Background())
			So(got, ShouldEqual, want)
			So(err, ShouldBeNil)
		})
	})
}

//grpc
func TestUpdateGrpcPingCount(t *testing.T) {
	Convey("TestUpdateGrpcPingCount should return nil", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		daom := mock.NewMockDao(ctrl)
		svc := service{dao: daom}

		var pc PingCount = 2
		daom.EXPECT().
			UpdatePingCount(gomock.Any(), GRPC, pc).
			Return(nil)

		err := svc.UpdateGrpcPingCount(context.Background(), pc)
		So(err, ShouldBeNil)
	})
}

func TestReadGrpcPingCount(t *testing.T) {
	Convey("TestReadGrpcPingCount ", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		daom := mock.NewMockDao(ctrl)
		svc := service{dao: daom}

		Convey("for succ", func() {
			var want PingCount = 2
			daom.EXPECT().
				ReadPingCount(gomock.Any(), GRPC).
				Return(want, nil)

			got, err := svc.ReadGrpcPingCount(context.Background())
			So(got, ShouldEqual, want)
			So(err, ShouldBeNil)
		})
	})
}
