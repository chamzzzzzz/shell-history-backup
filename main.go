package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Config struct {
	DestinationFileName string `json:"destination_file_name"`
}

func main() {
	file := "config.json"
	b, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("read config failed: %v\n", err)
		os.Exit(1)
	}
	var config Config
	err = json.Unmarshal(b, &config)
	if err != nil {
		fmt.Printf("unmarshal config failed: %v\n", err)
		os.Exit(1)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("get user home dir failed: %v\n", err)
		os.Exit(1)
	}

	src := filepath.Join(home, ".zsh_history")
	dst := config.DestinationFileName
	err = exec.Command("cp", "-f", src, dst).Run()
	if err != nil {
		fmt.Printf("backup shell history file failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("backup shell history file success: %s -> %s\n", src, dst)
}
