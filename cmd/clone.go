package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"git.psu.edu/swe-golang/gitlabctl/client"

	gitlab "github.com/xanzy/go-gitlab"
)

var outputSshUrl bool
var outputHttpsUrl bool
var groupFilter string

// cloneCmd represents the command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "helper for cloning projects",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("Listing the clone url for all GitLab projects")

		client, err := client.GetClient()
		if err != nil {
			color.New(color.FgRed).Printf("Error: %s", err.Error())
			return
		}

		opt := &gitlab.ListProjectsOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: 100,
				Page:    1,
			},
		}

		out := os.Stdout
		writer := csv.NewWriter(out)

		header := []string{"id", "name", "path", "default branch", "visibility",
			"archived", "creator", "tags", "namespace path", "namespace name", "commit count", "repository size"}
		err = writer.Write(header)
		if err != nil {
			color.New(color.FgRed).Printf("Error: %s", err.Error())
		}

		for {
			projects, resp, err := client.Projects.ListProjects(opt)
			if err != nil {
				color.New(color.FgRed).Printf("Error: %s", err.Error())
				return
			}

			for _, p := range projects {

				if groupFilter != "" && !strings.HasPrefix(p.Namespace.FullPath, groupFilter) {
					continue
				}

				if outputHttpsUrl {
					fmt.Printf("git clone %s\n", p.HTTPURLToRepo)
				} else if outputSshUrl {
					fmt.Printf("git clone %s\n", p.SSHURLToRepo)
				} else {
					fmt.Printf("git clone %s\n", p.SSHURLToRepo)
				}
			}

			if resp.CurrentPage >= resp.TotalPages {
				break
			}

			opt.Page = resp.NextPage
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	cloneCmd.PersistentFlags().BoolVar(&outputSshUrl, "ssh", true, "output the SSH clone url")
	cloneCmd.PersistentFlags().BoolVar(&outputHttpsUrl, "https", false, "output the HTTPS clone url")

	cloneCmd.PersistentFlags().StringVar(&groupFilter, "group", "", "filter repositories to groups starting with this prefix")

}
