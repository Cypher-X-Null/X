package cypher

import (
	"encoding/json"
	"errors"
	"os"
)

type Config map[string]interface{}

func LoadConfig(file string) (Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}