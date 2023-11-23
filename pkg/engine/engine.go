package engine

import (
	"fmt"
	"github.com/dimiro1/banner"
	"github.com/engboustani/go-quartz/internal/database"
	"github.com/engboustani/go-quartz/internal/schedule"
	"github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:     "quartz [command]",
	Short:   "\nQuartz runs specified programs at scheduled times and related tools.",
	Long:    "\nQuartz runs specified programs at scheduled times and related tools.",
	Run:     magic,
	Version: "v1.0.0",
}

func Start() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func magic(cmd *cobra.Command, args []string) {

}

var jobCmd = &cobra.Command{
	Use:   "job [command]",
	Short: "Manage jobs",
	Long:  "Manage list of jobs scheduled",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var jobListCmd = &cobra.Command{
	Use:   "list [command]",
	Short: "Get list of jobs",
	Long:  "Get list of jobs scheduled",
	Run: func(cmd *cobra.Command, args []string) {
		database := database.Database{}
		database.Connect()
		defer database.Close()
		database.GetAllJobs()
	},
}

var jobTailCmd = &cobra.Command{
	Use:   "tail [command]",
	Short: "Get tail of jobs",
	Long:  "Get tail of jobs scheduled",
	Run: func(cmd *cobra.Command, args []string) {
		database := database.Database{}
		database.Connect()
		defer database.Close()
		database.Watch()
		fmt.Scanln()
	},
}

var eventCmd = &cobra.Command{
	Use:   "event [command]",
	Short: "Manage events",
	Long:  "Manage list of event triggered",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var commandCmd = &cobra.Command{
	Use:   "command [command]",
	Short: "Manage commands",
	Long:  "Manage list of command executed",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var runCmd = &cobra.Command{
	Use:   "run [command]",
	Short: "Run quartz",
	Long:  "Run quartz service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Run service")
		schedule := schedule.Schedule{}
		schedule.Init()
		schedule.Start()
		fmt.Scanln()
	},
}

func init() {
	template := `{{ .Title "quartz" "" 2 }}
	{{ .AnsiColor.BrightWhite }}v1.0{{ .AnsiColor.Default }}
	{{ .AnsiColor.BrightCyan }}https://github.com/engboustani/go-quartz{{ .AnsiColor.Default }}
	Now: {{ .Now "Monday, 2 Jan 2006" }}`

	banner.InitString(colorable.NewColorableStdout(), true, true, template)
	println()

	rootCmd.AddCommand(jobCmd, eventCmd, commandCmd, runCmd)
	jobCmd.AddCommand(jobListCmd, jobTailCmd)
}
