package dao

import (
	"os"
	"path/filepath"

	e "github.com/fuwensun/goms/eApi/internal/pkg/err"
	"github.com/fuwensun/goms/pkg/conf"

	_ "github.com/go-sql-driver/mysql" // for init()
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
)

// cccfg config of cache.
type cccfg struct {
	Addr string `yaml:"addr"`
	Pass string `yaml:"pass"`
}

// getCCConfig get cache config from file and env.
func getCCConfig(cfgpath string) (*cccfg, error) {
	var err error
	cfg := &cccfg{}
	// file
	path := filepath.Join(cfgpath, "redis.yaml")
	if err = conf.GetConf(path, &cfg); err != nil {
		log.Warn().Msgf("get cc config file error: %v", err)
	} else if cfg.Addr == "" {
		log.Warn().Msgf("get cc config file succ, but ADDR IS EMPTY")
	} else {
		log.Info().Msgf("get cc config file succ, Addr: %v", cfg.Addr)
		return cfg, nil
	}
	// env
	if addr, exist := os.LookupEnv("REDIS_SVC_ADDR"); exist == false {
		log.Warn().Msgf("get cc config env error: %v", e.ErrNotFoundData)
	} else if addr == "" {
		log.Warn().Msgf("get cc config env succ, but ADDR IS EMPTY")
	} else {
		log.Info().Msgf("get cc config env succ, Addr: %v", cfg.Addr)
		cfg.Addr = addr
		return cfg, nil
	}
	return nil, e.ErrNotFoundData
}

// newCC new cache and return.
func newCC(cfgpath string) (redis.Conn, func(), error) {
	if cf, err := getCCConfig(cfgpath); err != nil {
		log.Error().Msgf("get cc config error: %v", err)
		return nil, nil, err
	} else if cc, err := redis.Dial("tcp", cf.Addr, redis.DialPassword(cf.Pass)); err != nil {
		log.Error().Msgf("dial cc error: %v", err)
		return nil, nil, err
	} else if _, err = cc.Do("PING"); err != nil {
		log.Error().Msgf("ping cc error: %v", err)
		return nil, nil, err
	} else {
		return cc, func() {
			cc.Close()
		}, nil
	}
}
