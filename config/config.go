package config

type Config struct {
	ServerAddress string
	RateLimit     int
}

func Load() Config {
	return Config{
		ServerAddress: ":8080",
		RateLimit:     5,
	}
}
