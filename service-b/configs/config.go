package configs

import "github.com/spf13/viper"

type Conf struct {
	WebServerPort     			string `mapstructure:"WEB_SERVER_PORT"`
	WeatherApiKey   			string `mapstructure:"WEATHER_API_KEY"`
	OtelExporterOtlpEndpoint   	string `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	OtelServiceName   			string `mapstructure:"OTEL_SERVICE_NAME"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}