package client

import (
	"log"

	"github.com/fatih/color"
	"github.com/spf13/viper"
	gitlab "github.com/xanzy/go-gitlab"
)

func GetClient() (*gitlab.Client, error) {
	url := viper.GetString("gitlab.url")
	token := viper.GetString("gitlab.token")

	color.New(color.FgBlue).Printf("Using url: %s\n", url)

	client, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}

	return client, nil
}
