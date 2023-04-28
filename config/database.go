package config

type DataBaseConf struct {
	Type     string `yaml:"type"`
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db-name"`
	Charset  string `yaml:"charset"`
}
