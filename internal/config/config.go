// Package config provides a struct to store the applications config
package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/loggingx"
	"go.infratographer.com/x/otelx"
)

// AppConfig stores all the config values for our application
var AppConfig struct {
	CRDB    crdbx.Config
	Logging loggingx.Config
	Server  echox.Config
	Tracing otelx.Config
	URN     struct {
		Namespace string `mapstructure:"namespace"`
		Types     struct {
			Instance         string `mapstructure:"instance"`
			InstanceProvider string `mapstructure:"instance_provider"`
		} `mapstructure:"types"`
	}
}

// MustViperFlags returns the viper config for application specific config
func MustViperFlags(v *viper.Viper, _ *pflag.FlagSet) {
	v.SetDefault("urn.namespace", "com.infratographer")
	v.SetDefault("urn.types.instance", "instance")
	v.SetDefault("urn.types.instance_provider", "instance-provider")
}
