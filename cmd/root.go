package cmd

import (
	"fmt"
	"os"

	"github.com/benmatselby/lionel/trello"

	"github.com/benmatselby/lionel/version"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	initConfig()

	cmd := NewRootCommand()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NewRootCommand will return the application
func NewRootCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "lionel",
		Short:   "CLI application for retrieving data from Trello",
		Version: version.GITCOMMIT,
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lionel.yaml)")

	client := trello.New()

	cmd.AddCommand(
		NewListBoardsCommand(&client),
		NewBurndownCommand(&client),
		NewListCardsCommand(&client),
		NewListCardsPeopleCommand(&client),
		NewCompletionCommand(cmd),
	)

	return cmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".lionel" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".lionel")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if fmt.Sprintf("%T", err) == "ConfigParseError" {
		fmt.Fprintf(os.Stderr, "Failed to load config: %s\n", err)
	}
}
