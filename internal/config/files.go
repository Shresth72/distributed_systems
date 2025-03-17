package config

import (
	"path/filepath"
)

var (
	CAFile              = configFile("ca.pem")
	ServerCerFile       = configFile("server.pem")
	ServerKeyFile       = configFile("server-key.pem")
	RootClientCerFile   = configFile("root-client.pem")
	RootClientKeyFile   = configFile("root-client-key.pem")
	NobodyClientCerFile = configFile("nobody-client.pem")
	NobodyClientKeyFile = configFile("nobody-client-key.pem")
	ACLModelFile        = configFile("model.conf")
	ACLPolicyFile       = configFile("policy.csv")
)

func configFile(filename string) string {
	rootDir, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		panic("failed to resolve project root: " + err.Error())
	}
	return filepath.Join(rootDir, ".certs", filename)
}
