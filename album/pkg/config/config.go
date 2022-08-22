package config

import "github.com/spf13/viper"

type Config struct {
	Port       string `mapstructure:"PORT"`
	DBUrl      string `mapstructure:"DB_URL"`
	SongSvcUrl string `mapstructure:"SONG_SVC_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./album/pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
