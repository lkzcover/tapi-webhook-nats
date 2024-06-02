package config

import "os"

type Config struct {
	HttpServer struct {
		Port string
		Cert string
		Key  string
	}

	NatsConn string
	NastSubj string
}

func LoadEnv() (Config, error) {
	cfg := Config{}

	cfg.HttpServer.Port = os.Getenv("HTTP_SERVER_PORT")
	cfg.HttpServer.Cert = os.Getenv("HTTP_SERVER_CERT")
	cfg.HttpServer.Key = os.Getenv("HTTP_SERVER_KEY")

	cfg.NatsConn = os.Getenv("NATS_CONN")
	cfg.NastSubj = os.Getenv("NAST_SUBJ")

	// TODO add validation @lkzcover

	return cfg, nil
}
