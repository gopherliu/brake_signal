package config

var C struct {
	Debug bool `json:"Debug"`

	RedisAddr string `json:"redis_addr"`
	RedisAuth string `json:"redis_auth"`

	Bind string `json:"bind"`
}
