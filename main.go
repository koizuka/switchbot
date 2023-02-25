package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nasa9084/go-switchbot/v2"
	"os"
)

const (
	STATUS_INVALID_ARGUMENT = 2
	STATUS_CONFIG_ERROR     = 2
	STATUS_API_ERROR        = 1
)

type Config struct {
	OpenToken string
	SecretKey string
}

func help() {
	fmt.Println("Usage: switchbot <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  list")
	fmt.Println("  status <device_id...>")
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(STATUS_INVALID_ARGUMENT)
	}

	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Printf("config.toml error %v\n", err)
		os.Exit(STATUS_CONFIG_ERROR)
	}

	c := switchbot.New(config.OpenToken, config.SecretKey)
	cmd := NewCommands(c)

	switch os.Args[1] {
	case "list":
		// get physical devices and show
		pdev, err := cmd.ListDevices(context.Background())
		if err != nil {
			fmt.Printf("  error %v\n", err)
			os.Exit(STATUS_API_ERROR)
		} else {
			var results []*DeviceItem
			for _, d := range pdev {
				i := ParseDeviceItem(&d)
				results = append(results, i)
			}
			j, err := json.Marshal(results)
			if err != nil {
				fmt.Printf("  error %v\n", err)
				os.Exit(STATUS_API_ERROR)
			}
			fmt.Printf("%v\n", string(j))
		}
	case "status":
		if len(os.Args) < 3 {
			help()
			os.Exit(STATUS_INVALID_ARGUMENT)
		}
		var results []*DeviceStatus
		for i := 2; i < len(os.Args); i++ {
			id := os.Args[i]
			status, err := cmd.Status(context.Background(), id)
			if err != nil {
				fmt.Printf("%v  error %v\n", id, err)
				os.Exit(STATUS_API_ERROR)
			}
			results = append(results, ParseDeviceStatus(status))
		}
		j, err := json.Marshal(results)
		if err != nil {
			fmt.Printf("  error %v\n", err)
		} else {
			fmt.Printf("%v\n", string(j))
		}
	}
}
