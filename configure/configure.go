package configure

import (
	"flag"
	"fmt"
	"github.com/magiconair/properties"
	"github.com/zceuiv/gorest/godbc"
	"log"
	"net/http"
	"os"
	"time"
)

/*
 * 该struct要实现main包中的Config接口
 **/
type configProperties struct {
	HttpServerHost   string
	HttpServerPort   int
	HttpReadTimeout  time.Duration
	HttpWriteTimeout time.Duration
	HttpServerCross  bool
	server           *http.Server

	DbSqlName        string
	DbServerHost     string
	DbServerPort     int
	DbServerDatabase string
	DbServerUser     string
	DbServerPassword string
	dbOpt            *godbc.Options
}

var (
	__SingletonConfigProperties *configProperties
)

func init() {
	if __SingletonConfigProperties == nil {
		__SingletonConfigProperties = newConfig()
	}
	if __SingletonConfigProperties == nil {
		os.Exit(1)
	}
}

/*
 * configProperties的构造函数
 **/
func newConfig() (config *configProperties) {
	propertiesFile := flag.String("config", "gorest.conf", "the configuration file")
	var p *properties.Properties
	var err error
	if p, err = properties.LoadFile(*propertiesFile, properties.UTF8); err != nil {
		log.Fatalf("[error] Unable to read properties:%v\n", err)
		return nil
	}

	config = new(configProperties)
	config.HttpServerHost = p.GetString("http.server.host", "0.0.0.0")
	config.HttpServerPort = p.GetInt("http.server.port", 3000)
	config.HttpReadTimeout = p.GetDuration("http.server.readtimeout", 10)
	config.HttpWriteTimeout = p.GetDuration("http.server.writetimeout", 10)
	config.HttpServerCross = p.GetBool("http.server.cross", false)

	config.DbSqlName = p.GetString("database.sql.name", "mysql")
	config.DbServerHost = p.GetString("database.pg.host", "0.0.0.0")
	config.DbServerPort = p.GetInt("database.pg.port", 5432)
	config.DbServerDatabase = p.MustGetString("database.pg.database")
	config.DbServerUser = p.MustGetString("database.pg.user")
	config.DbServerPassword = p.MustGetString("database.pg.password")

	config.dbOpt, config.server = config.initialize()
	return config
}

/*
 * 返回Config的单个实例
 **/
func SingletonConfig() *configProperties {
	return __SingletonConfigProperties
}

func (config *configProperties) initialize() (dbOpt *godbc.Options, server *http.Server) {
	server = new(http.Server)
	server.Addr = fmt.Sprintf("%s:%d", config.HttpServerHost, config.HttpServerPort)
	server.ReadTimeout = config.HttpReadTimeout * time.Second
	server.WriteTimeout = config.HttpWriteTimeout * time.Second

	dbOpt = new(godbc.Options)
	dbOpt.SqlName = config.DbSqlName
	dbOpt.Addr = fmt.Sprintf("%s:%d", config.DbServerHost, config.DbServerPort)
	dbOpt.Database = config.DbServerDatabase
	dbOpt.User = config.DbServerUser
	dbOpt.Password = config.DbServerPassword
	return dbOpt, server
}

/*
 * 实现Config接口
 **/
func (config *configProperties) Server() *http.Server {
	return config.server
}

func (config *configProperties) DbOptions() *godbc.Options {
	return config.dbOpt
}
