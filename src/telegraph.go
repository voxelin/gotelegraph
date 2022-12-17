package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/kallydev/telegraph-go"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "telegraph",
		Usage:                "Create telegra.ph articles right in terminal.",
		EnableBashCompletion: true,
		Compiled:             time.Now(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "title",
				Value:   "Untitled",
				Usage:   "Set title for an article",
				Aliases: []string{"t"},
			},
			&cli.StringFlag{
				Name:    "author",
				Value:   "Anonymous",
				Usage:   "Set author for an article",
				Aliases: []string{"a"},
			},
			&cli.StringFlag{
				Name:     "content",
				Usage:    "Set content for an article",
				Aliases:  []string{"c"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			title := c.String("title")
			author := c.String("author")
			content := c.String("content")
			client, err := telegraph.NewClient("", nil)
			if err != nil {
				log.Panicln(err)
			}
			account, err := client.CreateAccount(author, &telegraph.CreateAccountOption{
				AuthorName: author,
			})
			if err != nil {
				log.Panicln(err)
			}
			client.AccessToken = account.AccessToken
			page, err := client.CreatePage(title, []telegraph.Node{
				telegraph.NodeElement{
					Tag: "p",
					Children: []telegraph.Node{
						content,
					},
				},
			}, &telegraph.CreatePageOption{
				ReturnContent: true,
			})
			if err != nil {
				log.Panicln(err)
			}
			fmt.Println(page.URL)

			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
