package config

import "os"

func Test() string {
	return os.Getenv("test")
}
