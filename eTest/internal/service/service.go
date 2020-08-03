package service

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/aivuca/goms/eTest/internal/dao"
	m "github.com/aivuca/goms/eTest/internal/model"
	"github.com/aivuca/goms/pkg/conf"

	"github.com/rs/zerolog/log"
)

// Svc service interface.
type Svc interface {
	HandPing(c context.Context, p *m.Ping) (*m.Ping, error)

	CreateUser(c context.Context, user *m.User) error
	ReadUser(c context.Context, uid int64) (*m.User, error)
	UpdateUser(c context.Context, user *m.User) error
	DeleteUser(c context.Context, uid int64) error

	Ping(c context.Context) (err error)
	Close()
}

// Service service.
type service struct {
	cfg *config
	dao dao.Dao
}

// Service config.
type config struct {
	Name    string `yaml:"name,omitempty"`
	Version string `yaml:"version,omitempty"`
}

//

// getConfig get config from config file.
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	filep := filepath.Join(cfgpath, "app.yaml")
	if err := conf.GetConf(filep, cfg); err != nil {
		log.Warn().Msgf("get config file: %v", err)
		err = fmt.Errorf("get config file: %w", err)
		return nil, err
	}
	log.Info().Msgf("config name: %v,version: %v", cfg.Name, cfg.Version)
	return cfg, nil
}

// New new a service and return.
func New(cfgpath string, dao dao.Dao) (Svc, func(), error) {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Error().Msgf("get config error")
		return nil, nil, err
	}
	svc := &service{cfg: cfg, dao: dao}

	log.Info().Msg("service ok")
	return svc, svc.Close, nil
}

// Ping ping the resource.
func (s *service) Ping(c context.Context) (err error) {
	return s.dao.Ping(c)
}

// Close close the resource.
func (s *service) Close() {
}
