package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

type App struct {
	// JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	DB     DB
	Server Server
	Port   int `mapstructure:"port" yaml:"port"`
	// Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	// Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	// Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	// Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	// System  System  `mapstructure:"system" json:"system" yaml:"system"`
	// Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// // auto
	// AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	// // gorm
	// // Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// // oss
	// Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	// Qiniu      Qiniu      `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	// AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`
	// TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencentCOS" yaml:"tencent-cos"`
	// Excel      Excel      `mapstructure:"excel" json:"excel" yaml:"excel"`
	// Timer      Timer      `mapstructure:"timer" json:"timer" yaml:"timer"`
}

func (app App) Validate() error {
	// return validation.ValidateStruct(&app,
	// 	validation.Field(&app.DB.Driver, validation.Required),
	// 	// validation.Field(&app.JWTSigningKey, validation.Required),
	// 	// validation.Field(&app.JWTVerificationKey, validation.Required),
	// )
	return nil

}

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (app *App, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	//
	// viper.SetEnvPrefix(viper.GetString("ENV"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// Read config file.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(fmt.Errorf("fatal error %w", err))
		} else {
			// Config file was found but another error was produced
			fmt.Printf("fatal error  %w", err)
		}
	}

	color.New(color.FgCyan, color.Bold).Printf("Using config file: %v\n", viper.ConfigFileUsed())

	if err := viper.Unmarshal(&app); err != nil {
		fmt.Println(err)
	}
	if err := app.Validate(); err != nil {
		color.New(color.FgRed, color.Bold).Printf("Missing required configs: %v\n", err)
		os.Exit(0)
	}
	return app, nil
}
