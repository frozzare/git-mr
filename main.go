package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/frozzare/go-gitlab-client"
	"github.com/urfave/cli"
)

var (
	description string
	project     string
	source      string
	target      string
	title       string
)

func main() {
	app := cli.NewApp()
	app.Name = "Git Merge Request"
	app.Usage = "Simple command line to do merge request to GitLab"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "project, p",
			Value:       "",
			Usage:       "project (namespace/project_name)",
			Destination: &project,
		},
		cli.StringFlag{
			Name:        "description, d",
			Value:       "",
			Usage:       "merge request description",
			Destination: &description,
		},
		cli.StringFlag{
			Name:        "source, s",
			Value:       "",
			Usage:       "source branch",
			Destination: &source,
		},
		cli.StringFlag{
			Name:        "target, t",
			Value:       "",
			Usage:       "target branch",
			Destination: &target,
		},
		cli.StringFlag{
			Name:        "title, m",
			Value:       "",
			Usage:       "merge request title",
			Destination: &title,
		},
	}

	app.Action = run

	app.Run(os.Args)
}

func client(addr string, token string) (*gogitlab.Gitlab, error) {
	u, err := url.Parse(addr)

	if err != nil {
		return nil, err
	}

	u.Path = ""

	return gogitlab.NewGitlab(u.String(), "/api/v3", token), nil
}

func run(c *cli.Context) error {
	if len(source) == 0 {
		fmt.Println("Source branch is not specified")
		return nil
	}

	if len(target) == 0 {
		fmt.Println("Target branch is not specified")
		return nil
	}

	if len(title) == 0 {
		title = source
	}

	if len(project) == 0 {
		fmt.Println("Project is not specified")
		return nil
	}

	token := os.Getenv("GITLAB_TOKEN")

	if len(token) == 0 {
		fmt.Println("GitLab token is empty")
		return nil
	}

	url := os.Getenv("GITLAB_URL")

	if len(url) == 0 {
		fmt.Println("GitLab url is empty")
		return nil
	}

	lab, err := client(url, token)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	p, err := lab.Project(strings.Replace(project, "/", "%2F", -1))

	if err != nil {
		log.Fatal(err)
		return nil
	}

	mr, err := lab.AddMergeRequest(&gogitlab.AddMergeRequestRequest{
		Description:     description,
		SourceBranch:    source,
		TargetBranch:    target,
		TargetProjectId: p.Id,
		Title:           title,
	})

	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println(fmt.Sprintf("%s/%s/merge_requests/%d", url, project, mr.Iid))
	return nil
}
