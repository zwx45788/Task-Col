package config

import (
	"log"
	"os"
	"project-common/logs"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var C = InitConfig()

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
	GC    *GrpcConfig
	EC    *EtcdConfig
	MC    *MysqlConfig
}
type ServerConfig struct {
	Name string
	Addr string
}
type GrpcConfig struct {
	Name    string
	Addr    string
	Version string
	Weight  int64
}
type EtcdConfig struct {
	Addrs []string
}
type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Db       string
}

func InitConfig() *Config {
	v := viper.New()
	conf := &Config{
		viper: v,
		SC:    &ServerConfig{},
		GC:    &GrpcConfig{},
	}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")

	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	// 初始化 ServerConfig
	conf.SC = conf.InitServerConfig()
	conf.GC = conf.InitGrpcConfig()
	conf.InitZapLog()
	conf.InitRedisOptions()
	conf.ReadEtcdConfig()

	return conf
}
func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}
func (c *Config) InitRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"), // no password set
		DB:       c.viper.GetInt("db"),                // use default DB
	}
}
func (c *Config) InitServerConfig() *ServerConfig {
	return &ServerConfig{
		Name: c.viper.GetString("server.name"),
		Addr: c.viper.GetString("server.addr"),
	}
}
func (c *Config) InitGrpcConfig() *GrpcConfig {
	return &GrpcConfig{
		Name: c.viper.GetString("grpc.name"),
		Addr: c.viper.GetString("grpc.addr"),
	}
}
func (c *Config) ReadEtcdConfig() {
	ec := &EtcdConfig{}
	c.viper.UnmarshalKey("etcd.addrs", ec)
	var addrs []string
	err := c.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalln(err)
	}
	ec.Addrs = addrs
	c.EC = ec
}
func (c *Config) InitMysqlConfig() *MysqlConfig {
	return &MysqlConfig{
		Username: c.viper.GetString("mysql.username"),
		Password: c.viper.GetString("mysql.password"),
		Host:     c.viper.GetString("mysql.host"),
		Port:     c.viper.GetInt("mysql.port"),
		Db:       c.viper.GetString("mysql.db"),
	}

}
