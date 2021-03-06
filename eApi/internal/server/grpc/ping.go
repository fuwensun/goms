package grpc

import (
	"context"

	api "github.com/gomsx/goms/eApi/api/v1"
	m "github.com/gomsx/goms/eApi/internal/model"
	e "github.com/gomsx/goms/pkg/err"

	log "github.com/sirupsen/logrus"
)

// setPingReplyMate set mate data to ping reply.
func setPingReplyMate(r *api.PingReply, ecode int64, err error) {
	r.Code = ecode
	if err != nil {
		r.Msg = err.Error()
	}
	r.Msg = "ok"
}

// Ping ping server.
func (s *Server) Ping(ctx context.Context, in *api.PingReq) (*api.PingReply, error) {
	svc := s.svc
	res := &api.PingReply{Data: &api.PingMsg{}}
	msg := ""
	if in.Data != nil {
		msg = in.Data.Message
	}

	//
	ping := &m.Ping{Type: "grpc"}
	ping, err := svc.HandPing(ctx, ping)
	if err != nil {
		setPingReplyMate(res, e.StatusInternalServerError, err)
		log.Infof("failed to hand ping, error: %v", err)
		return res, err
	}
	//
	res.Data.Message = m.MakePongMsg(msg)
	res.Data.Count = ping.Count
	setPingReplyMate(res, e.StatusOK, nil)
	log.Debugf("pong msg: %v, count: %v", res.Data.Message, res.Data.Count)
	return res, nil
}
