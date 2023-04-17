package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Drone Pushover"
	app.Usage = "Drone plugin for sending Pushover notifications"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "token",
			Required: true,
			Aliases:  []string{"t"},
			EnvVars:  []string{"PLUGIN_TOKEN"},
			Usage:    "Pushover application API token",
		},
		&cli.StringFlag{
			Name:     "user",
			Required: true,
			Aliases:  []string{"u"},
			EnvVars:  []string{"PLUGIN_USER"},
			Usage:    "Pushover user/group key",
		},
		&cli.StringFlag{
			Name:     "message",
			Required: true,
			Aliases:  []string{"m"},
			EnvVars:  []string{"PLUGIN_MESSAGE"},
			Usage:    "Pushover message",
		},
		&cli.StringFlag{
			Name:    "title",
			Aliases: []string{"s"},
			EnvVars: []string{"PLUGIN_TITLE"},
			Usage:   "Pushover message title",
		},
		&cli.StringFlag{
			Name:    "device",
			Aliases: []string{"d"},
			EnvVars: []string{"PLUGIN_DEVICE"},
			Usage:   "The name of one of your Pushover devices",
		},
	}
	app.Action = run
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	data := map[string]string{
		"token":   c.String("token"),
		"user":    c.String("user"),
		"message": c.String("message"),
		"title":   c.String("title"),
		"device":  c.String("device"),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", "https://api.pushover.net/1/messages.json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return fmt.Errorf("pushover API returned status code %d", response.StatusCode)
	}
	return nil
}
