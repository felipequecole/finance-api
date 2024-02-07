package configuration

type AppConfig struct {
	Database struct {
		Host   string `yaml:"host" validate:"required"`
		Schema string `yaml:"database" validate:"required"`
		Port   uint   `yaml:"port" validate:"required"`
	} `yaml:"database"`
	App struct {
		Port uint   `yaml:"port" validate:"required"`
		Host string `yaml:"host" validate:"required"`
	} `yaml:"app"`
}
