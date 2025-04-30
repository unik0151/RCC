package config

var GlobalConfig *Conf

type Conf struct {
	MySql *MySqlConf
	Jwt   *JwtConf
	Env   *EnvConf
	Email *EmailConf
}

type MySqlConf struct {
	Dsn          string `toml:"dsn"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxLifeTime  int    `toml:"max_life_time"`
}

type JwtConf struct {
	Secret string `toml:"secret"`
	Expire int    `toml:"expire"`
}

type EnvConf struct {
	Version int `toml:"version"`
	Port    int `toml:"port"`
}

type EmailConf struct {
	From         string `toml:"from"`
	Password     string `toml:"password"`
	Host         string `toml:"host"`
	Subject      string `toml:"subject"`
	Port         string `toml:"port"`
	ValidateTime int    `toml:"validate_time"`
}
