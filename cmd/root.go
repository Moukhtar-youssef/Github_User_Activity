/*
Copyright © 2025 Moukhtar youssef
*/
package cmd

import (
	"Github_User_Activity/internal/activity"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Github_User_Activity",
	Short: "Print the latest activity of the Github User provided",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunDisplayActivityCmd(args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Padding(1, 2).
			Bold(true).
			Render(fmt.Sprintf("Error: %s", err))
		fmt.Println(errorStyle)
	}
}

func RunDisplayActivityCmd(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Need the url of the Github account")
	}
	username := args[0]
	activities, err := activity.FetchGithubActivity(username)
	if err != nil {
		return err
	}
	return activity.DisplayActivity(username, activities)
}
