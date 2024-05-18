package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v62/github"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

func main() {
	app := cli.NewApp()
	app.Name = "ghban"
	app.HelpName = "ghban"
	app.Usage = "Block multiple accounts on multiple GitHub organizations."
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "token",
			Usage:   "GitHub personal access token",
			EnvVars: []string{"GITHUB_TOKEN"},
		},
		&cli.StringSliceFlag{
			Name:  "orgs",
			Usage: "GitHub's organization names",
		},
		&cli.StringSliceFlag{
			Name:  "users",
			Usage: "GitHub usernames",
		},
	}

	app.Action = func(cliCtx *cli.Context) error {
		run(context.Background(),
			cliCtx.String("token"),
			cliCtx.StringSlice("users"),
			cliCtx.StringSlice("orgs"),
		)

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, token string, users, orgs []string) {
	client := newGitHubClient(ctx, token)

	for _, user := range users {
		err := banUser(ctx, client, user, orgs)
		if err != nil {
			log.Println(err)
		}
	}
}

func banUser(ctx context.Context, client *github.Client, user string, orgs []string) error {
	fmt.Printf("User profile: https://github.com/%s\n\n", user)

	for _, org := range orgs {
		fmt.Printf("Ban %s from %s: https://github.com/organizations/%[2]s/settings/blocked_users\n", user, org)

		blocked, _, err := client.Organizations.IsBlocked(ctx, org, user)
		if err != nil {
			return err
		}

		if blocked {
			fmt.Println("Already blocked.")

			continue
		}

		_, err = client.Organizations.BlockUser(ctx, org, user)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Ban %s from my account: https://github.com/settings/blocked_users\n", user)

	_, err := client.Users.BlockUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func newGitHubClient(ctx context.Context, token string) *github.Client {
	if token == "" {
		return github.NewClient(nil)
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

	return github.NewClient(oauth2.NewClient(ctx, ts))
}
