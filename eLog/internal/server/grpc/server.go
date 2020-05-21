package grpc

import (
	"net"
	"path/filepath"

	"github.com/fuwensun/goms/eLog/api"
	"github.com/fuwensun/goms/eLog/internal/service"
	"github.com/fuwensun/goms/pkg/conf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/rs/zerolog/log"
)

type config struct {
	Addr string `yaml:"addr"`
}

//
type Server struct {
	cfg *config
	gs  *grpc.Server
	svc service.Svc
}

func getConfig(cfgpath string) (config, error) {
	var cfg config
	path := filepath.Join(cfgpath, "grpc.yml")
	if err := conf.GetConf(path, &cfg); err != nil {
		log.Info().Msgf("get config file: %v", err)
	}
	if cfg.Addr != "" {
		log.Info().Msgf("get config addr: %v", cfg.Addr)
		return cfg, nil
	}
	//todo get env
	cfg.Addr = ":50051"
	log.Info().Msgf("use default addr: %v", cfg.Addr)
	return cfg, nil
}

//
func New(cfgpath string, s service.Svc) (*Server, error) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		return nil, err
	}
	gs := grpc.NewServer()
	server := &Server{
		cfg: &cfg,
		gs:  gs,
		svc: s,
	}
	api.RegisterUserServer(gs, server)
	reflection.Register(gs)
	return server, nil
}

func (s *Server) Start() {
	addr := s.cfg.Addr
	gs := s.gs
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal().Msgf("tcp listen: %v", err)
	}
	go func() {
		if err := gs.Serve(lis); err != nil {
			log.Fatal().Msgf("failed to serve: %v", err)
		}
	}()
	return
}
