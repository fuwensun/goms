package dao

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/gomsx/goms/pkg/conf"
	e "github.com/gomsx/goms/pkg/err"

	_ "github.com/go-sql-driver/mysql" // for init()
)

// dbcfg config of db.
type dbcfg struct {
	DSN string `yaml:"dsn"`
}

// getDBConfig get db config from file and env.
func getDBConfig(cfgpath string) (*dbcfg, error) {
	var err error
	cfg := &dbcfg{}
	// file
	path := filepath.Join(cfgpath, "mysql.yaml")
	if err = conf.GetConf(path, &cfg); err != nil {
		log.Printf("get db config file error: %v", err)
	} else if cfg.DSN == "" {
		log.Printf("get db config file succeeded, but DSN IS EMPTY")
	} else {
		log.Printf("get db config file succeeded, DSN: %v", "***")
		return cfg, nil
	}
	// env
	if dsn, exist := os.LookupEnv("MYSQL_SVC_DSN"); exist == false {
		log.Printf("get db config env error: %v", e.ErrNotFoundData)
	} else if dsn == "" {
		log.Printf("get db config env succeeded, but DSN IS EMPTY")
	} else {
		log.Printf("get db config env succeeded, DSN: %v", "***")
		cfg.DSN = dsn
		return cfg, nil
	}
	return nil, e.ErrNotFoundData
}

// newDB new database and return.
func newDB(cfgpath string) (*sql.DB, func(), error) {
	if df, err := getDBConfig(cfgpath); err != nil {
		log.Printf("get db config error: %v", err)
		return nil, nil, err
	} else if db, err := sql.Open("mysql", df.DSN); err != nil {
		log.Printf("open db error: %v", err)
		return nil, nil, err
	} else if err := db.Ping(); err != nil {
		log.Printf("ping db error: %v", err)
		return nil, nil, err
	} else {
		return db, func() {
			db.Close()
		}, nil
	}
}
