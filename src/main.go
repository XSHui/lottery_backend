package main

import (
	"flag"
	"fmt"
	"lottery_backend/src/config"
	api "lottery_backend/src/server"
	"lottery_backend/src/xlog"
	"lottery_backend/src/xorm"
	"os"
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

	// init db
	err := xorm.GetInstance().Init(
		cfg.DbInfo.DbIp,
		int(cfg.DbInfo.DbPort),
		cfg.DbInfo.DbUser,
		cfg.DbInfo.DbPassword,
		cfg.DbInfo.DbName,
		"info")
	if err != nil {
		fmt.Println("Init Db Error: %v", err)
		os.Exit(1)
	}

	// init server
	api.StartHttpServer(cfg.ServerInfo.Ip, int(cfg.ServerInfo.ListenPort))

	os.Exit(0)
}
