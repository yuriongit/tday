/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tday",
	Short: `TDay: A to-do CLI tool that actually makes me look at my
to-do list.`,
	Long: `TDay: A to-do CLI tool that actually makes me look at my
to-do list. 

It's quick, feather-weight, and simply straightforward.
TDay is just to make all the CRUD-related operations for
all my to-dos as simple as and efficent as possible.

---

Because I don't want to deal with GUI-based to-do apps
that are sluggish, inefficent and bloated: I've decided
to create TDay as I like avoiding extensive mouse-use
where I can.

I should be able to: create, read, update, and delete a 
to-do. Commands that I should be available to use:
  tday new ""
  tday ls
  tday fix
  tday rm
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tday.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


