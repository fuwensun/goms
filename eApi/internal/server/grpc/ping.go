package grpc

import (
	"context"

	api "github.com/aivuca/goms/eApi/api/v1"
	m "github.com/aivuca/goms/eApi/internal/model"
	e "github.com/aivuca/goms/pkg/err"
	ms "github.com/aivuca/goms/pkg/misc"

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
	res.Data.Message = ms.MakePongMsg(msg)
	res.Data.Count = ping.Count
	setPingReplyMate(res, e.StatusOK, nil)
	log.Debugf("pong msg: %v, count: %v", res.Data.Message, res.Data.Count)
	return res, nil
}
