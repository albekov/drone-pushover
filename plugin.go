package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/drone/drone-template-lib/template"
)

type (
	Repo struct {
		Owner string
		Name  string
	}

	Build struct {
		Tag      string
		Event    string
		Number   int
		Parent   int
		Commit   string
		Ref      string
		Branch   string
		Author   Author
		Pull     string
		Message  string
		DeployTo string
		Status   string
		Link     string
		Started  int64
		Created  int64
	}

	Author struct {
		Username string
		Name     string
		Email    string
		Avatar   string
	}

	Config struct {
		Token   string
		User    string
		Message string
		Title   string
		Device  string
	}

	Job struct {
		Started int64
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

func (p Plugin) Exec() error {
	message, err := template.RenderTrim(p.Config.Message, p)
	if err != nil {
		return err
	}
	title, err := template.RenderTrim(p.Config.Title, p)
	if err != nil {
		return err
	}
	data := map[string]string{
		"token":   p.Config.Token,
		"user":    p.Config.User,
		"message": message,
		"title":   title,
		"device":  p.Config.Device,
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
