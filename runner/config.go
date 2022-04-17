package runner

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server  string            `json:"server"`
	Players map[string]string `json:"players"`
	Timeout float32           `json:"timeout"`
}

func (c Config) Save(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(c)
}
