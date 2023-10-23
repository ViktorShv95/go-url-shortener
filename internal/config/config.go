package config

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost"`
	Port string `yaml:"port" env-default:"8082"`
	Timeout string `yaml:"timeout" env-default:"4s"`
	IddleTimeout string `yaml:"iddle_timeout" env-default:"45s"`
}
