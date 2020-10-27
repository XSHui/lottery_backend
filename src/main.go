package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"

	"lottery_backend/src/config"
	xredis "lottery_backend/src/redis"
	api "lottery_backend/src/server"
	"lottery_backend/src/xlog"
	"lottery_backend/src/xorm"
)

var cfg *config.Config

// initConfig: init lottery server info and db info
func initConfig() {
	cfg = config.NewConfig()
	// server info
	flag.StringVar(&cfg.ServerInfo.Ip, "ip", "0.0.0.0", "ip address to listen")
	flag.UintVar(&cfg.ServerInfo.ListenPort, "listen_port", 8888, "listen port")
	flag.UintVar(&cfg.ServerInfo.ManagePort, "manage_port", 9999, "manage port")

	// db info
	flag.StringVar(&cfg.DbInfo.DbIp, "db_ip", "0.0.0.0", "db ip")
	flag.UintVar(&cfg.DbInfo.DbPort, "db_port", 6033, "db port")
	flag.StringVar(&cfg.DbInfo.DbUser, "db_user", "root", "db user")
	flag.StringVar(&cfg.DbInfo.DbPassword, "db_password", "root-root", "db password")
	flag.StringVar(&cfg.DbInfo.DbName, "db_name", "lottery", "db name")

	// redis info
	flag.StringVar(&cfg.RedisInfo.RedisIp, "redis_ip", "0.0.0.0", "redis ip")
	flag.UintVar(&cfg.RedisInfo.RedisPort, "redis_port", 6379, "redis port")
	flag.StringVar(&cfg.RedisInfo.RedisUser, "redis_user", "root", "redis user")
	flag.StringVar(&cfg.RedisInfo.RedisPassword, "redis_password", "root-root", "redis password")

	// parse
	flag.Parse()
}

// configCheck: simple check for lottery configuration
func configCheck() {
	// TODO
}

func main() {
	initConfig()
	// TODO: configuration check

	// TODO: init log
	xlog.Init("/data/lottery.log", "debug", 168, 24)
	xlog.DebugSimple("WelCome Lottery!", xlog.Fields{})

	// init redis
	rdb := xredis.GetRedisInstance()
	rdb.RedisPool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp",
				fmt.Sprintf("%s:%s", cfg.RedisInfo.RedisIp, cfg.RedisInfo.RedisPort),
				//redis.DialUsername(cfg.RedisInfo.RedisUser),
				redis.DialPassword(cfg.RedisInfo.RedisPassword),
			)
			if err != nil {
				xlog.ErrorSimple("Init Redis Failed!", xlog.Fields{
					"err": err,
				})
				panic(conn)
			}
			return
		},
		MaxIdle:     50,
		MaxActive:   2000,
		IdleTimeout: 180 * time.Second,
	}

	// init db
	err := xorm.GetInstance().Init(
		cfg.DbInfo.DbIp,
		int(cfg.DbInfo.DbPort),
		cfg.DbInfo.DbUser,
		cfg.DbInfo.DbPassword,
		cfg.DbInfo.DbName,
		"info")
	if err != nil {
		xlog.ErrorSimple("Init DB Failed!", xlog.Fields{
			"err": err,
		})
		os.Exit(1)
	}
	xlog.DebugSimple("Init DB Success!", xlog.Fields{})

	// init server
	api.StartHttpServer(cfg.ServerInfo.Ip, int(cfg.ServerInfo.ListenPort))

	os.Exit(0)
}
