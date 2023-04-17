package main

import (
	"log"
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
		&cli.StringFlag{
			Name:    "repo.name",
			EnvVars: []string{"DRONE_REPO_NAME"},
		},
		&cli.StringFlag{
			Name:    "commit.sha",
			EnvVars: []string{"DRONE_COMMIT_SHA"},
		},
		&cli.StringFlag{
			Name:    "commit.ref",
			EnvVars: []string{"DRONE_COMMIT_REF"},
		},
		&cli.StringFlag{
			Name:    "commit.branch",
			EnvVars: []string{"DRONE_COMMIT_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "commit.author",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR"},
		},
		&cli.StringFlag{
			Name:    "commit.author.email",
			Usage:   "git author email",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR_EMAIL"},
		},
		&cli.StringFlag{
			Name:    "commit.author.avatar",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR_AVATAR"},
		},
		&cli.StringFlag{
			Name:    "commit.author.name",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR_NAME"},
		},
		&cli.StringFlag{
			Name:    "commit.pull",
			EnvVars: []string{"DRONE_PULL_REQUEST"},
		},
		&cli.StringFlag{
			Name:    "commit.message",
			EnvVars: []string{"DRONE_COMMIT_MESSAGE"},
		},
		&cli.StringFlag{
			Name:    "build.event",
			EnvVars: []string{"DRONE_BUILD_EVENT"},
		},
		&cli.IntFlag{
			Name:    "build.number",
			EnvVars: []string{"DRONE_BUILD_NUMBER"},
		},
		&cli.IntFlag{
			Name:    "build.parent",
			EnvVars: []string{"DRONE_BUILD_PARENT"},
		},
		&cli.StringFlag{
			Name:    "build.status",
			EnvVars: []string{"DRONE_BUILD_STATUS"},
		},
		&cli.StringFlag{
			Name:    "build.link",
			EnvVars: []string{"DRONE_BUILD_LINK"},
		},
		&cli.Int64Flag{
			Name:    "build.started",
			EnvVars: []string{"DRONE_BUILD_STARTED"},
		},
		&cli.Int64Flag{
			Name:    "build.created",
			EnvVars: []string{"DRONE_BUILD_CREATED"},
		},
		&cli.StringFlag{
			Name:    "build.tag",
			EnvVars: []string{"DRONE_TAG"},
		},
		&cli.StringFlag{
			Name:    "build.deployTo",
			EnvVars: []string{"DRONE_DEPLOY_TO"},
		},
		&cli.Int64Flag{
			Name:    "job.started",
			EnvVars: []string{"DRONE_JOB_STARTED"},
		},
	}

	app.Action = run
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Tag:    c.String("build.tag"),
			Number: c.Int("build.number"),
			Parent: c.Int("build.parent"),
			Event:  c.String("build.event"),
			Status: c.String("build.status"),
			Commit: c.String("commit.sha"),
			Ref:    c.String("commit.ref"),
			Branch: c.String("commit.branch"),
			Author: Author{
				Username: c.String("commit.author"),
				Name:     c.String("commit.author.name"),
				Email:    c.String("commit.author.email"),
				Avatar:   c.String("commit.author.avatar"),
			},
			Pull:     c.String("commit.pull"),
			Message:  c.String("commit.message"),
			DeployTo: c.String("build.deployTo"),
			Link:     c.String("build.link"),
			Started:  c.Int64("build.started"),
			Created:  c.Int64("build.created"),
		},
		Job: Job{
			Started: c.Int64("job.started"),
		},
		Config: Config{
			Token:   c.String("token"),
			User:    c.String("user"),
			Message: c.String("message"),
			Title:   c.String("title"),
			Device:  c.String("device"),
		},
	}
	return plugin.Exec()
}
