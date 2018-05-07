package config

type Config struct {
	Debug  bool
	Server struct {
		Address string
	}
	Database struct {
		Host string
		Port int
		User string
		Name string
	}
}
