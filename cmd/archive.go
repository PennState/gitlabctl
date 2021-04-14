package cmd

import (
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"

	"git.psu.edu/swe-golang/gitlabctl/client"
	"git.psu.edu/swe-golang/gitlabctl/data"
)

// projectCmd represents the command
var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "archive gitlab projects",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("Archiving your GitLab projects")

		client, err := client.GetClient()
		if err != nil {
			color.New(color.FgRed).Printf("Error: %s", err.Error())
			return
		}

		pid := viper.GetInt("pid")
		if pid != 0 {
			archiveProject(client, pid)
		}

		file := viper.GetString("file")
		if file != "" {
			data, err := data.LoadProjects(file)
			if err != nil {
				color.New(color.FgRed).Printf("Error: %s", err.Error())
				return
			}

			for _, d := range data.Data {
				pid, err := strconv.Atoi(d[0])
				if err != nil {
					color.New(color.FgRed).Printf("Error: %s", err.Error())
				}
				archiveProject(client, pid)
			}
		}

		if pid == 0 && file == "" {
			color.Red("Either --pid or --file are Required")
		}

	},
}

func archiveProject(client *gitlab.Client, pid int) {
	if !dryRun {
		// proj, resp, err := client.Projects.ArchiveProject(pid)
		// if err != nil {
		// 	color.New(color.FgRed).Printf("Error: %s", err.Error())
		// 	return
		// }

		// if resp.StatusCode == http.StatusCreated {
		// 	color.New(color.FgYellow).Printf("Project '%s' Archived.\n", proj.Name)
		// } else {
		// 	color.New(color.FgRed).Printf("Failed to Archive Project: %s\n", resp.Status)
		// }
	} else {
		color.New(color.FgYellow).Printf("<dry-run> Archive Project: %v\n", pid)
	}
}

func init() {
	rootCmd.AddCommand(archiveCmd)

	archiveCmd.Flags().Int("pid", 0, "project id")
	err := viper.BindPFlag("pid", archiveCmd.Flags().Lookup("pid"))
	if err != nil {
		panic("invalid arg")
	}

	archiveCmd.Flags().String("file", "", "file containing list of projects to archive")
	err = viper.BindPFlag("file", archiveCmd.Flags().Lookup("file"))
	if err != nil {
		panic("invalid arg")
	}
}
