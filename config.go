package scut

import (
	"github.com/BurntSushi/toml"
	"github.com/zazab/zhash"
)

type Config struct {
	zhash.Hash
	filepath string
}

func NewConfig(filepath string) (*Config, error) {
	tomlMap := map[string]interface{}{}
	_, err := toml.DecodeFile(filepath, &tomlMap)
	if err != nil {
		return nil, err
	}

	hash := zhash.HashFromMap(tomlMap)
	return &Config{Hash: hash, filepath: filepath}, nil
}
