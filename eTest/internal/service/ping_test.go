package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/gomsx/goms/eTest/internal/dao/mock"
	m "github.com/gomsx/goms/eTest/internal/model"

	"github.com/golang/mock/gomock"
)

func TestHandPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dao := mock.NewMockDao(ctrl)
	//
	svc := service{dao: dao}
	ping := &m.Ping{Type: "http", Count: 2}
	want := &m.Ping{Type: "http", Count: 3}
	//1
	dao.EXPECT().
		ReadPing(ctxb, ping.Type).
		Return(ping, nil)

	dao.EXPECT().
		UpdatePing(ctxb, ping).
		Return(nil)
	//2
	dao.EXPECT().
		ReadPing(ctxb, ping.Type).
		Return(ping, errx)
	//3
	dao.EXPECT().
		ReadPing(ctxb, ping.Type).
		Return(ping, nil)

	dao.EXPECT().
		UpdatePing(ctxb, ping).
		Return(errx)
	//
	type args struct {
		c context.Context
		p *m.Ping
	}
	tests := []struct {
		name    string
		args    args
		want    *m.Ping
		wantErr bool
	}{
		{name: "Hand ping when read succeeded and update succeeded", args: args{ctxb, ping}, want: want, wantErr: false},
		{name: "Hand ping when read failed", args: args{ctxb, ping}, want: nil, wantErr: true},
		{name: "Hand ping when read succeeded and update failed", args: args{ctxb, ping}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := svc.HandPing(tt.args.c, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.HandPing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.HandPing() = %v, want %v", got, tt.want)
			}
		})
	}
}
