package grpc

import (
	"context"
	"log"

	"github.com/aivuca/goms/eRedis/api"
	m "github.com/aivuca/goms/eRedis/internal/model"
	e "github.com/aivuca/goms/pkg/err"
	ms "github.com/aivuca/goms/pkg/misc"
)

// Ping ping server.
func (s *Server) Ping(ctx context.Context, req *api.Request) (*api.Reply, error) {
	svc := s.svc
	//
	var res *api.Reply
	ping := &m.Ping{Type: "grpc"}
	ping, err := svc.HandPing(ctx, ping)
	if err != nil {
		res = &api.Reply{
			Message: e.ErrInternalError.Error(),
		}
		return res, err
	}
	//
	res = &api.Reply{
		Message: ms.MakePongMsg(req.Message),
		Count:   ping.Count,
	}
	log.Printf("pong msg: %v, count: %v", res.Message, res.Count)
	return res, nil
}
