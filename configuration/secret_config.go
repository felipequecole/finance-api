package configuration

type Secrets struct {
	Database struct {
		Username string `yaml:"user" validate:"required"`
		Password string `yaml:"password" validate:"required"`
	} `yaml:"database"`
}
