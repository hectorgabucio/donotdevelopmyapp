package config

import "os"

type ConfigProvider interface {
	Get(key string) string
}

type OsEnv struct{}

func (o OsEnv) Get(key string) string {
	return os.Getenv(key)
}
