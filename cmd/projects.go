package cmd

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"git.psu.edu/swe-golang/gitlabctl/client"

	gitlab "github.com/xanzy/go-gitlab"
)

// projectCmd represents the command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "list all gitlab projects",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("Listing your GitLab projects")

		client, err := client.GetClient()
		if err != nil {
			color.New(color.FgRed).Printf("Error: %s", err.Error())
			return
		}

		statistics := true
		opt := &gitlab.ListProjectsOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: 100,
				Page:    1,
			},
			Statistics: &statistics,
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
				cols := []string{
					strconv.Itoa(p.ID),
					p.Name,
					p.Path,
					p.DefaultBranch,
					string(p.Visibility),
					strconv.FormatBool(p.Archived),
					strconv.Itoa(p.CreatorID),
					strings.Join(p.TagList, ","),
					p.Namespace.Path,
					p.Namespace.Name,
				}

				if p.Statistics != nil {
					cols = append(cols, strconv.Itoa(p.Statistics.CommitCount))
					cols = append(cols, strconv.FormatInt(p.Statistics.RepositorySize, 10))
				} else {
					cols = append(cols, "")
					cols = append(cols, "")
				}

				err = writer.Write(cols)
				if err != nil {
					color.New(color.FgRed).Printf("Error: %s", err.Error())
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
	rootCmd.AddCommand(projectsCmd)
}
