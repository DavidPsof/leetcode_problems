package config

import "github.com/spf13/viper"

type Conf struct {
	HttpServerConfig `mapstructure:",squash"`
	PostgresConfig   `mapstructure:",squash"`
	LogSettings      `mapstructure:",squash"`
}

type HttpServerConfig struct {
	Port         string `mapstructure:"HTTP_PORT" envDefault:"8080"`
	ReadTimeout  uint8  `mapstructure:"HTTP_READ_TIMEOUT" envDefault:"60"`
	WriteTimeout uint8  `mapstructure:"HTTP_WRITE_TIMEOUT" envDefault:"60"`
	IdleTimeout  uint8  `mapstructure:"HTTP_PORT_IDLE_TIMEOUT" envDefault:"60"`
	JwtSignKey   string `mapstructure:"JWT_SIGN_KEY" envDefault:"dda28eba-3670-11ed-a261-0242ac120002"`
}

type PostgresConfig struct {
	Host         string `mapstructure:"POSTGRES_HOST,required"`
	Port         string `mapstructure:"POSTGRES_PORT,required"`
	User         string `mapstructure:"POSTGRES_USER,required"`
	Pass         string `mapstructure:"POSTGRES_PASSWORD,required"`
	DBName       string `mapstructure:"POSTGRES_DB_NAME,required"`
	ModeSSL      string `mapstructure:"POSTGRES_SSL_MODE,required"`
	MaxOpenConns uint   `mapstructure:"POSTGRES_MAX_OPEN_CONNS" envDefault:"10"`
	MaxIdleConns uint   `mapstructure:"POSTGRES_MAX_IDLE_CONNS" envDefault:"5"`
}

type LogSettings struct {
	LogLevel          string `mapstructure:"LOG_LEVEL" envDefault:"INFO"`
	LogInConsole      bool   `mapstructure:"LOG_CONSOLE" envDefault:"true"`
	LogInFile         bool   `mapstructure:"LOG_FILE" envDefault:"false"`
	LogFileName       string `mapstructure:"LOG_FILE_NAME" envDefault:"log"`
	LogFileMaxSize    int    `mapstructure:"LOG_FILE_MAX_SIZE" envDefault:"2048"`
	LogFileMaxBackups int    `mapstructure:"LOG_FILE_MAX_BACKUPS" envDefault:"10"`
	LogFileMaxMaxAge  int    `mapstructure:"LOG_FILE_MAX_AGE" envDefault:"10"` // in days
	LogReportCaller   bool   `mapstructure:"LOG_REPORT_CALLER" envDefault:"true"`
}

func InitConfig(name string) *Conf {
	var conf Conf

	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName(name)
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
