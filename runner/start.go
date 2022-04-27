package runner

import (
	"context"
	"ksp.sk/proboj/web/config"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

type GameResult struct {
}

func RunGame(conf Config, game Game) error {
	temp, err := os.MkdirTemp("", "proboj-")
	if err != nil {
		return err
	}

	defer os.RemoveAll(temp)

	err = game.Save(path.Join(temp, "games.json"))
	if err != nil {
		return err
	}

	err = conf.Save(path.Join(temp, "config.json"))
	if err != nil {
		return err
	}

	log.Printf("Starting game in %v.\n", temp)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, config.Configuration.RunnerCommand, path.Join(temp, "config.json"), path.Join(temp, "games.json"))

	if config.Configuration.RunnerDebug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	err = cmd.Run()
	log.Printf("Game %s over.\n", temp)
	return err
}
