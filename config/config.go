package config

import "github.com/spf13/viper"

func LoadConfig() (*viper.Viper, error) {
	viper := viper.GetViper()

	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return viper, nil
}
