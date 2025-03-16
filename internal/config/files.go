package config

import (
	"path/filepath"
)

var (
	CAFile        = configFile("ca.pem")
	ServerCerFile = configFile("server.pem")
	ServerKeyFile = configFile("server-key.pem")
	ClientCerFile = configFile("client.pem")
	ClientKeyFile = configFile("client-key.pem")
)

func configFile(filename string) string {
	rootDir, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		panic("failed to resolve project root: " + err.Error())
	}
	return filepath.Join(rootDir, ".proglog", filename)
}
