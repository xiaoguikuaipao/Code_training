package config

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `mapstructure:"host" json:"host,omitempty"`
	Port int    `mapstructure:"port" json:"port,omitempty"`
	Name string `mapstructure:"name" json:"name,omitempty"`

	Version        string `mapstructure:"version" json:"version,omitempty"`
	MaxConn        int    `mapstructure:"max_conn" json:"max_conn,omitempty"`
	MaxPackageSize int    `mapstructure:"max_package_size" json:"max_package_size,omitempty"`

	MaxQueuePoolSize uint8  `mapstructure:"max_queue_pool_size"`
	MaxQueueWorker   uint32 `mapstructure:"max_queue_worker"`
	MaxQueueSize     uint32 `mapstructure:"max_queue_size"`
}

var DefaultConfig *ServerConfig

var (
	cfg  = pflag.String("server-config", "", "init xyz server with config,only support yaml")
	host = pflag.String("host", "0.0.0.0", "assign listen IP")
	port = pflag.Int("port", 8080, "assign listen Port")
)

func (sc *ServerConfig) reload() error {
	pflag.Parse()
	_ = viper.BindPFlag("host", pflag.Lookup("host"))
	_ = viper.BindPFlag("port", pflag.Lookup("port"))
	if *cfg != "" {
		viper.SetConfigFile(*cfg)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("server")
	}
	if err := viper.ReadInConfig(); err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return err
	}
	err := viper.Unmarshal(sc)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	DefaultConfig = &ServerConfig{
		Host:             *host,
		Port:             *port,
		Name:             "Default XYZ Server",
		Version:          "v0.3",
		MaxConn:          100,
		MaxPackageSize:   4096,
		MaxQueuePoolSize: 10,
		MaxQueueSize:     1024,
		MaxQueueWorker:   uint32(runtime.NumCPU()),
	}
	err := DefaultConfig.reload()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config Init Successfully:%v\n", DefaultConfig)
}
