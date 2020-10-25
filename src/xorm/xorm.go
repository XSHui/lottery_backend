package xorm

import (
	"fmt"
	"sync"

	model "lottery_backend/src/xorm/model"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/go-xorm/core"
	goxorm "github.com/go-xorm/xorm"
	"xorm.io/core"
)

type XormStoreManager struct {
	sync.Mutex
	dbName     string // db名字
	dbUser     string //登录用户名
	dbIp       string // db ip
	dbPort     int    // db port
	dbPw       string // db password
	inited     bool   // 是否已经初始化
	connnected bool   // 是否连接
	logLevel   string // 日志等级
	// xorm
	engine *goxorm.Engine
}

var xormStoreManager *XormStoreManager = nil
var xormOnce sync.Once

func GetInstance() *XormStoreManager {
	xormOnce.Do(func() {
		xormStoreManager = &XormStoreManager{}
		xormStoreManager.inited = false
		xormStoreManager.connnected = false
	})
	return xormStoreManager
}

func (s *XormStoreManager) Init(ip string, port int, user string, password string, dbname string, logLevel string) error {

	s.dbUser = user
	s.dbPw = password
	s.dbName = dbname
	s.dbIp = ip
	s.dbPort = port
	s.inited = true
	s.logLevel = logLevel

	return s.initTables()
}

func (s *XormStoreManager) connect() error {

	if !s.connnected {
		// 数据库名称:数据库连接密码@(数据库地址:3306)/数据库实例名称?charset=utf8
		params := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", s.dbUser, s.dbPw, s.dbIp, s.dbPort, s.dbName)
		engine, err := goxorm.NewEngine("mysql", params)
		if err != nil {
			return err
		}

		logLevel := getXormLogLevel(s.logLevel)

		engine.Logger().SetLevel(logLevel)
		s.engine = engine

		err = s.engine.Ping()
		if err != nil {
			return err
		}

		if logLevel == core.LOG_DEBUG {
			s.engine.ShowSQL(true)
		}

		s.Lock()
		s.connnected = true
		s.Unlock()
	}
	return nil
}

// InitTables sync tables
func (s *XormStoreManager) initTables() error {
	var _tables []interface{}
	_tables = append(_tables, new(model.User))
	_tables = append(_tables, new(model.Article))
	_tables = append(_tables, new(model.Permission))
	_tables = append(_tables, new(model.Prize))
	_tables = append(_tables, new(model.Record))
	if err := s.connect(); err != nil {
		return err
	}
	return s.engine.Sync2(_tables...)
}

func GetDB() *goxorm.Engine {
	instance := GetInstance()
	if !instance.inited {
		return nil
	}

	if !instance.connnected {
		if instance.connect() != nil {
			return nil
		}
	}
	return instance.engine
}

// getXormLogLevel
// convert level string to core log level
func getXormLogLevel(level string) core.LogLevel {
	xormLog := core.LOG_OFF
	switch level {
	case "debug":
		xormLog = core.LOG_DEBUG
	case "info":
		xormLog = core.LOG_INFO
	case "warn":
		xormLog = core.LOG_WARNING
	case "error":
		xormLog = core.LOG_ERR
	}
	return xormLog
}
