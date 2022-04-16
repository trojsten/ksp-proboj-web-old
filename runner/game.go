package runner

import (
	"encoding/json"
	"os"
)

type Game struct {
	Gamefolder string   `json:"gamefolder"`
	Players    []string `json:"players"`
	Args       string   `json:"args"`
}

func (g Game) Save(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode([]Game{g})
}
