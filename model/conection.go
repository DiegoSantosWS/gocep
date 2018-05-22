package model

//Config recebe dados de conexão com banco de dados
type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	DB   string `yaml:"bd"`
}
