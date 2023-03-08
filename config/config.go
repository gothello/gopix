package config

import "github.com/spf13/viper"

//LoadConfig is function to load files in one archive YML.
//If success create new viper instancer e erro nil.
//If error occurred return nil and error.

var (
	PATH = "./"
)

func LoadConfig() (*viper.Viper, error) {
	viper := viper.GetViper()

	viper.AddConfigPath(PATH)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return viper, nil
}
