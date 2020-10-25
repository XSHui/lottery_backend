package config

type Config struct {
	// server
	ServerInfo ServerInfo `toml:"server_info" json:"server_info"`
	// db
	DbInfo DbInfo `toml:"db_info" json:"db_info"`
	// redis
	RedisInfo RedisInfo `toml:"redis_info" json:"redis_info"`

}

type ServerInfo struct {
	Ip         string `toml:"ip" json:"ip"`
	ListenPort uint   `toml:"listen_port" json:"listen_port"`
	ManagePort uint   `toml:"manage_port" json:"manage_port"`
}

type DbInfo struct {
	DbIp       string `toml:"db_ip" json:"db_ip"`
	DbPort     uint   `toml:"db_port" json:"db_port"`
	DbUser     string `toml:"db_user" json:"db_user"`
	DbPassword string `toml:"db_password" json:"db_password"`
	DbName     string `toml:"db_name" json:"db_name"`
}

type RedisInfo struct {
	RedisIp       string `toml:"redis_ip" json:"redis_ip"`
	RedisPort     uint   `toml:"redis_port" json:"redis_port"`
	RedisUser     string `toml:"redis_user" json:"redis_user"`
	RedisPassword string `toml:"redis_password" json:"redis_password"`
}



var (
	defaultConf = Config{
		ServerInfo: ServerInfo{
			Ip:         "0.0.0.0",
			ListenPort: 8888,
			ManagePort: 9999,
		},
		DbInfo: DbInfo{
			DbIp:       "0.0.0.0",
			DbPort:     6033,
			DbUser:     "root",
			DbPassword: "root-root",
			DbName:     "lottery",
		},
		RedisInfo: RedisInfo{
			RedisIp:       "0.0.0.0",
			RedisPort:     6033,
			RedisUser:     "root",
			RedisPassword: "root-root",
		},
	}
)

// NewConfig creates a new config instance with default value.
func NewConfig() *Config {
	conf := defaultConf
	return &conf
}

