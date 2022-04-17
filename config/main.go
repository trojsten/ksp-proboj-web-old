package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	RunnerCommand string  `json:"runner_command"`
	DataFolder    string  `json:"data_folder"`
	UploadFolder  string  `json:"upload_folder"`
	Database      string  `json:"database"`
	ServerCommand string  `json:"server_command"`
	PlayerTimeout float32 `json:"player_timeout"`
	GamesAhead    int     `json:"games_ahead"`
	MakeCommand   string  `json:"make_command"`
	RunGames      bool    `json:"run_games"`
}

var Configuration Config

func LoadConfig(path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	return decoder.Decode(&Configuration)
}
