package config_test

import (
	"bufio"
	"os"
	"rest-mongo/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	file, err := os.Open("../env.yml")
	if err != nil {
		panic(err)
	}

	config.Load(bufio.NewReader(file), "yaml")
}
